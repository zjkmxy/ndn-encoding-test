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

void encode(string name){
  uint8_t buf[10];
  const uint8_t* ptr;

  Data data(name);
  data.setFreshnessPeriod(4000_ms);
  data.setContentType(0);
  data.setFinalBlock(name::Component("10000"));
  data.setContent(buf);
  // keyChain.sign(data, signingWithSha256());
  data.setSignatureInfo(SignatureInfo(tlv::NullSignature));
  data.setSignatureValue(std::make_shared<Buffer>());
  ptr = data.wireEncode().data();
}

int main(int argc, char* argv[]){
  char str[100];
  for(int i = 0; i < 1000000; i ++){
    sprintf(str, "/go-ndn/test/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/%010d", i);
    encode(string(str));
  }
  return 0;
}
