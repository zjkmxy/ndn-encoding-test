#pragma once

#include <cstdlib>
#include <memory>
#include <sstream>

namespace tlv {

class IndexError: public std::runtime_error{
public:
  IndexError(const std::string& message):std::runtime_error(message){}
};

// Currently allocator is not supported, so T must have a default constructor.
template<typename T>
class Slice {
private:
  std::shared_ptr<T[]> ptr;
  size_t start;
  size_t len;
  size_t cap;

  Slice(const std::shared_ptr<T[]>& ptr, size_t start, size_t len, size_t cap):
    ptr(ptr), start(start), len(len), cap(cap){}

public:
  Slice(): ptr(nullptr), start(0), len(0), cap(0){};
  Slice(size_t length): ptr(new T[length]()), start(0), len(length), cap(length){};
  Slice(const Slice&) = default;
  Slice(Slice&&) = default;

  Slice& operator=(const Slice&& other){
    ptr = std::move(other.ptr);
    start = other.start;
    len = other.len;
    cap = other.cap;
    return *this;
  }

  size_t size() const {
    return len;
  }
  
  size_t capacity() const {
    return cap;
  }
  
  void resize(size_t newSize) {
    if(newSize <= cap) {
      len = newSize;
    } else {
      std::shared_ptr<T[]> newptr(new T[newSize]());
      std::copy_n(&ptr[start], len, &newptr[0]);
      ptr = newptr;
      start = 0;
      cap = newSize;
      len = newSize;
    }
  }
  
  Slice range(size_t first, size_t last) const {
    if(last < first || last > cap){
      std::ostringstream stringStream;
      if(last < first){
        stringStream << "IndexError: [" << first << ", " << last << "] is not a legal range.";
      } else {
        stringStream << "IndexError: [" << first << ", " << last << "] out of scope " << cap << ".";
      }
      throw new IndexError(stringStream.str());
    }
    return Slice(ptr, start + first, last - first, cap - first);
  }

  T& operator[](size_t index) {
    return ptr[start + index];
  }
  
  const T& operator[](size_t index) const {
    return ptr[start + index];
  }
  
  T* begin() {
    return &ptr[start];
  }
  
  const T* begin() const {
    return &ptr[start];
  }
  
  T* end() {
    return &ptr[start+len];
  }
  
  const T* end() const {
    return &ptr[start+len];
  }
  
  T* data() {
    return &ptr[start];
  }
};

// Buffer represents a byte buffer with its length
// returning a vector will not copy its content.
using Buffer = Slice<uint8_t>;

};
