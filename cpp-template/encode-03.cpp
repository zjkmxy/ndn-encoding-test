#include "ndn-format-0-3.hpp"

using namespace ndn;

void encode(std::string name){
  Data data{
    .name = NameFromString(name),
    .metaInfo = MetaInfo {
      .contentType = 0,
      .freshnessPeriod = 4000,
      .finalBlockId = GenericNameComponent(std::string("10000"))
    },
    .content = Buffer(10),
    .signatureInfo = SignatureInfo{
      .signatureType = 0
    },
    .signatureValue = std::nullopt
  };
  auto ptr = data.Encode();
}

int main() {
  char str[100];
  for(int i = 0; i < 1000000; i ++){
    sprintf(str, "/go-ndn/test/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/a/%010d", i);
    encode(std::string(str));
  }
  return 0;
}
