#include "ndn-format-0-3.hpp"

using namespace ndn;
using namespace tlv;

void encode(std::string name){
  Data data{
    NameFromString(name),
    MetaInfo {
      0,
      4000,
      NameComponentFromString("10000")
    },
    std::vector<uint8_t>(100, 0),
    SignatureInfo{
      0
    }
  };
  auto ptr = data.Encode();
}

int main() {
  char str[100];
  for(int i = 0; i < 1000000; i ++){
    sprintf(str, "/go-ndn/test/%d", i);
    encode(std::string(str));
  }
  return 0;
}