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

void decode(){
  static uint8_t buf[] =
    {6, 187, 7, 116, 8, 6, 103, 111, 45, 110, 100, 110, 8, 4, 116, 101, 115, 116, 8, 1,
    97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8,
    1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97,
    8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1,
    97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8, 1, 97, 8,
    1, 97, 8, 1, 97, 8, 1, 97, 8, 10, 48, 48, 48, 48, 48, 48, 48, 48, 48, 49,
    20, 16, 24, 1, 0, 25, 2, 15, 160, 26, 7, 8, 5, 49, 48, 48, 48, 48, 21, 10,
    82, 253, 252, 7, 33, 130, 101, 79, 22, 63, 22, 3, 27, 1, 0, 23, 32, 163, 60, 194,
    104, 248, 118, 79, 231, 250, 185, 231, 57, 251, 224, 71, 220, 255, 220, 10, 60, 77, 181, 242,
    75, 60, 83, 208, 210, 38, 129, 234, 198, };

  Data data;
  data.wireDecode(Block(buf, sizeof(buf)));
}

int main(int argc, char* argv[]){
  for(int i = 0; i < 1000000; i ++){
    decode();
  }
  return 0;
}
