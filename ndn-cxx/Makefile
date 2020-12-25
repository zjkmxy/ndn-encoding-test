CFLAGS = -Wall -O3 `pkg-config --cflags --libs libndn-cxx`
CFLAGS_DBG = -Wall -ggdb -O0 `pkg-config --cflags --libs libndn-cxx`

all: main

main: main.cpp
	g++ -std=c++14 $< -o $@ $(CFLAGS)

clean:
	rm -rf *.dSYM
	rm -rf main

run:
	time ./main
