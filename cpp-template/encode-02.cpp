#include "ndn-format-0-3.hpp"

using namespace ndn;

void encode(std::string name){
  Data data{
    NameFromString(name),
    MetaInfo {
      0,
      4000,
      GenericNameComponent(std::string("10000"))
    },
    Buffer(4000),
    SignatureInfo{
      0
    }
  };
  auto ptr = data.Encode();
}

int main() {
  char str[100];
  for(int i = 0; i < 1000000; i ++){
    sprintf(str, "/go-ndn/test/%010d", i);
    encode(std::string(str));
  }
  return 0;
}
