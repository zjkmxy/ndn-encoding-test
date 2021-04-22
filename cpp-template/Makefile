COMPILER = g++-10
# COMPILER = /usr/local/opt/llvm/bin/clang++
INCL_PATHS = -I/usr/local/opt/openssl@1.1/include -L/usr/local/opt/openssl@1.1/lib
CFLAGS = -Wall -O3 -lcrypto -lssl
CFLAGS_DBG = -Wall -ggdb -O0 -lcrypto -lssl

all: encode-01 encode-02 encode-03 decode-01 decode-02 decode-03

encode-01: encode-01.cpp ndn-format-0-3.hpp tlv-encoder.hpp slice.hpp
	$(COMPILER) -std=c++20 $(INCL_PATHS) $< -o $@ $(CFLAGS)

encode-02: encode-02.cpp ndn-format-0-3.hpp tlv-encoder.hpp slice.hpp
	$(COMPILER) -std=c++20 $(INCL_PATHS) $< -o $@ $(CFLAGS)

encode-03: encode-03.cpp ndn-format-0-3.hpp tlv-encoder.hpp slice.hpp
	$(COMPILER) -std=c++20 $(INCL_PATHS) $< -o $@ $(CFLAGS)

decode-01: decode-01.cpp ndn-format-0-3.hpp tlv-encoder.hpp slice.hpp
	$(COMPILER) -std=c++20 $(INCL_PATHS) $< -o $@ $(CFLAGS)

decode-02: decode-02.cpp ndn-format-0-3.hpp tlv-encoder.hpp slice.hpp
	$(COMPILER) -std=c++20 $(INCL_PATHS) $< -o $@ $(CFLAGS)

decode-03: decode-03.cpp ndn-format-0-3.hpp tlv-encoder.hpp slice.hpp
	$(COMPILER) -std=c++20 $(INCL_PATHS) $< -o $@ $(CFLAGS)

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
