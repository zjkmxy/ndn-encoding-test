#pragma once

#include "tlv-encoder.hpp"
#include "sha256/sha256.h"

namespace ndn {

using namespace tlv;

struct MetaInfo {
  uint64_t ContentType;
  uint64_t FreshnessPeriod;
  NameComponent FinalBlockId;
  
  using Encoder = StructEncoder<MetaInfo,
    NaturalFieldEncoder<0x18, MetaInfo, &MetaInfo::ContentType>,
    NaturalFieldEncoder<0x19, MetaInfo, &MetaInfo::FreshnessPeriod>,
    NameComponentFieldEncoder<0x1a, MetaInfo, &MetaInfo::FinalBlockId>>;
};

struct SignatureInfo {
  uint64_t SignatureType;

  using Encoder = StructEncoder<SignatureInfo,
    NaturalFieldEncoder<0x1b, SignatureInfo, &SignatureInfo::SignatureType>>;
};

struct SignatureValueEncoder {
  constexpr size_t EncodeSize() const {
    return 32;
  }
  constexpr size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    return 32;
  }
};

struct Data {
  Name Name;
  MetaInfo MetaInfo;
  std::vector<uint8_t> Content;
  SignatureInfo SignatureInfo;
  std::vector<uint8_t> SignatureValue;
  
  using Encoder = StructEncoder<Data,
    NameFieldEncoder<0x07, Data, &Data::Name>,
    StructFieldEncoder<0x14, Data, decltype(MetaInfo), &Data::MetaInfo>,
    NameComponentFieldEncoder<0x15, Data, &Data::Content>,
    StructFieldEncoder<0x16, Data, decltype(SignatureInfo), &Data::SignatureInfo>,
    FieldEncoder<0x17, Data, SignatureValueEncoder>>;
    
  BufferPtr Encode() const {
    Encoder dataEncoder(*this);
    NaturalNumberEncoder<true> lengthEncoder(dataEncoder.EncodeSize());
    size_t size = TlvConstEncoder<0x06>::EncodeSize()
                + lengthEncoder.EncodeSize()
                + dataEncoder.EncodeSize();
    BufferPtr ret(size, 0);
    size_t pos = TlvConstEncoder<0x06>::EncodeInto(ret.data(), size);
    pos += lengthEncoder.EncodeInto(ret.data() + pos, size - pos);
    dataEncoder.EncodeInto(ret.data() + pos, size - pos);

    SHA256 ctx = SHA256();
    ctx.init();
    ctx.update(ret.data() + pos, size - pos - 32);
    ctx.final(ret.data() + size - 32);

    return ret;
  }
};


};