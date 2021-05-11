#ifndef _GNU_SOURCE
#define _GNU_SOURCE
#endif
#include "../memory_monitor/memory_monitor.hpp"
#include <ndn-cxx/face.hpp>
#include <ndn-cxx/security/signing-helpers.hpp>
#include <string>
using namespace ndn;
using namespace std;
using security::signingWithSha256;

KeyChain keyChain;

void encode(string name){
  uint8_t buf[100];
  const uint8_t* ptr;

  auto& inst = monitor::Monitor::getInstance();
  inst.start();
  Data data(name);
  data.setFreshnessPeriod(4000_ms);
  data.setContentType(0);
  data.setFinalBlock(name::Component("10000"));
  data.setContent(buf, 100);
  keyChain.sign(data, signingWithSha256());
  ptr = data.wireEncode().wire();
  inst.stopAll();
  inst.writeLogFile("./raw.log");
  size_t size = 0;
  for(const auto& blk: inst.getRecords()){
    size += blk->size;
  }
  printf("sizeof(data) = %lu\n", sizeof(data));
  printf("Referred to Size = %lu\n", size);
}

int main(int argc, char* argv[]){
  char str[100];
  for(int i = 0; i < 1; i ++){
    sprintf(str, "/go-ndn/test/%010d", i);
    encode(string(str));
  }
  return 0;
}
