CFLAGS = -Wall -O3 `pkg-config --cflags --libs libndn-cxx`
CFLAGS_DBG = -Wall -ggdb -O0 `pkg-config --cflags --libs libndn-cxx`
# EXTRA_CPP_INPUT = ../memory_monitor/memory_monitor.cpp
EXTRA_CPP_INPUT =

all: encode-01 encode-02 encode-03 decode-01 decode-02 decode-03

encode-01: encode-01.cpp
	g++ -std=c++17 $< $(EXTRA_CPP_INPUT) -o $@ $(CFLAGS)

encode-02: encode-02.cpp
	g++ -std=c++17 $< $(EXTRA_CPP_INPUT) -o $@ $(CFLAGS)

encode-03: encode-03.cpp
	g++ -std=c++17 $< $(EXTRA_CPP_INPUT) -o $@ $(CFLAGS)

decode-01: decode-01.cpp
	g++ -std=c++17 $< $(EXTRA_CPP_INPUT) -o $@ $(CFLAGS)

decode-02: decode-02.cpp
	g++ -std=c++17 $< $(EXTRA_CPP_INPUT) -o $@ $(CFLAGS)

decode-03: decode-03.cpp
	g++ -std=c++17 $< $(EXTRA_CPP_INPUT) -o $@ $(CFLAGS)

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
