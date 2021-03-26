#include "ndn-format-0-3.hpp"

using namespace ndn;

void decode(const Buffer& wire){
  auto d = Data::Parse(wire);
}

int main() {
  static uint8_t buf[] = {6, 187, 7, 26, 8, 6, 103, 111, 45, 110, 100, 110,
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
    60, 83, 208, 210, 38, 129, 234, 198};

  Buffer wire;
  wire.resize(sizeof(buf));
  std::copy(buf, buf + sizeof(buf), wire.begin());
  for(int i = 0; i < 1000000; i ++){
    decode(wire);
  }
  return 0;
}
