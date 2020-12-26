#pragma once

#include "tlv-encoder.hpp"
#include "sha256/sha256.h"

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
};

struct Data {
  Name name;
  std::optional<MetaInfo> metaInfo;
  // Should be replaced with a memory view to avoid copying on decoding
  std::optional<std::vector<uint8_t>> content;
  SignatureInfo signatureInfo;
  mutable std::vector<uint8_t> signatureValue;
  
  using Encodable = Struct<Data,
    NameField<0x07, Data, &Data::name>,
    StructFieldOpt<0x14, Data, struct MetaInfo, &Data::metaInfo>,
    BytesFieldOpt<0x15, Data, std::vector<uint8_t>, &Data::content>,
    StructField<0x16, Data, decltype(signatureInfo), &Data::signatureInfo>,
    Field<Data, TlvBlock<0x17, Buffer, SignatureValueEncodable<Buffer>>, &Data::signatureValue>>;
    
  Buffer Encode() const {
    Encodable data(*this);
    TlvVar length(data.EncodeSize());
    TlvConst<0x06> typeNum;
    size_t size = typeNum.EncodeSize()
                + length.EncodeSize()
                + data.EncodeSize();
    Buffer ret(size, 0);
    size_t pos = typeNum.EncodeInto(ret.data(), size);
    pos += length.EncodeInto(ret.data() + pos, size - pos);
    data.EncodeInto(ret.data() + pos, size - pos);

    SHA256 ctx = SHA256();
    ctx.init();
    ctx.update(ret.data() + pos, size - pos - 32);
    ctx.final(ret.data() + size - 32);

    return ret;
  }
};

};
