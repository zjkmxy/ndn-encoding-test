CFLAGS = -Wall -O3
CFLAGS_DBG = -Wall -ggdb -O0

all: encode-01 encode-02 encode-03 decode-01 decode-02 decode-03

encode-01: encode-01.cpp sha256/sha256.cpp sha256/sha256.h ndn-format-0-3.hpp tlv-encoder.hpp slice.hpp
	g++-10 -std=c++20 $< sha256/sha256.cpp -o $@ $(CFLAGS)

encode-02: encode-02.cpp sha256/sha256.cpp sha256/sha256.h ndn-format-0-3.hpp tlv-encoder.hpp slice.hpp
	g++-10 -std=c++20 $< sha256/sha256.cpp -o $@ $(CFLAGS)

encode-03: encode-03.cpp sha256/sha256.cpp sha256/sha256.h ndn-format-0-3.hpp tlv-encoder.hpp slice.hpp
	g++-10 -std=c++20 $< sha256/sha256.cpp -o $@ $(CFLAGS)

decode-01: decode-01.cpp sha256/sha256.cpp sha256/sha256.h ndn-format-0-3.hpp tlv-encoder.hpp slice.hpp
	g++-10 -std=c++20 $< sha256/sha256.cpp -o $@ $(CFLAGS)

decode-02: decode-02.cpp sha256/sha256.cpp sha256/sha256.h ndn-format-0-3.hpp tlv-encoder.hpp slice.hpp
	g++-10 -std=c++20 $< sha256/sha256.cpp -o $@ $(CFLAGS)

decode-03: decode-03.cpp sha256/sha256.cpp sha256/sha256.h ndn-format-0-3.hpp tlv-encoder.hpp slice.hpp
	g++-10 -std=c++20 $< sha256/sha256.cpp -o $@ $(CFLAGS)

clean:
	rm -rf *.dSYM
	rm -rf encode-01 encode-02 encode-03 decode-01 decode-02 decode-03

run:
	time ./encode-01
	time ./encode-02
	time ./encode-03
	time ./decode-01
	time ./decode-02
	time ./decode-03
