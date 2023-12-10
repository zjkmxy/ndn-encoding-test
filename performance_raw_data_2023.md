Performance Raw Data
====================

Date
----

2023-12-09

Environments
------------

- CPU: AMD EPYC 7702P 64-Core Processor 2GHz
- Memory: 256 GiB
- Operating System: Ubuntu 22.04
- Compilers
  - gcc 11.3.0
  - go 1.21.5
  - Python 3.12.0
  - Pypy 3.10 v7.3.13
  - Deno 1.38.5
- Method: 5 times, take average. Encode does not include signing.
- Library Versions
  - ndn-cxx: [0.8.1](https://github.com/named-data/ndn-cxx/commit/81de5d98e92c30dfb7827f8143080b379dc2a51d)
  - go-ndn: [v0.0.6](https://github.com/zjkmxy/go-ndn/commit/2e42d595c02b3b60ec55893099b885bcbd546101)
  - python-ndn: [v0.4.1](https://github.com/named-data/python-ndn/commit/7b9ab6ed23d1260f6592d4529fc15c8121456a64)
  - NDNts: 0.0.20231126-nightly-2870797

Raw Data - Go
-------------

### T1
```text
========== ENCODING ==========
#1: 1000000 x 100B
reflection:     3.188163318s
codegen:        445.350274ms
go-ndn:         1.149100919s
protobuf:       1.128982309s
json:   1.754602141s
=== Total Bytes: 157000000 ===

#2: 1000000 x 4000B
reflection:     5.123540072s
codegen:        1.20744011s
go-ndn:         1.074130352s
protobuf:       2.118787161s
json:   5.659988608s
=== Total Bytes: 4061000000 ===

#3: 1000000 x 10B, very long name
reflection:     4.306128037s
codegen:        2.633603477s
go-ndn:         3.857531932s
protobuf:       3.385729225s
json:   4.31562892s
=== Total Bytes: 157000000 ===


========== DECODING ==========
#1: 1000000 x 100B
reflection:     3.34773905s
codegen:        464.38013ms
go-ndn:         1.753672847s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
reflection:     3.457089312s
codegen:        456.395132ms
go-ndn:         2.055107058s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
reflection:     5.499433051s
codegen:        1.943487774s
go-ndn:         3.691994261s
=== Total Bytes: 189000000 ===
```

### T2
```text
========== ENCODING ==========
#1: 1000000 x 100B
reflection:     3.162634277s
codegen:        456.354049ms
go-ndn:         1.198807459s
protobuf:       1.162460412s
json:   1.727906004s
=== Total Bytes: 157000000 ===

#2: 1000000 x 4000B
reflection:     5.142724516s
codegen:        1.166518881s
go-ndn:         1.141949464s
protobuf:       2.156521597s
json:   5.6505255s
=== Total Bytes: 4061000000 ===

#3: 1000000 x 10B, very long name
reflection:     5.358019072s
codegen:        2.088750306s
go-ndn:         3.933235018s
protobuf:       3.426065535s
json:   4.439639507s
=== Total Bytes: 157000000 ===


========== DECODING ==========
#1: 1000000 x 100B
reflection:     3.383600445s
codegen:        427.662595ms
go-ndn:         1.766459813s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
reflection:     3.427622616s
codegen:        449.577941ms
go-ndn:         2.05149809s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
reflection:     5.345699627s
codegen:        1.920306255s
go-ndn:         3.776472857s
=== Total Bytes: 189000000 ===
```

### T3
```text
========== ENCODING ==========
#1: 1000000 x 100B
reflection:     3.166665524s
codegen:        394.87695ms
go-ndn:         1.166050806s
protobuf:       1.155040429s
json:   1.799294816s
=== Total Bytes: 157000000 ===

#2: 1000000 x 4000B
reflection:     5.114050966s
codegen:        1.281657157s
go-ndn:         1.082722699s
protobuf:       2.152142947s
json:   5.729377411s
=== Total Bytes: 4061000000 ===

#3: 1000000 x 10B, very long name
reflection:     4.290383556s
codegen:        2.587040378s
go-ndn:         3.901498515s
protobuf:       3.392390111s
json:   4.419956046s
=== Total Bytes: 157000000 ===


========== DECODING ==========
#1: 1000000 x 100B
reflection:     3.415045548s
codegen:        431.113311ms
go-ndn:         1.812330514s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
reflection:     3.44277798s
codegen:        456.715862ms
go-ndn:         2.027367517s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
reflection:     5.604319083s
codegen:        1.888457719s
go-ndn:         3.779904408s
=== Total Bytes: 189000000 ===
```

### T4
```text
========== ENCODING ==========
#1: 1000000 x 100B
reflection:     3.245529237s
codegen:        400.359004ms
go-ndn:         1.207028271s
protobuf:       1.161596218s
json:   1.803668145s
=== Total Bytes: 157000000 ===

#2: 1000000 x 4000B
reflection:     5.11098378s
codegen:        1.233512597s
go-ndn:         1.101945905s
protobuf:       2.185063836s
json:   5.701998554s
=== Total Bytes: 4061000000 ===

#3: 1000000 x 10B, very long name
reflection:     4.304546047s
codegen:        2.429778034s
go-ndn:         3.993229918s
protobuf:       3.372845963s
json:   4.431327542s
=== Total Bytes: 157000000 ===


========== DECODING ==========
#1: 1000000 x 100B
reflection:     3.344392896s
codegen:        417.33244ms
go-ndn:         1.819887811s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
reflection:     3.451446055s
codegen:        454.970694ms
go-ndn:         2.031603966s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
reflection:     5.19626933s
codegen:        1.876178557s
go-ndn:         3.774154374s
=== Total Bytes: 189000000 ===
```

### T5
```text
========== ENCODING ==========
#1: 1000000 x 100B
reflection:     3.207666438s
codegen:        415.052414ms
go-ndn:         1.192679959s
protobuf:       1.165053077s
json:   1.80058587s
=== Total Bytes: 157000000 ===

#2: 1000000 x 4000B
reflection:     5.137068992s
codegen:        1.155407953s
go-ndn:         1.098535324s
protobuf:       2.173579792s
json:   5.635312819s
=== Total Bytes: 4061000000 ===

#3: 1000000 x 10B, very long name
reflection:     5.075902603s
codegen:        2.079120932s
go-ndn:         3.88834131s
protobuf:       3.336134292s
json:   4.322799631s
=== Total Bytes: 157000000 ===


========== DECODING ==========
#1: 1000000 x 100B
reflection:     3.340894137s
codegen:        447.571516ms
go-ndn:         1.844968752s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
reflection:     3.456192094s
codegen:        448.691409ms
go-ndn:         2.084533663s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
reflection:     5.522068235s
codegen:        1.833785383s
go-ndn:         3.79991212s
=== Total Bytes: 189000000 ===
```

Raw Data - ndn-cxx
------------------

### T1
```text
time ./encode-01
4.64user 0.00system 0:04.66elapsed 99%CPU (0avgtext+0avgdata 11000maxresident)k
64inputs+0outputs (0major+468minor)pagefaults 0swaps
time ./encode-02
4.94user 0.00system 0:04.94elapsed 99%CPU (0avgtext+0avgdata 11236maxresident)k
0inputs+0outputs (0major+477minor)pagefaults 0swaps
time ./encode-03
24.53user 0.00system 0:24.53elapsed 99%CPU (0avgtext+0avgdata 11040maxresident)k
0inputs+0outputs (0major+473minor)pagefaults 0swaps
time ./decode-01
1.01user 0.00system 0:01.01elapsed 100%CPU (0avgtext+0avgdata 11096maxresident)k
0inputs+0outputs (0major+471minor)pagefaults 0swaps
time ./decode-02
1.00user 0.00system 0:01.00elapsed 100%CPU (0avgtext+0avgdata 10656maxresident)k
0inputs+0outputs (0major+464minor)pagefaults 0swaps
time ./decode-03
2.12user 0.00system 0:02.12elapsed 99%CPU (0avgtext+0avgdata 10744maxresident)k
0inputs+0outputs (0major+468minor)pagefaults 0swaps
```

### T2
```text
time ./encode-01
4.64user 0.00system 0:04.64elapsed 99%CPU (0avgtext+0avgdata 11028maxresident)k
0inputs+0outputs (0major+469minor)pagefaults 0swaps
time ./encode-02
4.93user 0.00system 0:04.93elapsed 99%CPU (0avgtext+0avgdata 11236maxresident)k
0inputs+0outputs (0major+475minor)pagefaults 0swaps
time ./encode-03
24.79user 0.00system 0:24.80elapsed 99%CPU (0avgtext+0avgdata 10844maxresident)k
0inputs+0outputs (0major+468minor)pagefaults 0swaps
time ./decode-01
0.98user 0.00system 0:00.98elapsed 99%CPU (0avgtext+0avgdata 10936maxresident)k
0inputs+0outputs (0major+466minor)pagefaults 0swaps
time ./decode-02
1.01user 0.00system 0:01.01elapsed 99%CPU (0avgtext+0avgdata 10656maxresident)k
0inputs+0outputs (0major+465minor)pagefaults 0swaps
time ./decode-03
2.13user 0.00system 0:02.13elapsed 100%CPU (0avgtext+0avgdata 10748maxresident)k
0inputs+0outputs (0major+470minor)pagefaults 0swaps
```

### T3
```text
time ./encode-01
4.64user 0.00system 0:04.65elapsed 99%CPU (0avgtext+0avgdata 10728maxresident)k
0inputs+0outputs (0major+467minor)pagefaults 0swaps
time ./encode-02
4.86user 0.00system 0:04.87elapsed 99%CPU (0avgtext+0avgdata 10660maxresident)k
0inputs+0outputs (0major+467minor)pagefaults 0swaps
time ./encode-03
25.09user 0.00system 0:25.09elapsed 99%CPU (0avgtext+0avgdata 11232maxresident)k
0inputs+0outputs (0major+475minor)pagefaults 0swaps
time ./decode-01
1.00user 0.00system 0:01.00elapsed 100%CPU (0avgtext+0avgdata 10588maxresident)k
0inputs+0outputs (0major+462minor)pagefaults 0swaps
time ./decode-02
1.01user 0.00system 0:01.01elapsed 100%CPU (0avgtext+0avgdata 10740maxresident)k
0inputs+0outputs (0major+470minor)pagefaults 0swaps
time ./decode-03
2.13user 0.00system 0:02.13elapsed 99%CPU (0avgtext+0avgdata 11108maxresident)k
0inputs+0outputs (0major+474minor)pagefaults 0swaps
```

### T4
```text
time ./encode-01
4.67user 0.00system 0:04.67elapsed 99%CPU (0avgtext+0avgdata 10760maxresident)k
0inputs+0outputs (0major+467minor)pagefaults 0swaps
time ./encode-02
4.78user 0.00system 0:04.79elapsed 99%CPU (0avgtext+0avgdata 10740maxresident)k
0inputs+0outputs (0major+468minor)pagefaults 0swaps
time ./encode-03
24.50user 0.00system 0:24.50elapsed 99%CPU (0avgtext+0avgdata 11012maxresident)k
0inputs+0outputs (0major+473minor)pagefaults 0swaps
time ./decode-01
1.01user 0.00system 0:01.01elapsed 99%CPU (0avgtext+0avgdata 10720maxresident)k
0inputs+0outputs (0major+468minor)pagefaults 0swaps
time ./decode-02
0.98user 0.00system 0:00.99elapsed 100%CPU (0avgtext+0avgdata 10696maxresident)k
0inputs+0outputs (0major+465minor)pagefaults 0swaps
time ./decode-03
2.13user 0.00system 0:02.13elapsed 99%CPU (0avgtext+0avgdata 10808maxresident)k
0inputs+0outputs (0major+470minor)pagefaults 0swaps
```

### T5
```text
time ./encode-01
4.61user 0.00system 0:04.61elapsed 100%CPU (0avgtext+0avgdata 11012maxresident)k
0inputs+0outputs (0major+469minor)pagefaults 0swaps
time ./encode-02
4.83user 0.00system 0:04.83elapsed 99%CPU (0avgtext+0avgdata 10716maxresident)k
0inputs+0outputs (0major+468minor)pagefaults 0swaps
time ./encode-03
24.55user 0.00system 0:24.56elapsed 99%CPU (0avgtext+0avgdata 10620maxresident)k
0inputs+0outputs (0major+465minor)pagefaults 0swaps
time ./decode-01
0.98user 0.00system 0:00.98elapsed 100%CPU (0avgtext+0avgdata 10688maxresident)k
0inputs+0outputs (0major+463minor)pagefaults 0swaps
time ./decode-02
1.01user 0.00system 0:01.01elapsed 100%CPU (0avgtext+0avgdata 10912maxresident)k
0inputs+0outputs (0major+469minor)pagefaults 0swaps
time ./decode-03
2.09user 0.00system 0:02.10elapsed 100%CPU (0avgtext+0avgdata 10648maxresident)k
0inputs+0outputs (0major+463minor)pagefaults 0swaps
```

Raw Data - cpp-template
-----------------------

### T1
```text
time ./encode-01
0.64user 0.00system 0:00.64elapsed 100%CPU (0avgtext+0avgdata 4536maxresident)k
0inputs+0outputs (0major+243minor)pagefaults 0swaps
time ./encode-02
0.84user 0.00system 0:00.84elapsed 100%CPU (0avgtext+0avgdata 4608maxresident)k
0inputs+0outputs (0major+243minor)pagefaults 0swaps
time ./encode-03
3.97user 0.00system 0:03.97elapsed 99%CPU (0avgtext+0avgdata 4484maxresident)k
0inputs+0outputs (0major+240minor)pagefaults 0swaps
time ./decode-01
0.43user 0.00system 0:00.43elapsed 100%CPU (0avgtext+0avgdata 2760maxresident)k
0inputs+0outputs (0major+122minor)pagefaults 0swaps
time ./decode-02
0.78user 0.00system 0:00.78elapsed 100%CPU (0avgtext+0avgdata 2952maxresident)k
0inputs+0outputs (0major+126minor)pagefaults 0swaps
time ./decode-03
1.49user 0.00system 0:01.49elapsed 100%CPU (0avgtext+0avgdata 2872maxresident)k
0inputs+0outputs (0major+125minor)pagefaults 0swaps
```

### T2
```text
time ./encode-01
0.64user 0.00system 0:00.64elapsed 100%CPU (0avgtext+0avgdata 4592maxresident)k
0inputs+0outputs (0major+243minor)pagefaults 0swaps
time ./encode-02
0.85user 0.00system 0:00.86elapsed 99%CPU (0avgtext+0avgdata 4488maxresident)k
0inputs+0outputs (0major+242minor)pagefaults 0swaps
time ./encode-03
3.94user 0.00system 0:03.94elapsed 99%CPU (0avgtext+0avgdata 4596maxresident)k
0inputs+0outputs (0major+242minor)pagefaults 0swaps
time ./decode-01
0.44user 0.00system 0:00.44elapsed 100%CPU (0avgtext+0avgdata 2948maxresident)k
0inputs+0outputs (0major+126minor)pagefaults 0swaps
time ./decode-02
0.76user 0.00system 0:00.76elapsed 99%CPU (0avgtext+0avgdata 2952maxresident)k
0inputs+0outputs (0major+125minor)pagefaults 0swaps
time ./decode-03
1.49user 0.00system 0:01.49elapsed 100%CPU (0avgtext+0avgdata 2764maxresident)k
0inputs+0outputs (0major+123minor)pagefaults 0swaps
```

### T3
```text
time ./encode-01
0.64user 0.00system 0:00.64elapsed 100%CPU (0avgtext+0avgdata 4596maxresident)k
0inputs+0outputs (0major+242minor)pagefaults 0swaps
time ./encode-02
0.86user 0.00system 0:00.86elapsed 100%CPU (0avgtext+0avgdata 4548maxresident)k
0inputs+0outputs (0major+245minor)pagefaults 0swaps
time ./encode-03
3.94user 0.00system 0:03.94elapsed 100%CPU (0avgtext+0avgdata 4596maxresident)k
0inputs+0outputs (0major+242minor)pagefaults 0swaps
time ./decode-01
0.44user 0.00system 0:00.44elapsed 100%CPU (0avgtext+0avgdata 2756maxresident)k
0inputs+0outputs (0major+123minor)pagefaults 0swaps
time ./decode-02
0.78user 0.00system 0:00.78elapsed 99%CPU (0avgtext+0avgdata 2960maxresident)k
0inputs+0outputs (0major+125minor)pagefaults 0swaps
time ./decode-03
1.49user 0.00system 0:01.49elapsed 100%CPU (0avgtext+0avgdata 2764maxresident)k
0inputs+0outputs (0major+124minor)pagefaults 0swaps
```

### T4
```text
time ./encode-01
0.64user 0.00system 0:00.64elapsed 100%CPU (0avgtext+0avgdata 4676maxresident)k
0inputs+0outputs (0major+240minor)pagefaults 0swaps
time ./encode-02
0.84user 0.00system 0:00.85elapsed 99%CPU (0avgtext+0avgdata 4604maxresident)k
0inputs+0outputs (0major+244minor)pagefaults 0swaps
time ./encode-03
3.97user 0.00system 0:03.97elapsed 100%CPU (0avgtext+0avgdata 4596maxresident)k
0inputs+0outputs (0major+242minor)pagefaults 0swaps
time ./decode-01
0.44user 0.00system 0:00.44elapsed 100%CPU (0avgtext+0avgdata 2764maxresident)k
0inputs+0outputs (0major+122minor)pagefaults 0swaps
time ./decode-02
0.77user 0.00system 0:00.78elapsed 99%CPU (0avgtext+0avgdata 2792maxresident)k
0inputs+0outputs (0major+125minor)pagefaults 0swaps
time ./decode-03
1.47user 0.00system 0:01.48elapsed 99%CPU (0avgtext+0avgdata 2952maxresident)k
0inputs+0outputs (0major+125minor)pagefaults 0swaps
```

### T5
```text
time ./encode-01
0.63user 0.00system 0:00.63elapsed 100%CPU (0avgtext+0avgdata 4608maxresident)k
0inputs+0outputs (0major+244minor)pagefaults 0swaps
time ./encode-02
0.84user 0.00system 0:00.84elapsed 100%CPU (0avgtext+0avgdata 4544maxresident)k
0inputs+0outputs (0major+244minor)pagefaults 0swaps
time ./encode-03
3.91user 0.00system 0:03.92elapsed 99%CPU (0avgtext+0avgdata 4604maxresident)k
0inputs+0outputs (0major+243minor)pagefaults 0swaps
time ./decode-01
0.43user 0.00system 0:00.43elapsed 100%CPU (0avgtext+0avgdata 2780maxresident)k
0inputs+0outputs (0major+123minor)pagefaults 0swaps
time ./decode-02
0.74user 0.00system 0:00.75elapsed 99%CPU (0avgtext+0avgdata 2776maxresident)k
0inputs+0outputs (0major+124minor)pagefaults 0swaps
time ./decode-03
1.48user 0.00system 0:01.48elapsed 100%CPU (0avgtext+0avgdata 2940maxresident)k
0inputs+0outputs (0major+124minor)pagefaults 0swaps
```

Raw Data - python-ndn
---------------------

### T1
```text
python encode-01.py && python encode-02.py && python encode-03.py && python decode-01.py && python decode-02.py && python decode-03.py
Total time: 48.2844 seconds
Total time: 50.3509 seconds
Total time: 119.1322 seconds
Total time: 22.3633 seconds
Total time: 24.2563 seconds
Total time: 33.6972 seconds
```

### T2
```text
python encode-01.py && python encode-02.py && python encode-03.py && python decode-01.py && python decode-02.py && python decode-03.py
Total time: 48.7112 seconds
Total time: 50.4541 seconds
Total time: 118.2867 seconds
Total time: 22.2933 seconds
Total time: 24.7141 seconds
Total time: 33.9607 seconds
```

### T3
```text
python encode-01.py && python encode-02.py && python encode-03.py && python decode-01.py && python decode-02.py && python decode-03.py
Total time: 48.4792 seconds
Total time: 50.0068 seconds
Total time: 119.9308 seconds
Total time: 23.2791 seconds
Total time: 25.0208 seconds
Total time: 33.2429 seconds
```

### T4
```text
python encode-01.py && python encode-02.py && python encode-03.py && python decode-01.py && python decode-02.py && python decode-03.py
Total time: 48.2575 seconds
Total time: 50.0900 seconds
Total time: 121.7798 seconds
Total time: 22.8126 seconds
Total time: 24.8047 seconds
Total time: 34.0893 seconds
```

### T5
```text
python encode-01.py && python encode-02.py && python encode-03.py && python decode-01.py && python decode-02.py && python decode-03.py
Total time: 48.0565 seconds
Total time: 50.5088 seconds
Total time: 119.1157 seconds
Total time: 22.3893 seconds
Total time: 24.2955 seconds
Total time: 33.3656 seconds
```

Raw Data - python-ndn (pypy3)
---------------------

### T1
```text
pypy3.10 encode-01.py && pypy3.10 encode-02.py && pypy3.10 encode-03.py && pypy3.10 decode-01.py && pypy3.10 decode-02.py && pypy3.10 decode-03.py
Total time: 13.8104 seconds
Total time: 16.3660 seconds
Total time: 34.8110 seconds
Total time: 4.6478 seconds
Total time: 4.4920 seconds
Total time: 4.8683 seconds
```

### T2
```text
Total time: 13.7078 seconds
Total time: 16.3718 seconds
Total time: 34.8080 seconds
Total time: 4.4539 seconds
Total time: 5.2632 seconds
Total time: 4.8263 seconds
```

### T3
```text
Total time: 13.7765 seconds
Total time: 16.3752 seconds
Total time: 34.8798 seconds
Total time: 4.4436 seconds
Total time: 5.2494 seconds
Total time: 4.8332 seconds
```

### T4
```text
Total time: 13.6655 seconds
Total time: 16.7786 seconds
Total time: 36.2705 seconds
Total time: 4.5095 seconds
Total time: 4.5166 seconds
Total time: 4.8824 seconds
```

### T5
```text
Total time: 13.8712 seconds
Total time: 16.3159 seconds
Total time: 37.5697 seconds
Total time: 4.5002 seconds
Total time: 4.6258 seconds
Total time: 5.0124 seconds
```

Raw Data - NDNts (Deno)
---------------------

### T1
```text
deno task encode-01 && deno task encode-02 && deno task encode-03 && deno task decode-01 && deno task decode-02 && deno task decode-03
Task encode-01 deno run --unstable encode-01.ts
Total time: 49.802 seconds
Task encode-02 deno run --unstable encode-02.ts
Total time: 52.882 seconds
Task encode-03 deno run --unstable encode-03.ts
Total time: 86.482 seconds
Task decode-01 deno run --unstable decode-01.ts
Total time: 12.352 seconds
Task decode-02 deno run --unstable decode-02.ts
Total time: 4.78 seconds
Task decode-03 deno run --unstable decode-03.ts
Total time: 13.392 seconds
```

### T2
```text
Task encode-01 deno run --unstable encode-01.ts
Total time: 50.462 seconds
Task encode-02 deno run --unstable encode-02.ts
Total time: 54.42 seconds
Task encode-03 deno run --unstable encode-03.ts
Total time: 85.134 seconds
Task decode-01 deno run --unstable decode-01.ts
Total time: 13.518 seconds
Task decode-02 deno run --unstable decode-02.ts
Total time: 4.768 seconds
Task decode-03 deno run --unstable decode-03.ts
Total time: 13.484 seconds
```

### T3
```text
Task encode-01 deno run --unstable encode-01.ts
Total time: 51.018 seconds
Task encode-02 deno run --unstable encode-02.ts
Total time: 55.424 seconds
Task encode-03 deno run --unstable encode-03.ts
Total time: 88.404 seconds
Task decode-01 deno run --unstable decode-01.ts
Total time: 13.158 seconds
Task decode-02 deno run --unstable decode-02.ts
Total time: 4.778 seconds
Task decode-03 deno run --unstable decode-03.ts
Total time: 13.316 seconds
```

### T4
```text
Task encode-01 deno run --unstable encode-01.ts
Total time: 50.894 seconds
Task encode-02 deno run --unstable encode-02.ts
Total time: 54.242 seconds
Task encode-03 deno run --unstable encode-03.ts
Total time: 88.614 seconds
Task decode-01 deno run --unstable decode-01.ts
Total time: 12.422 seconds
Task decode-02 deno run --unstable decode-02.ts
Total time: 4.764 seconds
Task decode-03 deno run --unstable decode-03.ts
Total time: 13.232 seconds
```

### T5
```text
Task encode-01 deno run --unstable encode-01.ts
Total time: 50.828 seconds
Task encode-02 deno run --unstable encode-02.ts
Total time: 53.402 seconds
Task encode-03 deno run --unstable encode-03.ts
Total time: 88.178 seconds
Task decode-01 deno run --unstable decode-01.ts
Total time: 13.636 seconds
Task decode-02 deno run --unstable decode-02.ts
Total time: 5.008 seconds
Task decode-03 deno run --unstable decode-03.ts
Total time: 11.888 seconds
```
