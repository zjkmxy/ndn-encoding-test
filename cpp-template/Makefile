CFLAGS = -Wall -O3
CFLAGS_DBG = -Wall -ggdb -O0

all: main

main: main.cpp sha256/sha256.cpp sha256/sha256.h ndn-format-0-3.hpp tlv-encoder.hpp
	g++ -std=c++17 $< sha256/sha256.cpp -o $@ $(CFLAGS)

clean:
	rm -rf *.dSYM
	rm -rf main

run:
	time ./main
