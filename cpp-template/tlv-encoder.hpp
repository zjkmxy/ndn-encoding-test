#pragma once

#include <cstdio>
#include <cstdint>
#include <cstddef>
#include <type_traits>
#include <list>
#include <tuple>
#include <vector>
#include <string>

namespace tlv {

using BufferPtr = std::vector<uint8_t>;

namespace endian {
  inline void BigEndianWord(uint16_t val, uint8_t* buf) {
    buf[0] = ((val >> 8) & 0xFF);
    buf[1] = ((val >> 0) & 0xFF);
  }
  
  inline void BigEndianDword(uint32_t val, uint8_t* buf) {
    buf[0] = ((val >> 24) & 0xFF);
    buf[1] = ((val >> 16) & 0xFF);
    buf[2] = ((val >> 8) & 0xFF);
    buf[3] = ((val >> 0) & 0xFF);
  }
  
  inline void BigEndianQword(uint64_t val, uint8_t* buf) {
    buf[0] = ((val >> 56) & 0xFF);
    buf[1] = ((val >> 48) & 0xFF);
    buf[2] = ((val >> 40) & 0xFF);
    buf[3] = ((val >> 32) & 0xFF);
    buf[4] = ((val >> 24) & 0xFF);
    buf[5] = ((val >> 16) & 0xFF);
    buf[6] = ((val >> 8) & 0xFF);
    buf[7] = ((val >> 0) & 0xFF);
  }
};

template<uint64_t num>
struct TlvConstEncoder {
  static constexpr size_t EncodeSize() {
    if constexpr (num <= 0xfc){
      return 1;
    } else if constexpr (num <= 0xffff){
      return 3;
    } else if constexpr (num <= 0xffffffff){
      return 5;
    } else {
      return 9;
    }
  }
  
  static inline size_t EncodeInto(uint8_t* buf, size_t buflen) {
    if constexpr (num <= 0xfc){
      buf[0] = uint8_t(num);
      return 1;
    } else if constexpr (num <= 0xffff){
      buf[0] = 0xfd;
      endian::BigEndianWord(uint16_t(num), buf + 1);
      return 3;
    } else if constexpr (num <= 0xffffffff){
      buf[0] = 0xfe;
      endian::BigEndianDword(uint32_t(num), buf + 1);
      return 5;
    } else {
      buf[0] = 0xff;
      endian::BigEndianQword(uint64_t(num), buf + 1);
      return 9;
    }
  }
};

template<bool isTlvVar>
struct NaturalNumberEncoder {
  uint64_t value;
  inline NaturalNumberEncoder(uint64_t value):value(value){}
  
  inline size_t EncodeSize() const {
    if constexpr (isTlvVar){
      if (value <= 0xfc){
        return 1;
      } else if (value <= 0xffff){
        return 3;
      } else if (value <= 0xffffffff){
        return 5;
      } else {
        return 9;
      }
    } else {
      if (value <= 0xff){
        return 1;
      } else if (value <= 0xffff){
        return 2;
      } else if (value <= 0xffffffff){
        return 4;
      } else {
        return 8;
      }
    }
  }
  
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    if constexpr (isTlvVar) {
      if (value <= 0xfc){
        buf[0] = uint8_t(value);
        return 1;
      } else if (value <= 0xffff){
        buf[0] = 0xfd;
        endian::BigEndianWord(uint16_t(value), buf + 1);
        return 3;
      } else if (value <= 0xffffffff){
        buf[0] = 0xfe;
        endian::BigEndianDword(uint32_t(value), buf + 1);
        return 5;
      } else {
        buf[0] = 0xff;
        endian::BigEndianQword(uint64_t(value), buf + 1);
        return 9;
      }
    } else {
      if (value <= 0xff){
        buf[0] = uint8_t(value);
        return 1;
      } else if (value <= 0xffff){
        endian::BigEndianWord(uint16_t(value), buf);
        return 2;
      } else if (value <= 0xffffffff){
        endian::BigEndianDword(uint32_t(value), buf);
        return 4;
      } else {
        endian::BigEndianQword(uint64_t(value), buf);
        return 8;
      }
    }
  }
};

struct BinaryStringEncoder {
  const uint8_t *value;
  size_t length;
  inline BinaryStringEncoder(const uint8_t *value, size_t length):value(value),length(length){}
  inline size_t EncodeSize() const {
    return length;
  }
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    std::copy_n(value, length, buf);
    return length;
  }
};

template<uint64_t typeNum, typename Model, typename Encoder, auto ...args>
struct FieldEncoder {
  Encoder encoder;
  size_t length;
  inline FieldEncoder(const Model& model):encoder(model.*args...),length(encoder.EncodeSize()){}
  inline size_t EncodeSize() const {
    size_t ret = TlvConstEncoder<typeNum>::EncodeSize();
    ret += NaturalNumberEncoder<true>(length).EncodeSize();
    ret += length;
    return ret;
  }
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    size_t ret = TlvConstEncoder<typeNum>::EncodeInto(buf, buflen);
    ret += NaturalNumberEncoder<true>(length).EncodeInto(buf + ret, buflen - ret);
    ret += encoder.EncodeInto(buf + ret, buflen - ret);
    return ret;
  }
};

template<uint64_t typeNum, typename Model, uint64_t Model::* offset>
using NaturalFieldEncoder = FieldEncoder<typeNum, Model, NaturalNumberEncoder<false>, offset>;

template<uint64_t typeNum, typename Model, uint8_t* Model::* bufOffset, size_t Model::* lengthOffset>
using BinaryFieldEncoder = FieldEncoder<typeNum, Model, BinaryStringEncoder, bufOffset, lengthOffset>;

template<typename Model, typename ...Encoder>
struct StructEncoder {
  std::tuple<Encoder...> encoders;

  inline StructEncoder(const Model& model):encoders(std::make_tuple(Encoder(model)...)){}

  template<std::size_t I = 0>
  inline typename std::enable_if<I == sizeof...(Encoder), size_t>::type
  EncodeSize() const {
    return 0;
  }
  
  template<std::size_t I = 0>
  inline typename std::enable_if<I < sizeof...(Encoder), size_t>::type
  EncodeSize() const {
    return std::get<I>(encoders).EncodeSize() + EncodeSize<I+1>();
  }
  
  template<std::size_t I = 0>
  inline typename std::enable_if<I == sizeof...(Encoder), size_t>::type
  EncodeInto(uint8_t* buf, size_t buflen) const {
    return 0;
  }
  
  template<std::size_t I = 0>
  inline typename std::enable_if<I < sizeof...(Encoder), size_t>::type
  EncodeInto(uint8_t* buf, size_t buflen) const {
    size_t pos = std::get<I>(encoders).EncodeInto(buf, buflen);
    return pos + EncodeInto<I+1>(buf + pos, buflen - pos);
  }
  
  BufferPtr Encode() const {
    size_t length = EncodeSize();
    BufferPtr ret(length, 0);
    EncodeInto(ret.data(), length);
    return ret;
  }
};

template<uint64_t typeNum, typename Model, typename FieldType, FieldType Model::* offset>
using StructFieldEncoder = FieldEncoder<typeNum, Model, typename FieldType::Encoder, offset>;

using NameComponent = std::vector<uint8_t>;
using Name = std::vector<NameComponent>;

template<typename Vector>
struct ByteVectorEncoder {
  const Vector& value;
  inline ByteVectorEncoder(const Vector& value):value(value){}
  inline size_t EncodeSize() const {
    return value.size();
  }
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    std::copy_n(value.begin(), value.size(), buf);
    return value.size();
  }
};

using StringEncoder = ByteVectorEncoder<std::string>;
template<uint64_t typeNum, typename Model, std::string Model::* offset>
using StringFieldEncoder = FieldEncoder<typeNum, Model, StringEncoder, offset>;


NameComponent NameComponentFromString(const std::string& str){
  NaturalNumberEncoder<true> lengthEncoder(str.length());
  StringEncoder stringEncoder(str);
  size_t size = TlvConstEncoder<8>::EncodeSize() + lengthEncoder.EncodeSize() + str.length();
  std::vector<uint8_t> ret(size, 0);
  size_t pos = TlvConstEncoder<8>::EncodeInto(ret.data(), size);
  pos += lengthEncoder.EncodeInto(ret.data() + pos, size - pos);
  stringEncoder.EncodeInto(ret.data() + pos, size - pos);
  return ret; // Automatically moved
}

Name NameFromString(const std::string& str){
  auto start = 1UL;
  auto end = str.find("/", start);
  size_t cnt = 0;
  while(end != std::string::npos) {
    cnt ++;
    start = end + 1;
    end = str.find("/", start);
  }
  cnt ++;
  Name ret(cnt);
  start = 1UL;
  end = str.find("/", start);
  cnt = 0;
  while(end != std::string::npos) {
    ret[cnt] = NameComponentFromString(str.substr(start, end - start));
    cnt ++;
    start = end + 1;
    end = str.find("/", start);
  }
  ret[cnt] = NameComponentFromString(str.substr(start, end - start));
  return ret;
}

template<typename Type, typename Encoder>
struct SequenceEncoder {
  const std::vector<Type>& value;
  std::vector<Encoder> encoders;
  size_t length;
  inline SequenceEncoder(const std::vector<Type>& value):value(value),length(0){
    encoders.reserve(value.size());
    for(const auto& v: value) {
      encoders.push_back(Encoder(v));
      length += encoders.back().EncodeSize();
    }
  }
  inline size_t EncodeSize() const {
    return length;
  }
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    size_t pos = 0;
    for(const auto& e: encoders) {
      pos += e.EncodeInto(buf + pos, buflen - pos);
    }
    return pos;
  }
};

using NameComponentEncoder = ByteVectorEncoder<NameComponent>;
template<uint64_t typeNum, typename Model, NameComponent Model::* offset>
using NameComponentFieldEncoder = FieldEncoder<typeNum, Model, NameComponentEncoder, offset>;
using NameEncoder = SequenceEncoder<NameComponent, NameComponentEncoder>;
template<uint64_t typeNum, typename Model, Name Model::* offset>
using NameFieldEncoder = FieldEncoder<typeNum, Model, NameEncoder, offset>;

};
