CFLAGS = `pkg-config --cflags --libs libndn-cxx`

all: main

main: main.cpp
	g++ -std=c++14 $< -o $@ -Wall -ggdb -O0 $(CFLAGS)

clean:
	rm -rf *.dSYM
	rm -rf main

run:
	time ./main
