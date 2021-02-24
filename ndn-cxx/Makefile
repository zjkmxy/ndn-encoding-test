CFLAGS = -Wall -O3 `pkg-config --cflags --libs libndn-cxx`
CFLAGS_DBG = -Wall -ggdb -O0 `pkg-config --cflags --libs libndn-cxx`

all: encode-01 encode-02 encode-03 decode-01 decode-02 decode-03

encode-01: encode-01.cpp
	g++ -std=c++14 $< -o $@ $(CFLAGS)

encode-02: encode-02.cpp
	g++ -std=c++14 $< -o $@ $(CFLAGS)

encode-03: encode-03.cpp
	g++ -std=c++14 $< -o $@ $(CFLAGS)

decode-01: decode-01.cpp
	g++ -std=c++14 $< -o $@ $(CFLAGS)

decode-02: decode-02.cpp
	g++ -std=c++14 $< -o $@ $(CFLAGS)

decode-03: decode-03.cpp
	g++ -std=c++14 $< -o $@ $(CFLAGS)

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
