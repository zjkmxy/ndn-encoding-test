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
    {6, 253, 15, 249, 7, 26, 8, 6, 103, 111, 45, 110, 100, 110, 8, 4, 116, 101, 115, 116,
    8, 10, 48, 48, 48, 48, 48, 48, 48, 48, 48, 49, 20, 16, 24, 1, 0, 25, 2, 15,
    160, 26, 7, 8, 5, 49, 48, 48, 48, 48, 21, 253, 15, 160, 82, 253, 252, 7, 33, 130,
    101, 79, 22, 63, 95, 15, 154, 98, 29, 114, 149, 102, 199, 77, 16, 3, 124, 77, 123, 187,
    4, 7, 209, 226, 198, 73, 129, 133, 90, 216, 104, 29, 13, 134, 209, 233, 30, 0, 22, 121,
    57, 203, 102, 148, 210, 196, 34, 172, 210, 8, 160, 7, 41, 57, 72, 127, 105, 153, 235, 157,
    24, 164, 71, 132, 4, 93, 135, 243, 198, 124, 242, 39, 70, 233, 149, 175, 90, 37, 54, 121,
    81, 186, 162, 255, 108, 212, 113, 196, 131, 241, 95, 185, 11, 173, 179, 124, 88, 33, 182, 217,
    85, 38, 164, 26, 149, 4, 104, 11, 78, 124, 139, 118, 58, 27, 29, 73, 212, 149, 92, 132,
    134, 33, 99, 37, 37, 63, 236, 115, 141, 215, 169, 226, 139, 249, 33, 17, 156, 22, 15, 7,
    2, 68, 134, 21, 187, 218, 8, 49, 63, 106, 142, 182, 104, 210, 11, 245, 5, 152, 117, 146,
    30, 102, 138, 91, 223, 44, 127, 196, 132, 69, 146, 210, 87, 43, 205, 6, 104, 210, 214, 197,
    47, 80, 84, 226, 208, 131, 107, 248, 76, 113, 116, 203, 116, 118, 54, 76, 195, 219, 217, 104,
    176, 247, 23, 46, 216, 87, 148, 187, 53, 139, 12, 59, 82, 93, 161, 120, 111, 159, 255, 9,
    66, 121, 219, 25, 68, 235, 215, 161, 157, 15, 123, 186, 203, 224, 37, 90, 165, 183, 212, 75,
    236, 64, 248, 76, 137, 43, 155, 255, 212, 54, 41, 176, 34, 59, 238, 165, 244, 247, 67, 145,
    244, 69, 209, 90, 253, 66, 148, 4, 3, 116, 246, 146, 75, 152, 203, 248, 113, 63, 141, 150,
    45, 124, 141, 1, 145, 146, 194, 66, 36, 226, 202, 252, 202, 227, 166, 31, 181, 134, 177, 67,
    35, 166, 188, 143, 158, 125, 241, 217, 41, 51, 63, 249, 147, 147, 59, 234, 111, 91, 58, 246,
    222, 3, 116, 54, 108, 71, 25, 228, 58, 27, 6, 125, 137, 188, 127, 1, 241, 245, 115, 152,
    22, 89, 164, 79, 241, 122, 76, 114, 21, 163, 181, 57, 235, 30, 88, 73, 198, 7, 125, 187,
    87, 34, 245, 113, 122, 40, 154, 38, 111, 151, 100, 121, 129, 153, 142, 190, 168, 156, 11, 75,
    55, 57, 112, 17, 94, 130, 237, 111, 65, 37, 200, 250, 115, 17, 228, 215, 222, 250, 146, 45,
    170, 231, 120, 102, 103, 247, 233, 54, 205, 79, 36, 171, 247, 223, 134, 107, 170, 86, 3, 131,
    103, 173, 97, 69, 222, 30, 232, 244, 168, 176, 153, 62, 189, 248, 136, 58, 10, 216, 190, 156,
    57, 120, 176, 72, 131, 229, 106, 21, 106, 141, 229, 99, 175, 164, 103, 212, 157, 236, 106, 64,
    233, 161, 208, 7, 240, 51, 194, 130, 48, 97, 189, 208, 234, 165, 159, 142, 77, 166, 67, 1,
    5, 34, 13, 11, 41, 104, 139, 115, 75, 142, 160, 243, 202, 153, 54, 232, 70, 31, 16, 215,
    124, 150, 234, 128, 167, 166, 101, 246, 6, 246, 166, 59, 127, 61, 253, 37, 103, 193, 137, 121,
    228, 214, 15, 38, 104, 109, 155, 242, 251, 38, 201, 1, 255, 53, 76, 222, 22, 7, 238, 41,
    75, 57, 243, 43, 124, 120, 34, 186, 100, 248, 74, 180, 60, 160, 198, 230, 185, 28, 31, 211,
    190, 137, 144, 67, 65, 121, 211, 175, 68, 145, 163, 105, 1, 45, 185, 45, 24, 79, 195, 157,
    23, 52, 255, 87, 22, 66, 137, 83, 187, 104, 101, 252, 249, 43, 12, 58, 23, 201, 2, 139,
    233, 145, 78, 183, 100, 156, 108, 147, 71, 128, 9, 121, 209, 131, 3, 86, 242, 165, 76, 61,
    234, 178, 164, 180, 71, 93, 99, 175, 190, 143, 181, 105, 135, 199, 127, 88, 24, 82, 111, 24,
    20, 190, 130, 51, 80, 234, 177, 57, 53, 243, 29, 132, 72, 69, 23, 233, 36, 174, 247, 138,
    225, 81, 192, 7, 85, 146, 88, 54, 183, 7, 88, 133, 101, 12, 48, 236, 41, 163, 112, 57,
    52, 191, 80, 162, 141, 161, 2, 151, 93, 237, 167, 126, 117, 133, 121, 234, 61, 254, 65, 54,
    171, 247, 82, 179, 184, 39, 29, 3, 233, 68, 179, 201, 219, 54, 107, 117, 4, 95, 142, 253,
    105, 210, 42, 229, 65, 25, 71, 203, 85, 61, 118, 148, 38, 122, 239, 78, 188, 234, 64, 107,
    50, 214, 16, 139, 214, 133, 132, 245, 126, 55, 202, 172, 110, 51, 254, 170, 50, 99, 163, 153,
    67, 112, 36, 186, 156, 155, 20, 103, 138, 39, 79, 1, 169, 16, 174, 41, 95, 110, 251, 254,
    95, 90, 191, 68, 204, 222, 38, 59, 86, 6, 99, 62, 43, 240, 0, 111, 40, 41, 93, 125,
    57, 6, 159, 1, 162, 57, 196, 54, 88, 84, 195, 175, 127, 107, 65, 214, 49, 249, 43, 154,
    141, 18, 244, 18, 87, 50, 95, 255, 51, 47, 117, 118, 176, 98, 5, 86, 48, 74, 62, 62,
    174, 20, 194, 141, 12, 234, 57, 210, 144, 26, 82, 114, 13, 168, 92, 161, 228, 179, 142, 175,
    63, 68, 198, 198, 239, 131, 98, 242, 245, 79, 192, 14, 9, 214, 252, 37, 100, 8, 84, 193,
    93, 252, 172, 170, 138, 44, 236, 206, 90, 58, 186, 83, 171, 112, 91, 24, 219, 148, 180, 211,
    56, 165, 20, 62, 99, 64, 141, 135, 36, 176, 207, 63, 174, 23, 163, 247, 155, 225, 7, 47,
    182, 60, 53, 214, 4, 44, 65, 96, 243, 142, 233, 226, 169, 243, 251, 79, 251, 0, 25, 180,
    84, 213, 34, 181, 255, 161, 118, 4, 25, 63, 184, 150, 103, 16, 167, 150, 7, 50, 202, 82,
    207, 83, 195, 245, 32, 200, 137, 183, 155, 245, 4, 207, 181, 124, 118, 1, 35, 45, 88, 155,
    172, 206, 169, 214, 226, 99, 226, 92, 39, 116, 29, 63, 108, 98, 203, 187, 21, 217, 175, 188,
    191, 127, 125, 164, 26, 176, 64, 142, 57, 105, 194, 226, 205, 207, 35, 52, 56, 191, 23, 116,
    172, 231, 112, 154, 79, 9, 30, 154, 131, 253, 234, 224, 236, 85, 235, 35, 58, 155, 83, 148,
    203, 60, 120, 86, 181, 70, 211, 19, 200, 163, 180, 193, 192, 224, 84, 71, 244, 186, 55, 14,
    179, 109, 188, 253, 236, 144, 179, 2, 220, 220, 59, 158, 245, 34, 226, 166, 241, 237, 10, 254,
    193, 248, 226, 15, 170, 190, 223, 107, 22, 46, 113, 125, 58, 116, 138, 88, 103, 122, 12, 86,
    52, 143, 137, 33, 162, 102, 177, 29, 15, 51, 76, 98, 254, 82, 186, 83, 175, 25, 119, 156,
    178, 148, 139, 101, 112, 255, 160, 183, 115, 150, 60, 19, 10, 215, 151, 221, 234, 254, 78, 58,
    210, 155, 81, 37, 33, 15, 14, 241, 195, 20, 9, 15, 7, 199, 154, 111, 87, 28, 36, 111,
    62, 154, 192, 183, 65, 62, 241, 16, 189, 88, 176, 12, 231, 59, 255, 112, 111, 127, 244, 182,
    244, 64, 144, 163, 39, 17, 243, 32, 142, 78, 75, 137, 203, 81, 101, 206, 100, 0, 44, 189,
    156, 40, 135, 170, 17, 61, 242, 70, 137, 40, 213, 162, 59, 156, 167, 64, 248, 12, 147, 130,
    217, 198, 3, 74, 210, 150, 12, 121, 101, 3, 225, 206, 34, 23, 37, 245, 12, 175, 31, 191,
    232, 49, 177, 11, 123, 245, 177, 92, 71, 165, 61, 191, 142, 125, 202, 252, 158, 19, 134, 71,
    164, 180, 78, 212, 188, 233, 100, 237, 71, 247, 74, 165, 148, 70, 140, 237, 50, 60, 183, 111,
    13, 63, 172, 71, 108, 159, 176, 63, 201, 34, 143, 186, 232, 143, 213, 128, 102, 58, 4, 84,
    182, 131, 18, 32, 127, 10, 59, 88, 76, 98, 49, 100, 146, 180, 151, 83, 181, 213, 2, 124,
    225, 90, 79, 10, 88, 37, 13, 143, 181, 14, 119, 242, 191, 79, 1, 82, 229, 212, 148, 53,
    128, 127, 157, 75, 151, 190, 111, 183, 121, 112, 70, 106, 86, 38, 254, 51, 64, 140, 249, 232,
    142, 44, 121, 116, 8, 163, 45, 41, 65, 107, 175, 32, 106, 50, 156, 255, 253, 74, 117, 228,
    152, 50, 9, 130, 200, 90, 173, 112, 56, 72, 89, 192, 90, 75, 19, 161, 213, 178, 245, 191,
    239, 90, 110, 217, 45, 164, 130, 202, 169, 86, 142, 91, 111, 233, 216, 169, 221, 217, 235, 9,
    39, 123, 146, 206, 249, 4, 110, 250, 24, 80, 9, 68, 203, 232, 0, 160, 177, 82, 126, 166,
    71, 41, 168, 97, 210, 246, 73, 122, 50, 53, 195, 127, 65, 146, 119, 158, 193, 217, 107, 59,
    28, 84, 36, 252, 224, 183, 39, 176, 48, 114, 230, 65, 90, 118, 31, 3, 171, 170, 64, 171,
    201, 68, 143, 221, 235, 33, 145, 217, 69, 192, 71, 103, 175, 132, 122, 253, 14, 219, 93, 136,
    87, 183, 153, 172, 177, 142, 74, 255, 171, 227, 3, 127, 254, 127, 166, 138, 168, 175, 94, 57,
    204, 65, 110, 115, 77, 55, 60, 94, 190, 188, 156, 220, 197, 149, 188, 206, 60, 123, 211, 216,
    223, 147, 250, 183, 225, 37, 221, 235, 175, 230, 90, 49, 189, 93, 65, 226, 210, 206, 156, 43,
    23, 137, 47, 15, 234, 25, 49, 162, 144, 34, 7, 119, 169, 49, 67, 223, 220, 191, 166, 132,
    6, 232, 119, 7, 63, 240, 136, 52, 225, 151, 164, 3, 74, 164, 138, 250, 63, 133, 184, 166,
    39, 8, 202, 235, 186, 200, 128, 181, 184, 155, 147, 218, 83, 129, 1, 100, 64, 33, 4, 230,
    72, 182, 34, 106, 27, 120, 2, 24, 81, 245, 217, 172, 15, 49, 58, 137, 221, 252, 69, 76,
    95, 143, 114, 172, 137, 179, 139, 25, 245, 55, 132, 193, 158, 155, 234, 192, 60, 135, 90, 39,
    219, 2, 157, 227, 122, 227, 122, 66, 49, 136, 19, 72, 118, 133, 146, 147, 89, 202, 140, 94,
    185, 78, 21, 45, 193, 175, 66, 234, 61, 22, 118, 193, 189, 209, 154, 184, 226, 146, 92, 109,
    174, 228, 222, 94, 249, 249, 220, 240, 141, 252, 189, 2, 184, 8, 9, 57, 133, 133, 146, 138,
    15, 125, 229, 11, 225, 166, 220, 29, 87, 104, 232, 83, 121, 136, 253, 220, 229, 98, 233, 185,
    72, 201, 24, 187, 163, 233, 51, 229, 196, 0, 205, 229, 230, 12, 94, 173, 111, 199, 174, 119,
    186, 29, 37, 155, 24, 138, 75, 33, 200, 111, 188, 35, 215, 40, 180, 83, 71, 234, 218, 101,
    10, 242, 76, 86, 208, 128, 10, 134, 145, 51, 32, 136, 168, 5, 189, 85, 196, 70, 226, 94,
    176, 117, 144, 186, 252, 204, 190, 198, 23, 117, 54, 64, 29, 154, 43, 127, 81, 43, 84, 191,
    201, 208, 5, 50, 173, 245, 170, 167, 195, 169, 107, 197, 155, 72, 159, 119, 217, 4, 44, 91,
    206, 38, 177, 99, 222, 253, 229, 238, 106, 15, 187, 62, 147, 70, 206, 248, 31, 10, 233, 81,
    94, 243, 15, 164, 122, 54, 78, 117, 174, 169, 225, 17, 213, 150, 230, 133, 165, 145, 18, 25,
    102, 224, 49, 101, 13, 81, 3, 84, 170, 132, 85, 128, 255, 86, 7, 96, 253, 54, 81, 76,
    161, 151, 200, 117, 241, 208, 45, 146, 22, 235, 167, 98, 126, 35, 152, 50, 46, 181, 207, 67,
    215, 43, 210, 229, 184, 135, 212, 99, 15, 184, 212, 116, 126, 173, 110, 184, 42, 205, 28, 91,
    7, 129, 67, 238, 38, 165, 134, 173, 35, 19, 157, 80, 65, 114, 52, 112, 191, 36, 168, 101,
    131, 124, 145, 35, 70, 28, 65, 245, 255, 153, 170, 153, 206, 36, 235, 77, 120, 133, 118, 227,
    51, 110, 101, 73, 22, 34, 85, 143, 223, 41, 123, 159, 160, 7, 134, 75, 175, 215, 205, 76,
    161, 178, 251, 87, 102, 171, 67, 26, 3, 43, 114, 185, 167, 233, 55, 237, 100, 141, 8, 1,
    242, 144, 85, 211, 9, 13, 36, 99, 113, 130, 84, 249, 68, 36, 131, 199, 185, 139, 147, 128,
    69, 218, 81, 152, 67, 133, 75, 14, 211, 247, 186, 149, 26, 73, 63, 50, 31, 9, 102, 96,
    48, 34, 193, 223, 197, 121, 185, 158, 217, 210, 13, 87, 58, 213, 49, 113, 200, 254, 247, 241,
    244, 228, 97, 59, 179, 101, 178, 235, 180, 79, 15, 251, 105, 7, 19, 99, 133, 205, 200, 56,
    240, 189, 212, 200, 18, 240, 66, 87, 116, 16, 172, 160, 8, 194, 175, 188, 76, 121, 198, 37,
    114, 226, 15, 142, 217, 78, 230, 43, 77, 231, 170, 28, 200, 76, 136, 126, 31, 124, 49, 233,
    39, 223, 229, 42, 95, 143, 70, 98, 126, 181, 211, 164, 254, 22, 250, 252, 226, 54, 35, 225,
    150, 201, 223, 255, 127, 186, 255, 79, 254, 148, 244, 88, 151, 51, 229, 99, 225, 157, 48, 69,
    170, 211, 226, 38, 72, 138, 192, 44, 202, 66, 145, 174, 209, 105, 220, 229, 3, 157, 106, 176,
    14, 64, 246, 122, 171, 41, 51, 45, 225, 68, 139, 53, 80, 124, 124, 138, 9, 196, 219, 7,
    16, 93, 195, 16, 3, 98, 4, 5, 218, 59, 33, 105, 245, 169, 16, 201, 208, 9, 110, 94,
    62, 241, 181, 112, 104, 7, 70, 172, 208, 204, 119, 96, 51, 27, 102, 49, 56, 214, 211, 66,
    176, 81, 181, 223, 65, 6, 55, 207, 122, 238, 155, 12, 140, 16, 168, 249, 152, 6, 48, 243,
    76, 224, 1, 192, 171, 122, 198, 94, 80, 45, 57, 178, 22, 203, 197, 14, 115, 163, 46, 175,
    147, 100, 1, 226, 80, 107, 216, 184, 44, 48, 211, 70, 188, 75, 47, 163, 25, 242, 69, 168,
    101, 126, 193, 34, 234, 244, 173, 84, 37, 194, 73, 238, 22, 14, 23, 185, 85, 65, 194, 174,
    229, 223, 130, 10, 200, 93, 227, 248, 231, 132, 135, 15, 216, 122, 54, 204, 13, 22, 56, 51,
    223, 99, 102, 19, 169, 204, 148, 116, 55, 182, 89, 40, 53, 185, 246, 244, 248, 192, 231, 13,
    190, 235, 174, 123, 20, 205, 185, 188, 65, 3, 58, 165, 186, 244, 13, 69, 226, 77, 114, 234,
    196, 162, 142, 60, 160, 48, 201, 147, 122, 184, 64, 154, 124, 191, 5, 174, 33, 249, 116, 37,
    37, 69, 67, 217, 77, 17, 89, 0, 185, 10, 231, 3, 185, 125, 152, 86, 210, 68, 29, 20,
    186, 73, 166, 119, 222, 139, 24, 203, 69, 75, 153, 221, 217, 218, 167, 204, 187, 117, 0, 218,
    228, 226, 229, 223, 140, 243, 133, 158, 189, 218, 218, 103, 69, 251, 166, 160, 76, 92, 55, 199,
    202, 53, 3, 111, 17, 115, 44, 232, 188, 39, 180, 136, 104, 97, 31, 199, 60, 130, 164, 145,
    191, 171, 215, 161, 157, 245, 15, 220, 120, 165, 93, 187, 194, 253, 55, 249, 41, 101, 102, 85,
    127, 171, 136, 91, 3, 159, 48, 231, 6, 240, 205, 89, 97, 225, 155, 100, 34, 33, 219, 68,
    166, 148, 151, 184, 173, 153, 64, 143, 225, 224, 55, 198, 139, 247, 197, 229, 222, 29, 44, 104,
    25, 35, 72, 236, 17, 137, 251, 46, 54, 151, 60, 239, 9, 255, 20, 190, 35, 146, 40, 1,
    246, 234, 238, 65, 64, 145, 88, 180, 95, 45, 236, 130, 209, 124, 170, 186, 22, 12, 214, 64,
    255, 115, 73, 95, 228, 160, 92, 225, 32, 44, 167, 40, 126, 211, 35, 91, 149, 230, 159, 87,
    31, 165, 230, 86, 170, 165, 31, 174, 30, 189, 215, 170, 98, 105, 194, 236, 127, 64, 87, 179,
    53, 147, 188, 132, 136, 140, 151, 15, 213, 40, 212, 169, 154, 30, 171, 157, 36, 32, 19, 69,
    55, 205, 109, 2, 40, 46, 9, 129, 225, 64, 35, 42, 74, 135, 56, 58, 33, 209, 132, 92,
    64, 138, 215, 87, 4, 56, 19, 3, 42, 11, 213, 163, 13, 204, 166, 227, 170, 45, 240, 71,
    21, 216, 121, 39, 154, 150, 135, 154, 79, 54, 144, 172, 32, 37, 166, 12, 125, 177, 94, 5,
    1, 235, 195, 75, 115, 67, 85, 254, 74, 5, 155, 211, 137, 157, 146, 14, 149, 241, 196, 109,
    67, 47, 155, 8, 230, 77, 127, 155, 56, 150, 93, 90, 119, 167, 172, 24, 60, 56, 51, 225,
    163, 66, 94, 173, 105, 212, 249, 117, 1, 47, 209, 164, 158, 216, 50, 246, 158, 110, 156, 99,
    180, 83, 236, 4, 156, 158, 122, 92, 249, 68, 35, 45, 16, 53, 63, 100, 67, 74, 186, 224,
    96, 246, 80, 106, 211, 253, 177, 244, 65, 91, 10, 249, 206, 140, 32, 139, 194, 14, 229, 38,
    116, 21, 57, 250, 50, 3, 199, 126, 203, 164, 16, 253, 103, 24, 242, 39, 224, 180, 48, 249,
    188, 176, 73, 163, 211, 133, 64, 220, 34, 41, 105, 18, 12, 232, 15, 32, 7, 205, 66, 167,
    8, 167, 33, 170, 41, 152, 123, 69, 212, 228, 40, 129, 25, 132, 236, 173, 52, 156, 195, 93,
    217, 53, 21, 206, 254, 11, 0, 44, 238, 94, 113, 196, 121, 53, 226, 129, 235, 252, 75, 139,
    101, 43, 105, 204, 176, 146, 229, 90, 32, 241, 185, 249, 125, 4, 98, 150, 18, 70, 33, 146,
    135, 57, 168, 102, 113, 204, 24, 1, 82, 185, 83, 227, 191, 157, 25, 248, 37, 195, 221, 84,
    174, 22, 136, 228, 158, 251, 94, 254, 101, 220, 218, 211, 75, 200, 96, 1, 14, 124, 140, 153,
    124, 213, 249, 227, 32, 202, 125, 57, 212, 186, 128, 26, 23, 91, 28, 118, 240, 87, 131, 47,
    63, 54, 215, 216, 147, 226, 22, 228, 199, 187, 219, 84, 141, 11, 164, 132, 73, 51, 0, 39,
    54, 139, 52, 249, 198, 151, 118, 180, 89, 21, 50, 218, 28, 91, 230, 142, 244, 238, 190, 140,
    184, 250, 125, 197, 72, 63, 183, 12, 44, 137, 99, 52, 203, 31, 156, 181, 223, 224, 68, 250,
    8, 97, 151, 255, 93, 253, 2, 242, 186, 56, 132, 197, 61, 215, 24, 200, 86, 13, 167, 67,
    168, 233, 212, 174, 174, 32, 204, 239, 0, 45, 130, 202, 53, 37, 146, 184, 216, 242, 168, 223,
    59, 12, 53, 241, 91, 155, 55, 13, 202, 128, 212, 202, 142, 154, 19, 62, 181, 32, 148, 242,
    221, 92, 8, 115, 31, 82, 49, 93, 130, 136, 70, 227, 125, 246, 143, 209, 6, 88, 180, 128,
    242, 172, 132, 35, 54, 51, 149, 126, 104, 142, 146, 79, 254, 55, 19, 181, 44, 118, 253, 138,
    86, 218, 139, 176, 125, 170, 142, 180, 235, 143, 115, 52, 249, 146, 86, 226, 118, 106, 65, 9,
    21, 14, 237, 66, 79, 15, 116, 53, 67, 205, 234, 102, 229, 186, 170, 3, 237, 201, 24, 232,
    48, 91, 177, 159, 192, 198, 180, 221, 180, 170, 56, 134, 203, 80, 144, 148, 15, 198, 212, 202,
    190, 33, 83, 128, 158, 78, 214, 10, 14, 42, 240, 127, 27, 42, 107, 181, 166, 1, 122, 87,
    138, 39, 203, 220, 32, 161, 117, 159, 118, 176, 136, 154, 131, 206, 37, 206, 60, 169, 26, 78,
    181, 194, 248, 88, 8, 25, 218, 4, 208, 44, 65, 119, 12, 1, 116, 109, 228, 79, 61, 182,
    227, 64, 46, 120, 115, 219, 118, 53, 81, 110, 135, 179, 62, 75, 65, 43, 163, 223, 104, 84,
    73, 32, 245, 234, 39, 236, 9, 119, 16, 149, 79, 66, 21, 139, 219, 166, 109, 72, 20, 192,
    100, 180, 17, 37, 56, 103, 96, 149, 70, 124, 137, 186, 152, 230, 165, 67, 117, 141, 112, 147,
    164, 148, 223, 92, 195, 109, 9, 199, 166, 71, 42, 65, 242, 156, 56, 10, 152, 123, 30, 205,
    207, 132, 118, 95, 78, 93, 60, 238, 252, 28, 2, 24, 31, 87, 15, 68, 252, 214, 41, 240,
    141, 193, 239, 83, 201, 174, 13, 136, 105, 254, 103, 253, 199, 162, 198, 123, 66, 95, 19, 197,
    190, 141, 159, 99, 12, 29, 6, 60, 2, 253, 117, 207, 100, 193, 174, 201, 210, 226, 239, 110,
    100, 49, 213, 245, 173, 4, 137, 7, 141, 198, 31, 70, 73, 77, 204, 244, 3, 218, 215, 240,
    148, 23, 13, 44, 62, 41, 193, 152, 176, 243, 65, 226, 132, 196, 190, 143, 166, 12, 26, 71,
    141, 107, 213, 93, 210, 192, 77, 173, 134, 210, 5, 61, 93, 37, 176, 20, 227, 216, 182, 67,
    34, 205, 203, 80, 4, 250, 164, 108, 250, 45, 106, 210, 255, 147, 59, 195, 189, 154, 90, 116,
    102, 10, 243, 208, 72, 169, 164, 54, 52, 192, 37, 4, 39, 217, 166, 33, 145, 151, 163, 243,
    99, 63, 132, 23, 83, 186, 124, 39, 243, 97, 159, 56, 123, 107, 26, 108, 185, 193, 220, 34,
    118, 116, 170, 2, 7, 36, 209, 55, 218, 44, 184, 123, 22, 21, 213, 18, 151, 79, 164, 116,
    125, 209, 225, 125, 2, 201, 70, 42, 68, 254, 193, 80, 202, 58, 143, 153, 204, 30, 73, 83,
    54, 94, 66, 153, 86, 94, 16, 133, 53, 177, 246, 46, 29, 75, 161, 142, 23, 165, 33, 100,
    65, 139, 253, 26, 147, 63, 127, 179, 161, 38, 200, 96, 131, 10, 135, 41, 61, 146, 113, 218,
    115, 110, 67, 152, 193, 227, 127, 183, 92, 75, 240, 39, 134, 225, 250, 244, 182, 16, 205, 19,
    119, 251, 185, 174, 24, 6, 85, 160, 171, 239, 186, 215, 0, 192, 148, 115, 70, 159, 30, 202,
    90, 102, 213, 63, 163, 220, 124, 211, 231, 195, 176, 65, 29, 126, 20, 95, 150, 235, 150, 84,
    171, 148, 145, 61, 218, 80, 58, 80, 249, 231, 115, 132, 47, 77, 42, 95, 170, 96, 134, 155,
    243, 101, 131, 5, 17, 242, 237, 237, 208, 62, 10, 115, 0, 14, 219, 96, 201, 162, 154, 95,
    94, 25, 76, 243, 181, 102, 122, 105, 70, 144, 56, 69, 153, 209, 22, 248, 210, 253, 147, 178,
    174, 213, 91, 125, 68, 181, 176, 84, 243, 243, 142, 120, 142, 79, 223, 54, 229, 145, 86, 140,
    65, 209, 5, 44, 173, 15, 203, 104, 202, 76, 75, 245, 9, 13, 87, 223, 157, 182, 240, 217,
    29, 216, 177, 27, 128, 79, 51, 26, 219, 126, 251, 8, 122, 86, 4, 233, 226, 43, 77, 84,
    219, 64, 188, 188, 110, 39, 47, 245, 234, 221, 252, 20, 113, 69, 158, 89, 240, 85, 76, 88,
    37, 19, 66, 19, 74, 141, 170, 239, 20, 152, 6, 155, 165, 129, 239, 29, 162, 81, 11, 233,
    40, 67, 72, 122, 78, 184, 17, 28, 121, 166, 240, 25, 95, 195, 138, 214, 174, 233, 60, 29,
    242, 181, 137, 126, 170, 56, 173, 143, 71, 171, 47, 224, 227, 170, 62, 106, 204, 191, 212, 193,
    109, 70, 132, 51, 24, 95, 198, 28, 134, 27, 150, 202, 101, 227, 77, 49, 242, 77, 111, 86,
    238, 133, 9, 35, 20, 164, 215, 101, 98, 5, 193, 83, 34, 241, 201, 118, 19, 192, 121, 234,
    226, 146, 186, 150, 110, 16, 209, 231, 0, 22, 78, 81, 139, 36, 63, 66, 76, 70, 249, 234,
    99, 219, 28, 44, 52, 181, 18, 196, 3, 193, 40, 238, 25, 3, 10, 98, 38, 81, 123, 128,
    90, 7, 37, 18, 165, 228, 205, 39, 75, 127, 209, 250, 35, 248, 48, 5, 130, 8, 255, 26,
    6, 59, 65, 3, 156, 116, 3, 107, 91, 61, 168, 177, 160, 185, 49, 53, 167, 16, 53, 45,
    160, 246, 195, 18, 3, 160, 157, 31, 35, 41, 101, 27, 179, 171, 22, 3, 27, 1, 0, 23,
    32, 163, 60, 194, 104, 248, 118, 79, 231, 250, 185, 231, 57, 251, 224, 71, 220, 255, 220, 10,
    60, 77, 181, 242, 75, 60, 83, 208, 210, 38, 129, 234, 198}));

  for(int i = 0; i < 1000000; i ++){
    decode(buf);
  }
  return 0;
}
