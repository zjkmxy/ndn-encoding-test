#pragma once

#include <cstdio>
#include <cstdint>
#include <cstddef>
#include <type_traits>
#include <list>
#include <tuple>
#include <vector>
#include <string>
#include <optional>
#ifdef __cpp_concepts
#include <concepts>
#endif
#include "slice.hpp"

namespace tlv {


// Directly use binary vector as NameComponent.
using NameComponent = Buffer;
using Name = std::vector<Buffer>;

#ifdef __cpp_concepts
template<typename T>
concept EncodableType =
  requires(T e, uint8_t* buf, size_t buflen) {
    // EncodeSize returns the expected length after encoding
    { e.EncodeSize() } -> std::convertible_to<size_t>;

    // EncodeInto encodes the object into buf.
    // buflen is not checked in this quick and dirty implementation.
    // The actually used size is returned.
    { e.EncodeInto(buf, buflen) } -> std::convertible_to<bool>;
  };

template<typename T>
concept ModelType =
  requires(T model) {
    EncodableType<typename T::Encodable>;
    new T::Encodable(model);
  };

template<typename T, typename Encodable>
concept Encodes =
  requires(T value) {
    EncodableType<Encodable>;
    Encodable(value);
  };

#define REQUIRES_ENCODE(t, e) requires Encodes<t, e>

#else
#define EncodableType typename
#define ModelType typename
#define REQUIRES_ENCODE(t, e)
#endif


// Encoding of integers in BigEndian
namespace big_endian {
  inline void Word(uint16_t val, uint8_t* buf) {
    *(uint16_t*)buf = __builtin_bswap16(val);
  }
  
  inline void Dword(uint32_t val, uint8_t* buf) {
    *(uint32_t*)buf = __builtin_bswap32(val);
  }
  
  inline void Qword(uint64_t val, uint8_t* buf) {
    *(uint64_t*)buf = __builtin_bswap64(val);
  }
};


// In the following section, we design encoders for primitive types and their combinations.
// The idea is similar to Monadic Parsing.


// TlvConst is a TLV type number constant, known at compiling time.
template<uint64_t num>
struct TlvConst {
  TlvConst() = default;

  constexpr size_t EncodeSize() {
    // if constexpr will be computed at compile time.
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

  inline size_t EncodeInto(uint8_t* buf, size_t buflen) {
    if constexpr (num <= 0xfc){
      buf[0] = uint8_t(num);
      return 1;
    } else if constexpr (num <= 0xffff){
      buf[0] = 0xfd;
      big_endian::Word(uint16_t(num), buf + 1);
      return 3;
    } else if constexpr (num <= 0xffffffff){
      buf[0] = 0xfe;
      big_endian::Dword(uint32_t(num), buf + 1);
      return 5;
    } else {
      buf[0] = 0xff;
      big_endian::Qword(uint64_t(num), buf + 1);
      return 9;
    }
  }
};

// TlvVar is a TLV type number variable.
struct TlvVar {
  uint64_t value;
  TlvVar(uint64_t value):value(value){}
  
  inline size_t EncodeSize() const {
    if (value <= 0xfc){
      return 1;
    } else if (value <= 0xffff){
      return 3;
    } else if (value <= 0xffffffff){
      return 5;
    } else {
      return 9;
    }
  }
  
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    if (value <= 0xfc){
      buf[0] = uint8_t(value);
      return 1;
    } else if (value <= 0xffff){
      buf[0] = 0xfd;
      big_endian::Word(uint16_t(value), buf + 1);
      return 3;
    } else if (value <= 0xffffffff){
      buf[0] = 0xfe;
      big_endian::Dword(uint32_t(value), buf + 1);
      return 5;
    } else {
      buf[0] = 0xff;
      big_endian::Qword(uint64_t(value), buf + 1);
      return 9;
    }
  }
};

// NaturalNumber is a natural number, without type and length.
struct NaturalNumber {
  uint64_t value;
  NaturalNumber(uint64_t value):value(value){}

  inline size_t EncodeSize() const {
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

  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    if (value <= 0xff){
      buf[0] = uint8_t(value);
      return 1;
    } else if (value <= 0xffff){
      big_endian::Word(uint16_t(value), buf);
      return 2;
    } else if (value <= 0xffffffff){
      big_endian::Dword(uint32_t(value), buf);
      return 4;
    } else {
      big_endian::Qword(uint64_t(value), buf);
      return 8;
    }
  }
};

// Unit is void type.
// This is used to construct a bool field.
struct Unit {
  Unit() = default;
  constexpr size_t EncodeSize() const {
    return 0;
  }
  constexpr size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    return 0;
  }
};

// ByteString works for std::string, std::vector<uint8_t> and std::array<uint8_t, N>.
template<typename Vector>
struct ByteString {
  const Vector& value;
  inline ByteString(const Vector& value):value(value){}
  inline size_t EncodeSize() const {
    return value.size();
  }
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    std::copy_n(value.begin(), value.size(), buf);
    return value.size();
  }
};

// Sequence is a sequence of any type, i.e. std::vector<T>
// Elements will be encoded in order.
// To make users' life easier, here we separate T and Encodable,
// so users can define std::vector<unit64_t> instead of std::vector<NaturalNumber>.
// But they are fundamentally the same type.
template<typename T, EncodableType Encodable> REQUIRES_ENCODE(T, Encodable)
struct Sequence {
  std::vector<Encodable> encodables;
  size_t length;
  
  inline Sequence(const std::vector<T>& values):length(0){
    encodables.reserve(values.size());
    for(const auto& v: values) {
      encodables.push_back(Encodable(v));
      length += encodables.back().EncodeSize();
    }
  }
  
  inline size_t EncodeSize() const {
    return length;
  }
  
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    size_t pos = 0;
    for(const auto& e: encodables) {
      pos += e.EncodeInto(buf + pos, buflen - pos);
    }
    return pos;
  }
};

// TlvBlock encapsulate an encodable into a block with type and length.
template<uint64_t typeNum, typename T, EncodableType Encodable> REQUIRES_ENCODE(T, Encodable)
struct TlvBlock {
  Encodable encodable;
  size_t length;
  inline TlvBlock(const T& value):encodable(value),length(encodable.EncodeSize()){}
  inline size_t EncodeSize() const {
    size_t ret = TlvConst<typeNum>().EncodeSize();
    ret += TlvVar(length).EncodeSize();
    ret += length;
    return ret;
  }
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    size_t ret = TlvConst<typeNum>().EncodeInto(buf, buflen);
    ret += TlvVar(length).EncodeInto(buf + ret, buflen - ret);
    ret += encodable.EncodeInto(buf + ret, buflen - ret);
    return ret;
  }
};

// OptionalBlock is an optional TLV Block.
template<uint64_t typeNum, typename T, EncodableType Encodable> REQUIRES_ENCODE(T, Encodable)
struct OptionalBlock {
  std::optional<Encodable> encodable;
  size_t length;
  inline OptionalBlock(const std::optional<T>& value):encodable(value),length(0){
    if(encodable.has_value()){
      length = encodable.value().EncodeSize();
    }
  }
  inline size_t EncodeSize() const {
    size_t ret = 0;
    if(encodable.has_value()){
      ret += TlvConst<typeNum>().EncodeSize();
      ret += TlvVar(length).EncodeSize();
      ret += length;
    }
    return ret;
  }
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    size_t ret = 0;
    if(encodable.has_value()){
      ret += TlvConst<typeNum>().EncodeInto(buf, buflen);
      ret += TlvVar(length).EncodeInto(buf + ret, buflen - ret);
      ret += encodable.value().EncodeInto(buf + ret, buflen - ret);
    }
    return ret;
  }
};

// Boolean is a bool such as MustBeFresh.
// Boolean = OptionalBlock<Unit>
template<uint64_t typeNum>
struct Boolean {
  OptionalBlock<typeNum, Unit, Unit> value;
  
  static inline std::optional<Unit> BoolToOptionalUnit(bool val){
    return val ? std::make_optional(Unit()) : std::nullopt;
  }
  
  inline Boolean(bool value):value(BoolToOptionalUnit(value)){}
  
  inline size_t EncodeSize() const {
    return value.EncodeSize();
  }
  
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    return value.EncodeInto(buf, buflen);
  }
};

// Field wraps an encodable into a field of a struct or class.
// Model is the class of TLV model, i.e. the struct holding this field.
// Encodable is the class of original encodable, typically a TlvBlock or OptionalBlock.
// Args are arguments that used to initialize the encodable, in the form of offsets.
// For example, to initialize Encodable(m.a), we need to write
//   StructField<Model, Encodable, &Model::a>
template<typename Model, EncodableType Encodable, auto ...args>
struct Field {
  Encodable encodable;
  inline Field(const Model& model):encodable(model.*args...){}
  inline size_t EncodeSize() const {
    return encodable.EncodeSize();
  }
  inline size_t EncodeInto(uint8_t* buf, size_t buflen) const {
    return encodable.EncodeInto(buf, buflen);
  }
};

// Struct represents a struct that is encodable.
// Encoder parameters specify the encodable fields of this struct/class.
// They will be called in order.
// The only parameter passed to initialize field encoders are the instance of the struct/class.
template<typename Model, typename ...Fields>
struct Struct {
  // Store the encodable instances of fields.
  std::tuple<Fields...> fields;

  // Use the instance to initialize the field encoders.
  inline Struct(const Model& model):fields(std::make_tuple(Fields(model)...)){}

  // enable_if is the way to build recursive functions in compile time.
  // EncodeSize just adds up all sizes of its fields.
  template<std::size_t I = 0>
  inline typename std::enable_if<I == sizeof...(Fields), size_t>::type
  EncodeSize() const {
    return 0;
  }
  
  template<std::size_t I = 0>
  inline typename std::enable_if<I < sizeof...(Fields), size_t>::type
  EncodeSize() const {
    return std::get<I>(fields).EncodeSize() + EncodeSize<I+1>();
  }

  // Similarly, EncodeInto encodes the fields in order.
  template<std::size_t I = 0>
  inline typename std::enable_if<I == sizeof...(Fields), size_t>::type
  EncodeInto(uint8_t* buf, size_t buflen) const {
    return 0;
  }
  
  template<std::size_t I = 0>
  inline typename std::enable_if<I < sizeof...(Fields), size_t>::type
  EncodeInto(uint8_t* buf, size_t buflen) const {
    size_t pos = std::get<I>(fields).EncodeInto(buf, buflen);
    return pos + EncodeInto<I+1>(buf + pos, buflen - pos);
  }

  // Encode automatically allocates a buffer of proper size and encode into it.
  // Not used by Data, but could be used by every struct other than Data and Interest.
  Buffer Encode() const {
    size_t length = EncodeSize();
    Buffer ret(length);
    EncodeInto(ret.data(), length);
    return ret;
  }
};

template<uint64_t typeNum, typename Model, uint64_t Model::* offset>
using NaturalField = Field<Model, TlvBlock<typeNum, uint64_t, NaturalNumber>, offset>;

template<uint64_t typeNum, typename Model, std::optional<uint64_t> Model::* offset>
using NaturalFieldOpt = Field<Model, OptionalBlock<typeNum, uint64_t, NaturalNumber>, offset>;

template<uint64_t typeNum, typename Model, typename Vector, Vector Model::* offset>
using BytesField = Field<Model, TlvBlock<typeNum, Vector, ByteString<Vector>>, offset>;

template<uint64_t typeNum, typename Model, typename Vector, std::optional<Vector> Model::* offset>
using BytesFieldOpt = Field<Model, OptionalBlock<typeNum, Vector, ByteString<Vector>>, offset>;

template<uint64_t typeNum, typename Model, bool Model::* offset>
using BoolField = Field<Model, Boolean<typeNum>, offset>;

// StructField is a field whose type is another encodable struct.
// Every struct is required to have its encodable type defined as Model::Encodable.
template<uint64_t typeNum, typename Model, ModelType StructType, StructType Model::* offset>
using StructField = Field<Model,
                          TlvBlock<typeNum, StructType, typename StructType::Encodable>, offset>;

template<uint64_t typeNum, typename Model, ModelType StructType, std::optional<StructType> Model::* offset>
using StructFieldOpt = Field<Model,
                             OptionalBlock<typeNum, StructType, typename StructType::Encodable>, offset>;


// NameComponentFromString encodes a GenericNameComponent from a string.
// Note: an interesting fact is, the most time consuming part of the whole experiment is new() called by this function.
// Well, copying is necessary for this function.
// But for other usage, should we change to something like Go slice? (Seems that C++20 borrowed_range does not work)
NameComponent GenericNameComponent(const std::string_view& str){
  TlvBlock<8, std::string_view, ByteString<std::string_view>> encodable(str);
  size_t size = encodable.EncodeSize();
  Buffer ret(size);
  encodable.EncodeInto(ret.data(), size);
  return ret; // Automatically moved
}

// NameFromString encodes a Name from a string. Implemented quickly, without support of naming conventions.
// Assume str starts with a "/" and does not ent with a "/".
Name NameFromString(const std::string_view& str){
  auto start = 1UL;
  auto end = str.find("/", start);
  size_t cnt = 0;
  while(end != std::string::npos) {
    cnt ++;
    start = end + 1;
    end = str.find("/", start);
  }
  cnt ++;
  Name ret;
  ret.reserve(cnt);
  start = 1UL;
  end = str.find("/", start);
  while(end != std::string::npos) {
    ret.push_back(GenericNameComponent(str.substr(start, end - start)));
    start = end + 1;
    end = str.find("/", start);
  }
  ret.push_back(GenericNameComponent(str.substr(start, end - start)));
  return ret;
}

using EncodableName = Sequence<Buffer, ByteString<Buffer>>;

template<uint64_t typeNum, typename Model, Name Model::* offset>
using NameField = Field<Model, TlvBlock<typeNum, Name, EncodableName>, offset>;

template<uint64_t typeNum, typename Model, NameComponent Model::* offset>
using NameComponentField = Field<Model,
                                 TlvBlock<typeNum, NameComponent, ByteString<NameComponent>>,
                                 offset>;

template<uint64_t typeNum, typename Model, std::optional<NameComponent> Model::* offset>
using NameComponentFieldOpt = Field<Model,
                                    OptionalBlock<typeNum, NameComponent, ByteString<NameComponent>>,
                                    offset>;

};
