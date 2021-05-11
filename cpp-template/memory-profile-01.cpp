#include "../memory_monitor/memory_monitor.hpp"
#include "ndn-format-0-3.hpp"

using namespace ndn;

void encode(std::string name){
  auto& inst = monitor::Monitor::getInstance();
  inst.start();
  Data data{
    .name = NameFromString(name),
    .metaInfo = MetaInfo {
      .contentType = 0,
      .freshnessPeriod = 4000,
      .finalBlockId = GenericNameComponent(std::string("10000"))
    },
    .content = Buffer(100),
    .signatureInfo = SignatureInfo{
      .signatureType = 0
    }
  };
  inst.stopAll();
  inst.writeLogFile("./raw.log");
  size_t size = 0;
  for(const auto& blk: inst.getRecords()){
    size += blk->size;
  }
  printf("sizeof(data) = %lu\n", sizeof(data));
  printf("Referred to Size = %lu\n", size);
  auto ptr = data.Encode();
  printf("sizeof(wire) = %lu\n", sizeof(ptr));
  printf("Content Size = %lu\n", ptr.capacity());
}

int main() {
  char str[100];
  for(int i = 0; i < 1; i ++){
    sprintf(str, "/go-ndn/test/%010d", i);
    encode(std::string(str));
  }
  return 0;
}
