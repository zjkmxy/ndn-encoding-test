#ifndef _GNU_SOURCE
#define _GNU_SOURCE
#endif
#include <ndn-cxx/face.hpp>
#include <ndn-cxx/security/signing-helpers.hpp>
#include <string>
using namespace ndn;
using namespace std;
using security::signingWithSha256;

KeyChain keyChain;

void decode(ConstBufferPtr buf){
  Data data;
  data.wireDecode(Block(buf));
}

int main(int argc, char* argv[]){
  ConstBufferPtr buf = make_shared<const Buffer>(std::initializer_list<uint8_t>(
    {6, 187, 7, 26, 8, 6, 103, 111, 45, 110, 100, 110,
    8, 4, 116, 101, 115, 116, 8, 10, 48, 48, 48, 48, 48, 48, 48, 48, 48, 49,
    20, 16, 24, 1, 0, 25, 2, 15, 160, 26, 7, 8, 5, 49, 48, 48, 48, 48, 21, 100,
    82, 253, 252, 7, 33, 130, 101, 79, 22, 63, 95, 15, 154, 98, 29, 114, 149,
    102, 199, 77, 16, 3, 124, 77, 123, 187, 4, 7, 209, 226, 198, 73, 129, 133,
    90, 216, 104, 29, 13, 134, 209, 233, 30, 0, 22, 121, 57, 203, 102, 148,
    210, 196, 34, 172, 210, 8, 160, 7, 41, 57, 72, 127, 105, 153, 235, 157,
    24, 164, 71, 132, 4, 93, 135, 243, 198, 124, 242, 39, 70, 233, 149, 175,
    90, 37, 54, 121, 81, 186, 162, 255, 108, 212, 113, 196, 131, 241, 95, 185,
    11, 173, 22, 3, 27, 1, 0, 23, 32, 163, 60, 194, 104, 248, 118, 79, 231,
    250, 185, 231, 57, 251, 224, 71, 220, 255, 220, 10, 60, 77, 181, 242, 75,
    60, 83, 208, 210, 38, 129, 234, 198}));

  for(int i = 0; i < 1000000; i ++){
    decode(buf);
  }
  return 0;
}
