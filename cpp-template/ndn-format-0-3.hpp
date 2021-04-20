#pragma once

#include "tlv-encoder.hpp"
#include "openssl/sha.h"

namespace ndn {

using namespace tlv;

struct MetaInfo {
  std::optional<uint64_t> contentType;
  std::optional<uint64_t> freshnessPeriod;
  std::optional<NameComponent> finalBlockId;
  
  using Encodable = Struct<MetaInfo,
    NaturalFieldOpt<0x18, MetaInfo, &MetaInfo::contentType>,
    NaturalFieldOpt<0x19, MetaInfo, &MetaInfo::freshnessPeriod>,
    NameComponentFieldOpt<0x1a, MetaInfo, &MetaInfo::finalBlockId>>;
};

struct SignatureInfo {
  uint64_t signatureType;

  using Encodable = Struct<SignatureInfo,
    NaturalField<0x1b, SignatureInfo, &SignatureInfo::signatureType>>;
};

template<typename Vector>
struct SignatureValueEncodable {
  inline SignatureValueEncodable(const Vector& value){}
  inline size_t EncodeSize() const {
    return 32;
  }
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    return 32;
  }

  static inline ParseResult<Vector> Parse(const Buffer& wire) {
    return {wire, wire.size()};
  }
};

struct Data {
  Name name;
  std::optional<MetaInfo> metaInfo;
  // Should be replaced with a memory view to avoid copying on decoding
  std::optional<Buffer> content;
  SignatureInfo signatureInfo;
  mutable Buffer signatureValue;
  
  using Encodable = Struct<Data,
    NameField<0x07, Data, &Data::name>,
    StructFieldOpt<0x14, Data, struct MetaInfo, &Data::metaInfo>,
    BytesFieldOpt<0x15, Data, Buffer, &Data::content>,
    StructField<0x16, Data, decltype(signatureInfo), &Data::signatureInfo>,
    Field<Data, TlvBlock<0x17, Buffer, SignatureValueEncodable<Buffer>>, &Data::signatureValue>>;
    
  Buffer Encode() const {
    Encodable data(*this);
    TlvVar length(data.EncodeSize());
    TlvConst<0x06> typeNum;
    size_t size = typeNum.EncodeSize()
                + length.EncodeSize()
                + data.EncodeSize();
    Buffer ret(size);
    size_t pos = typeNum.EncodeInto(ret.data(), size);
    pos += length.EncodeInto(ret.data() + pos, size - pos);
    data.EncodeInto(ret.data() + pos, size - pos);

    SHA256_CTX ctx;
    SHA256_Init(&ctx);
    SHA256_Update(&ctx, ret.data() + pos, size - pos - 32);
    SHA256_Final(ret.data() + size - 32, &ctx);

    return ret;
  }

  static std::optional<Data> Parse(Buffer wire) {
    size_t pos = 0;
    const auto& [typ, tsiz] = TlvConst<0x06>::Parse(wire);
    pos += tsiz;
    if(!typ){
      return std::nullopt;
    }

    const auto& [length, lsiz] = TlvVar::Parse(wire.range(pos, wire.size()));
    pos += lsiz;
    if(!length || pos + length.value() > wire.size()){
      return std::nullopt;
    }

    const auto& [ret, vsiz] = Encodable::Parse(wire.range(pos, pos+length.value()));
    return ret;
  }
};

};
