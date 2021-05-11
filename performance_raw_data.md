Performance Raw Data
====================

Raw Data - Go
-------------

### T1
```text
========== ENCODEING ==========
#1: 1000000 x 100B
block: 		4.854765576s
reflection: 	3.807049976s
codegen: 	1.141233216s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		19.546341029s
reflection: 	14.431212399s
codegen: 	11.844070864s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		15.125524913s
reflection: 	5.890830713s
codegen: 	3.12920471s
=== Total Bytes: 189000000 ===


========== DECODING ==========
#1: 1000000 x 100B
block: 		2.50055776s
reflection: 	3.291573173s
codegen: 	460.77955ms
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		2.323502626s
reflection: 	3.243882034s
codegen: 	556.047842ms
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		8.598232285s
reflection: 	4.987587884s
codegen: 	1.969024137s
=== Total Bytes: 189000000 ===
```

### T2
```text
========== ENCODEING ==========
#1: 1000000 x 100B
block: 		4.936672426s
reflection: 	3.789766454s
codegen: 	1.12385554s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		19.529038133s
reflection: 	14.392564416s
codegen: 	11.915797907s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		15.321592914s
reflection: 	5.797209538s
codegen: 	3.121677155s
=== Total Bytes: 189000000 ===


========== DECODING ==========
#1: 1000000 x 100B
block: 		2.470289829s
reflection: 	3.236745777s
codegen: 	470.479219ms
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		2.722743899s
reflection: 	3.251319296s
codegen: 	562.28719ms
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		8.422697879s
reflection: 	4.942703808s
codegen: 	1.908739977s
=== Total Bytes: 189000000 ===
```

### T3
```text
========== ENCODEING ==========
#1: 1000000 x 100B
block: 		4.929584319s
reflection: 	3.786709561s
codegen: 	1.138760976s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		19.523884387s
reflection: 	14.48764612s
codegen: 	11.866381621s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		15.261806736s
reflection: 	5.773296797s
codegen: 	3.127485494s
=== Total Bytes: 189000000 ===


========== DECODING ==========
#1: 1000000 x 100B
block: 		2.430580519s
reflection: 	3.27619214s
codegen: 	477.900987ms
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		2.752812124s
reflection: 	3.237357446s
codegen: 	566.837059ms
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		8.323488166s
reflection: 	4.982947373s
codegen: 	1.903607742s
=== Total Bytes: 189000000 ===
```

### T4
```text
========== ENCODEING ==========
#1: 1000000 x 100B
block: 		4.882205334s
reflection: 	3.813996769s
codegen: 	1.093506218s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		19.427554844s
reflection: 	14.403119278s
codegen: 	11.9841946s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		14.891972644s
reflection: 	5.832692361s
codegen: 	3.126838255s
=== Total Bytes: 189000000 ===


========== DECODING ==========
#1: 1000000 x 100B
block: 		2.439436392s
reflection: 	3.277122611s
codegen: 	481.221132ms
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		2.567456601s
reflection: 	3.246431148s
codegen: 	557.789032ms
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		8.589623415s
reflection: 	5.023395564s
codegen: 	1.915830726s
=== Total Bytes: 189000000 ===
```

### T5
```text
========== ENCODEING ==========
#1: 1000000 x 100B
block: 		4.900603528s
reflection: 	3.792944088s
codegen: 	1.14173714s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		19.501905263s
reflection: 	14.420068126s
codegen: 	11.856724463s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		15.162291814s
reflection: 	5.850854205s
codegen: 	3.204544324s
=== Total Bytes: 189000000 ===


========== DECODING ==========
#1: 1000000 x 100B
block: 		2.435503377s
reflection: 	3.285752864s
codegen: 	493.424117ms
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		2.697374808s
reflection: 	3.245891018s
codegen: 	557.266169ms
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		8.020276705s
reflection: 	4.995550971s
codegen: 	1.968909545s
=== Total Bytes: 189000000 ===
```

Raw Data - ndn-cxx
------------------

### T1
```text
time ./encode-01
6.58user 0.00system 0:06.60elapsed 99%CPU (0avgtext+0avgdata 11628maxresident)k
1680inputs+0outputs (10major+467minor)pagefaults 0swaps
time ./encode-02
15.24user 0.00system 0:15.24elapsed 99%CPU (0avgtext+0avgdata 11992maxresident)k
0inputs+0outputs (0major+474minor)pagefaults 0swaps
time ./encode-03
26.19user 0.00system 0:26.20elapsed 99%CPU (0avgtext+0avgdata 11640maxresident)k
0inputs+0outputs (0major+474minor)pagefaults 0swaps
time ./decode-01
1.13user 0.00system 0:01.13elapsed 100%CPU (0avgtext+0avgdata 11500maxresident)k
0inputs+0outputs (0major+461minor)pagefaults 0swaps
time ./decode-02
1.13user 0.00system 0:01.14elapsed 99%CPU (0avgtext+0avgdata 11144maxresident)k
0inputs+0outputs (0major+458minor)pagefaults 0swaps
time ./decode-03
2.72user 0.00system 0:02.72elapsed 100%CPU (0avgtext+0avgdata 11228maxresident)k
0inputs+0outputs (0major+460minor)pagefaults 0swaps
```

### T2
```text
time ./encode-01
6.36user 0.00system 0:06.36elapsed 99%CPU (0avgtext+0avgdata 11664maxresident)k
0inputs+0outputs (0major+471minor)pagefaults 0swaps
time ./encode-02
15.34user 0.00system 0:15.34elapsed 99%CPU (0avgtext+0avgdata 11632maxresident)k
0inputs+0outputs (0major+472minor)pagefaults 0swaps
time ./encode-03
25.80user 0.00system 0:25.80elapsed 99%CPU (0avgtext+0avgdata 11628maxresident)k
0inputs+0outputs (0major+470minor)pagefaults 0swaps
time ./decode-01
1.11user 0.00system 0:01.12elapsed 100%CPU (0avgtext+0avgdata 11224maxresident)k
0inputs+0outputs (0major+460minor)pagefaults 0swaps
time ./decode-02
1.11user 0.00system 0:01.11elapsed 100%CPU (0avgtext+0avgdata 11364maxresident)k
0inputs+0outputs (0major+460minor)pagefaults 0swaps
time ./decode-03
2.68user 0.00system 0:02.68elapsed 100%CPU (0avgtext+0avgdata 11388maxresident)k
0inputs+0outputs (0major+462minor)pagefaults 0swaps
```

### T3
```text
time ./encode-01
6.61user 0.00system 0:06.62elapsed 99%CPU (0avgtext+0avgdata 12052maxresident)k
0inputs+0outputs (0major+475minor)pagefaults 0swaps
time ./encode-02
15.33user 0.00system 0:15.34elapsed 99%CPU (0avgtext+0avgdata 12020maxresident)k
0inputs+0outputs (0major+474minor)pagefaults 0swaps
time ./encode-03
26.11user 0.00system 0:26.11elapsed 99%CPU (0avgtext+0avgdata 12016maxresident)k
0inputs+0outputs (0major+474minor)pagefaults 0swaps
time ./decode-01
1.12user 0.00system 0:01.12elapsed 100%CPU (0avgtext+0avgdata 11348maxresident)k
0inputs+0outputs (0major+458minor)pagefaults 0swaps
time ./decode-02
1.14user 0.00system 0:01.14elapsed 99%CPU (0avgtext+0avgdata 11488maxresident)k
0inputs+0outputs (0major+462minor)pagefaults 0swaps
time ./decode-03
2.66user 0.00system 0:02.66elapsed 100%CPU (0avgtext+0avgdata 11388maxresident)k
0inputs+0outputs (0major+463minor)pagefaults 0swaps
```

### T4
```text
time ./encode-01
6.25user 0.00system 0:06.25elapsed 100%CPU (0avgtext+0avgdata 11888maxresident)k
0inputs+0outputs (0major+477minor)pagefaults 0swaps
time ./encode-02
14.96user 0.00system 0:14.96elapsed 99%CPU (0avgtext+0avgdata 11936maxresident)k
0inputs+0outputs (0major+475minor)pagefaults 0swaps
time ./encode-03
24.93user 0.00system 0:24.93elapsed 99%CPU (0avgtext+0avgdata 11892maxresident)k
0inputs+0outputs (0major+473minor)pagefaults 0swaps
time ./decode-01
1.11user 0.00system 0:01.11elapsed 99%CPU (0avgtext+0avgdata 11360maxresident)k
0inputs+0outputs (0major+463minor)pagefaults 0swaps
time ./decode-02
1.15user 0.00system 0:01.15elapsed 100%CPU (0avgtext+0avgdata 11464maxresident)k
0inputs+0outputs (0major+463minor)pagefaults 0swaps
time ./decode-03
2.68user 0.00system 0:02.68elapsed 99%CPU (0avgtext+0avgdata 11352maxresident)k
0inputs+0outputs (0major+463minor)pagefaults 0swaps
```

### T5
```text
time ./encode-01
6.09user 0.00system 0:06.10elapsed 99%CPU (0avgtext+0avgdata 11668maxresident)k
0inputs+0outputs (0major+473minor)pagefaults 0swaps
time ./encode-02
14.83user 0.00system 0:14.83elapsed 100%CPU (0avgtext+0avgdata 11960maxresident)k
0inputs+0outputs (0major+476minor)pagefaults 0swaps
time ./encode-03
25.75user 0.00system 0:25.75elapsed 99%CPU (0avgtext+0avgdata 11892maxresident)k
0inputs+0outputs (0major+477minor)pagefaults 0swaps
time ./decode-01
1.12user 0.00system 0:01.13elapsed 99%CPU (0avgtext+0avgdata 11344maxresident)k
0inputs+0outputs (0major+458minor)pagefaults 0swaps
time ./decode-02
1.10user 0.00system 0:01.10elapsed 100%CPU (0avgtext+0avgdata 11428maxresident)k
0inputs+0outputs (0major+462minor)pagefaults 0swaps
time ./decode-03
2.66user 0.00system 0:02.66elapsed 100%CPU (0avgtext+0avgdata 11376maxresident)k
0inputs+0outputs (0major+466minor)pagefaults 0swaps
```

Raw Data - cpp-template
-----------------------

### T1
```text
1.05user 0.00system 0:01.05elapsed 99%CPU (0avgtext+0avgdata 4068maxresident)k
0inputs+0outputs (0major+186minor)pagefaults 0swaps
time ./encode-02
9.76user 0.00system 0:09.77elapsed 99%CPU (0avgtext+0avgdata 4076maxresident)k
0inputs+0outputs (0major+188minor)pagefaults 0swaps
time ./encode-03
3.84user 0.00system 0:03.84elapsed 99%CPU (0avgtext+0avgdata 4052maxresident)k
0inputs+0outputs (0major+187minor)pagefaults 0swaps
time ./decode-01
0.37user 0.00system 0:00.37elapsed 100%CPU (0avgtext+0avgdata 2832maxresident)k
0inputs+0outputs (0major+112minor)pagefaults 0swaps
time ./decode-02
0.49user 0.00system 0:00.49elapsed 100%CPU (0avgtext+0avgdata 2848maxresident)k
0inputs+0outputs (0major+115minor)pagefaults 0swaps
time ./decode-03
1.34user 0.00system 0:01.34elapsed 99%CPU (0avgtext+0avgdata 2968maxresident)k
0inputs+0outputs (0major+114minor)pagefaults 0swaps
```

### T2
```text
time ./encode-01
1.05user 0.00system 0:01.05elapsed 99%CPU (0avgtext+0avgdata 4060maxresident)k
0inputs+0outputs (0major+186minor)pagefaults 0swaps
time ./encode-02
9.78user 0.00system 0:09.78elapsed 99%CPU (0avgtext+0avgdata 4096maxresident)k
0inputs+0outputs (0major+189minor)pagefaults 0swaps
time ./encode-03
3.79user 0.00system 0:03.79elapsed 100%CPU (0avgtext+0avgdata 4064maxresident)k
0inputs+0outputs (0major+186minor)pagefaults 0swaps
time ./decode-01
0.40user 0.00system 0:00.40elapsed 100%CPU (0avgtext+0avgdata 3012maxresident)k
0inputs+0outputs (0major+116minor)pagefaults 0swaps
time ./decode-02
0.47user 0.00system 0:00.47elapsed 100%CPU (0avgtext+0avgdata 3032maxresident)k
0inputs+0outputs (0major+116minor)pagefaults 0swaps
time ./decode-03
1.33user 0.00system 0:01.33elapsed 100%CPU (0avgtext+0avgdata 2928maxresident)k
0inputs+0outputs (0major+115minor)pagefaults 0swaps
```

### T3
```text
time ./encode-01
1.05user 0.00system 0:01.05elapsed 99%CPU (0avgtext+0avgdata 4212maxresident)k
0inputs+0outputs (0major+188minor)pagefaults 0swaps
time ./encode-02
9.83user 0.00system 0:09.83elapsed 99%CPU (0avgtext+0avgdata 4076maxresident)k
0inputs+0outputs (0major+190minor)pagefaults 0swaps
time ./encode-03
3.83user 0.00system 0:03.84elapsed 99%CPU (0avgtext+0avgdata 4152maxresident)k
0inputs+0outputs (0major+188minor)pagefaults 0swaps
time ./decode-01
0.39user 0.00system 0:00.40elapsed 99%CPU (0avgtext+0avgdata 2848maxresident)k
0inputs+0outputs (0major+112minor)pagefaults 0swaps
time ./decode-02
0.46user 0.00system 0:00.46elapsed 100%CPU (0avgtext+0avgdata 3016maxresident)k
0inputs+0outputs (0major+116minor)pagefaults 0swaps
time ./decode-03
1.32user 0.00system 0:01.32elapsed 99%CPU (0avgtext+0avgdata 2836maxresident)k
0inputs+0outputs (0major+113minor)pagefaults 0swaps
```

### T4
```text
time ./encode-01
1.08user 0.00system 0:01.08elapsed 100%CPU (0avgtext+0avgdata 4212maxresident)k
0inputs+0outputs (0major+188minor)pagefaults 0swaps
time ./encode-02
10.14user 0.00system 0:10.14elapsed 99%CPU (0avgtext+0avgdata 4176maxresident)k
0inputs+0outputs (0major+189minor)pagefaults 0swaps
time ./encode-03
3.85user 0.00system 0:03.85elapsed 100%CPU (0avgtext+0avgdata 4060maxresident)k
0inputs+0outputs (0major+186minor)pagefaults 0swaps
time ./decode-01
0.38user 0.00system 0:00.39elapsed 99%CPU (0avgtext+0avgdata 3020maxresident)k
0inputs+0outputs (0major+114minor)pagefaults 0swaps
time ./decode-02
0.46user 0.00system 0:00.46elapsed 100%CPU (0avgtext+0avgdata 2928maxresident)k
0inputs+0outputs (0major+116minor)pagefaults 0swaps
time ./decode-03
1.32user 0.00system 0:01.33elapsed 99%CPU (0avgtext+0avgdata 3004maxresident)k
0inputs+0outputs (0major+116minor)pagefaults 0swaps
```

### T5
```text
time ./encode-01
1.05user 0.00system 0:01.06elapsed 99%CPU (0avgtext+0avgdata 4096maxresident)k
0inputs+0outputs (0major+187minor)pagefaults 0swaps
time ./encode-02
9.79user 0.00system 0:09.79elapsed 99%CPU (0avgtext+0avgdata 4088maxresident)k
0inputs+0outputs (0major+188minor)pagefaults 0swaps
time ./encode-03
3.83user 0.00system 0:03.83elapsed 99%CPU (0avgtext+0avgdata 4104maxresident)k
0inputs+0outputs (0major+186minor)pagefaults 0swaps
time ./decode-01
0.39user 0.00system 0:00.39elapsed 100%CPU (0avgtext+0avgdata 2876maxresident)k
0inputs+0outputs (0major+113minor)pagefaults 0swaps
time ./decode-02
0.46user 0.00system 0:00.46elapsed 100%CPU (0avgtext+0avgdata 2932maxresident)k
0inputs+0outputs (0major+115minor)pagefaults 0swaps
time ./decode-03
1.32user 0.00system 0:01.32elapsed 99%CPU (0avgtext+0avgdata 3000maxresident)k
0inputs+0outputs (0major+115minor)pagefaults 0swaps
```

Raw Data - python-ndn
---------------------

### T1
```text
Total time: 118.1406 seconds
Total time: 144.1903 seconds
Total time: 222.5538 seconds
Total time: 37.1360 seconds
Total time: 40.0556 seconds
Total time: 54.7339 seconds
```

### T2
```text
Total time: 117.6044 seconds
Total time: 141.2825 seconds
Total time: 220.2581 seconds
Total time: 37.4249 seconds
Total time: 39.7109 seconds
Total time: 54.6434 seconds
```

### T3
```text
Total time: 116.9281 seconds
Total time: 140.2277 seconds
Total time: 221.0614 seconds
Total time: 38.0363 seconds
Total time: 40.0673 seconds
Total time: 55.5637 seconds
```

### T4
```text
Total time: 117.2334 seconds
Total time: 142.4005 seconds
Total time: 220.7994 seconds
Total time: 38.1088 seconds
Total time: 39.4046 seconds
Total time: 54.5650 seconds
```

### T5
```text
Total time: 120.2042 seconds
Total time: 139.4211 seconds
Total time: 223.6294 seconds
Total time: 37.0010 seconds
Total time: 39.5002 seconds
Total time: 53.8632 seconds
```

Raw Data - python-ndn (pypy3)
---------------------

### T1
```text
./pypy3 encode-01.py && ./pypy3 encode-02.py && ./pypy3 encode-03.py && ./pypy3 decode-01.py && ./pypy3 decode-02.py && ./pypy3 decode-03.py 
Total time: 23.2973 seconds
Total time: 44.3198 seconds
Total time: 39.9096 seconds
Total time: 3.8774 seconds
Total time: 4.0191 seconds
Total time: 4.6539 seconds
```

### T2
```text
./pypy3 encode-01.py && ./pypy3 encode-02.py && ./pypy3 encode-03.py && ./pypy3 decode-01.py && ./pypy3 decode-02.py && ./pypy3 decode-03.py 
Total time: 22.8894 seconds
Total time: 44.0513 seconds
Total time: 40.6131 seconds
Total time: 3.9117 seconds
Total time: 4.1674 seconds
Total time: 4.6047 seconds
```

### T3
```text
./pypy3 encode-01.py && ./pypy3 encode-02.py && ./pypy3 encode-03.py && ./pypy3 decode-01.py && ./pypy3 decode-02.py && ./pypy3 decode-03.py 
Total time: 22.7557 seconds
Total time: 46.2450 seconds
Total time: 42.4678 seconds
Total time: 3.9472 seconds
Total time: 3.9590 seconds
Total time: 4.5828 seconds
```

### T4
```text
./pypy3 encode-01.py && ./pypy3 encode-02.py && ./pypy3 encode-03.py && ./pypy3 decode-01.py && ./pypy3 decode-02.py && ./pypy3 decode-03.py 
Total time: 23.5257 seconds
Total time: 44.8939 seconds
Total time: 40.7754 seconds
Total time: 3.9525 seconds
Total time: 3.9701 seconds
Total time: 4.5507 seconds
```

### T5
```text
./pypy3 encode-01.py && ./pypy3 encode-02.py && ./pypy3 encode-03.py && ./pypy3 decode-01.py && ./pypy3 decode-02.py && ./pypy3 decode-03.py 
Total time: 22.4416 seconds
Total time: 45.0505 seconds
Total time: 42.3590 seconds
Total time: 4.0875 seconds
Total time: 3.9862 seconds
Total time: 4.6395 seconds
```

Raw Data - cloc
---------------

### YaNFD
```text
-------------------------------------------------------------------------------
File                             blank        comment           code
-------------------------------------------------------------------------------
./name.go                          125            120            688
./tlv/block.go                      46             58            229
./tlv/helpers.go                    19             16            127
./tlv/errors.go                      3              7              8
-------------------------------------------------------------------------------
SUM:                               193            201           1052
-------------------------------------------------------------------------------
```

### Reflect
```text
-------------------------------------------------------------------------------
File                             blank        comment           code
-------------------------------------------------------------------------------
./struct.go                         12              7            119
./encode.go                          4              5             63
./name.go                           10              1             63
./natural.go                         6              1             58
./tlvvar.go                          5              1             47
./sequence.go                        5              1             37
./binary.go                          5              1             24
./interfaces.go                      3              1             11
-------------------------------------------------------------------------------
SUM:                                50             18            422
-------------------------------------------------------------------------------
```

### CodeGen
```text
-------------------------------------------------------------------------------
File                             blank        comment           code
-------------------------------------------------------------------------------
./fields.go                         27              1            518
./codegen.go                        12              9            148
-------------------------------------------------------------------------------
SUM:                                39             10            666
-------------------------------------------------------------------------------
```

### ndn-cxx
```text
---------------------------------------------------------------------------------------
File                                     blank        comment           code
---------------------------------------------------------------------------------------
./block.cpp                                 92             34            411
./tlv.hpp                                   59            143            337
./encoder.cpp                               36             20            218
./block.hpp                                 69            257            182
./encoder.hpp                               52            141            150
./block-helpers.hpp                         47            131            149
./block-helpers.cpp                         43             25            131
./estimator.cpp                             12             20             75
./buffer.hpp                                23             58             62
./estimator.hpp                             19             71             60
./tlv.cpp                                    7             20             56
./encoding-buffer.hpp                        8             26             34
./buffer-stream.hpp                         16             41             33
./buffer-stream.cpp                          9             20             31
./encoding-buffer-fwd.hpp                   10             20             31
./buffer.cpp                                 4             20             10
---------------------------------------------------------------------------------------
SUM:                                       506           1047           1970
---------------------------------------------------------------------------------------
```

However, in `tlv.hpp` there are 80 LOC defining constants for the spec, excluded.

### C++ template
```text
-------------------------------------------------------------------------------
File                             blank        comment           code
-------------------------------------------------------------------------------
./tlv-encoder.hpp                   83             79            566
./slice.hpp                         21              3             87
-------------------------------------------------------------------------------
SUM:                               104             82            653
-------------------------------------------------------------------------------
```

### python-ndn
```text
---------------------------------------------------------------------------------
File                               blank        comment           code
---------------------------------------------------------------------------------
./tlv_model.py                       135            239            530
./name/Component.py                   63            137            154
./name/Name.py                        42             81             96
./tlv_var.py                          23             55             86
./__init__.py                          3              0             15
./signer.py                            9             34             14
./tlv_type.py                         12             27             13
./name/__init__.py                     0              0              1
---------------------------------------------------------------------------------
SUM:                                 287            573            909
---------------------------------------------------------------------------------
```

Memory - Go
-----------

```text
========== ENCODING ==========
block:
#1: 1 x 100B
Size of Data (before): 575
Size of Data (after): 1719
Size of Wire:  213
reflection:
Size of Data:  341
Size of Wire:  213
codegen:
Size of Data:  341
Size of Wire:  213
=== Total Bytes: 189 ===

#2: 1 x 4000B
block:
Size of Data (before): 4475
Size of Data (after): 13425
Size of Wire:  4117
reflection:
Size of Data:  4241
Size of Wire:  4117
codegen:
Size of Data:  4241
Size of Wire:  4117
=== Total Bytes: 4093 ===

#3: 1 x 10B, very long name
block:
Size of Data (before): 2255
Size of Data (after): 4569
Size of Wire:  213
reflection:
Size of Data:  1061
Size of Wire:  213
codegen:
Size of Data:  1061
Size of Wire:  213
=== Total Bytes: 189 ===
```

Memory - ndn-cxx
----------------

### encode-01 (before encoding)

```text
sizeof(data) = 1016
Referred to Size = 639

>>TIME:2
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-01                           0x000000010ae7b0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010aec46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010aeb8966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010af494bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010af49fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010af498c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010af52050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010af51c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-01                           0x000000010ae78e47 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-01                           0x000000010ae7930d main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:3
>>SIZE:8
>>DEPT:12
>>ADDR:
0   encode-01                           0x000000010ae7b0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010aec4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010aeb8966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010af494bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010af49fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010af498c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010af52050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010af51c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-01                           0x000000010ae78e47 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-01                           0x000000010ae7930d main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:5
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-01                           0x000000010ae7b0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010aec46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010aeb8966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010af494bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010af49fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010af498c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010af52050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010af51c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-01                           0x000000010ae78e47 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-01                           0x000000010ae7930d main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:6
>>SIZE:6
>>DEPT:12
>>ADDR:
0   encode-01                           0x000000010ae7b0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010aec4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010aeb8966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010af494bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010af49fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010af498c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010af52050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010af51c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-01                           0x000000010ae78e47 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-01                           0x000000010ae7930d main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:8
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-01                           0x000000010ae7b0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010aec46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010aeb8966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010af494bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010af49fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010af498c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010af52050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010af51c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-01                           0x000000010ae78e47 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-01                           0x000000010ae7930d main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:9
>>SIZE:12
>>DEPT:12
>>ADDR:
0   encode-01                           0x000000010ae7b0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010aec4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010aeb8966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010af494bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010af49fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010af498c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010af52050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010af51c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-01                           0x000000010ae78e47 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-01                           0x000000010ae7930d main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:11
>>SIZE:264
>>DEPT:8
>>ADDR:
0   encode-01                           0x000000010ae7b0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010aec17e5 _ZNSt3__16vectorIN3ndn5BlockENS_9allocatorIS2_EEEC2ERKS5_ + 101
2   libndn-cxx.0.7.1.dylib              0x000000010aeb963f _ZN3ndn5BlockC1ERKS0_ + 95
3   libndn-cxx.0.7.1.dylib              0x000000010ae97fb7 _ZN3ndn4DataC2ERKNS_4NameE + 71
4   encode-01                           0x000000010ae78e5a _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 106
5   encode-01                           0x000000010ae7930d main + 189
6   libdyld.dylib                       0x00007fff20544f3d start + 1
7   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:12
>>SIZE:48
>>DEPT:9
>>ADDR:
0   encode-01                           0x000000010ae7b0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010aec46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010aeb8966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010aeb835e _ZN3ndn8encoding15makeBinaryBlockEjPKcm + 14
4   libndn-cxx.0.7.1.dylib              0x000000010af49652 _ZN3ndn4name9ComponentC1EPKc + 50
5   encode-01                           0x000000010ae78ead _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 189
6   encode-01                           0x000000010ae7930d main + 189
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:13
>>SIZE:7
>>DEPT:9
>>ADDR:
0   encode-01                           0x000000010ae7b0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010aec4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010aeb8966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010aeb835e _ZN3ndn8encoding15makeBinaryBlockEjPKcm + 14
4   libndn-cxx.0.7.1.dylib              0x000000010af49652 _ZN3ndn4name9ComponentC1EPKc + 50
5   encode-01                           0x000000010ae78ead _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 189
6   encode-01                           0x000000010ae7930d main + 189
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:14
>>SIZE:48
>>DEPT:8
>>ADDR:
0   encode-01                           0x000000010ae7b0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010aec46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010aeb8966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010ae9b35c _ZN3ndn4Data10setContentEPKhm + 60
4   encode-01                           0x000000010ae79080 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 656
5   encode-01                           0x000000010ae7930d main + 189
6   libdyld.dylib                       0x00007fff20544f3d start + 1
7   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:15
>>SIZE:102
>>DEPT:8
>>ADDR:
0   encode-01                           0x000000010ae7b0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010aec4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010aeb8966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010ae9b35c _ZN3ndn4Data10setContentEPKhm + 60
4   encode-01                           0x000000010ae79080 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 656
5   encode-01                           0x000000010ae7930d main + 189
6   libdyld.dylib                       0x00007fff20544f3d start + 1
7   ???                                 0x0000000000000001 0x0 + 1
<<
```

### encode-02 (before encoding)

```text
sizeof(data) = 1016
Referred to Size = 4541

>>TIME:2
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-02                           0x000000010ca800d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010cac96f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010cabd966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010cb4e4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010cb4efe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010cb4e8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010cb57050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010cb56c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-02                           0x000000010ca7de2f _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 95
9   encode-02                           0x000000010ca7e2fd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:3
>>SIZE:8
>>DEPT:12
>>ADDR:
0   encode-02                           0x000000010ca800d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010cac9741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010cabd966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010cb4e4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010cb4efe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010cb4e8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010cb57050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010cb56c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-02                           0x000000010ca7de2f _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 95
9   encode-02                           0x000000010ca7e2fd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:5
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-02                           0x000000010ca800d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010cac96f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010cabd966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010cb4e4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010cb4efe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010cb4e8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010cb57050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010cb56c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-02                           0x000000010ca7de2f _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 95
9   encode-02                           0x000000010ca7e2fd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:6
>>SIZE:6
>>DEPT:12
>>ADDR:
0   encode-02                           0x000000010ca800d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010cac9741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010cabd966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010cb4e4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010cb4efe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010cb4e8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010cb57050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010cb56c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-02                           0x000000010ca7de2f _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 95
9   encode-02                           0x000000010ca7e2fd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:8
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-02                           0x000000010ca800d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010cac96f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010cabd966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010cb4e4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010cb4efe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010cb4e8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010cb57050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010cb56c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-02                           0x000000010ca7de2f _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 95
9   encode-02                           0x000000010ca7e2fd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:9
>>SIZE:12
>>DEPT:12
>>ADDR:
0   encode-02                           0x000000010ca800d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010cac9741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010cabd966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010cb4e4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010cb4efe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010cb4e8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010cb57050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010cb56c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-02                           0x000000010ca7de2f _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 95
9   encode-02                           0x000000010ca7e2fd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:11
>>SIZE:264
>>DEPT:8
>>ADDR:
0   encode-02                           0x000000010ca800d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010cac67e5 _ZNSt3__16vectorIN3ndn5BlockENS_9allocatorIS2_EEEC2ERKS5_ + 101
2   libndn-cxx.0.7.1.dylib              0x000000010cabe63f _ZN3ndn5BlockC1ERKS0_ + 95
3   libndn-cxx.0.7.1.dylib              0x000000010ca9cfb7 _ZN3ndn4DataC2ERKNS_4NameE + 71
4   encode-02                           0x000000010ca7de42 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 114
5   encode-02                           0x000000010ca7e2fd main + 189
6   libdyld.dylib                       0x00007fff20544f3d start + 1
7   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:12
>>SIZE:48
>>DEPT:9
>>ADDR:
0   encode-02                           0x000000010ca800d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010cac96f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010cabd966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010cabd35e _ZN3ndn8encoding15makeBinaryBlockEjPKcm + 14
4   libndn-cxx.0.7.1.dylib              0x000000010cb4e652 _ZN3ndn4name9ComponentC1EPKc + 50
5   encode-02                           0x000000010ca7de95 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 197
6   encode-02                           0x000000010ca7e2fd main + 189
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:13
>>SIZE:7
>>DEPT:9
>>ADDR:
0   encode-02                           0x000000010ca800d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010cac9741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010cabd966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010cabd35e _ZN3ndn8encoding15makeBinaryBlockEjPKcm + 14
4   libndn-cxx.0.7.1.dylib              0x000000010cb4e652 _ZN3ndn4name9ComponentC1EPKc + 50
5   encode-02                           0x000000010ca7de95 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 197
6   encode-02                           0x000000010ca7e2fd main + 189
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:14
>>SIZE:48
>>DEPT:8
>>ADDR:
0   encode-02                           0x000000010ca800d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010cac96f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010cabd966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010caa035c _ZN3ndn4Data10setContentEPKhm + 60
4   encode-02                           0x000000010ca7e070 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 672
5   encode-02                           0x000000010ca7e2fd main + 189
6   libdyld.dylib                       0x00007fff20544f3d start + 1
7   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:15
>>SIZE:4004
>>DEPT:8
>>ADDR:
0   encode-02                           0x000000010ca800d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010cac9741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010cabd966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010caa035c _ZN3ndn4Data10setContentEPKhm + 60
4   encode-02                           0x000000010ca7e070 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 672
5   encode-02                           0x000000010ca7e2fd main + 189
6   libdyld.dylib                       0x00007fff20544f3d start + 1
7   ???                                 0x0000000000000001 0x0 + 1
<<
```

### encode-03 (before encoding)

```text
sizeof(data) = 1016
Referred to Size = 4719

>>TIME:2
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:3
>>SIZE:8
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:5
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:6
>>SIZE:6
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:8
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:9
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:11
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:12
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:13
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:14
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:16
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:17
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:18
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:19
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:20
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:21
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:22
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:23
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:25
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:26
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:27
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:28
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:29
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:30
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:31
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:32
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:33
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:34
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:35
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:36
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:37
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:38
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:39
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:40
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:42
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:43
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:44
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:45
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:46
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:47
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:48
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:49
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:50
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:51
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:52
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:53
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:54
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:55
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:56
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:57
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:58
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:59
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:60
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:61
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:62
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:63
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:64
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:65
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:66
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:67
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:68
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:69
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:70
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:71
>>SIZE:3
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:72
>>SIZE:48
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:73
>>SIZE:12
>>DEPT:12
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010fa294bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010fa29fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010fa298c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010fa32050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010fa31c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   encode-03                           0x000000010f95de07 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 87
9   encode-03                           0x000000010f95e2cd main + 189
10  libdyld.dylib                       0x00007fff20544f3d start + 1
11  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:75
>>SIZE:2904
>>DEPT:8
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a17e5 _ZNSt3__16vectorIN3ndn5BlockENS_9allocatorIS2_EEEC2ERKS5_ + 101
2   libndn-cxx.0.7.1.dylib              0x000000010f99963f _ZN3ndn5BlockC1ERKS0_ + 95
3   libndn-cxx.0.7.1.dylib              0x000000010f977fb7 _ZN3ndn4DataC2ERKNS_4NameE + 71
4   encode-03                           0x000000010f95de1a _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 106
5   encode-03                           0x000000010f95e2cd main + 189
6   libdyld.dylib                       0x00007fff20544f3d start + 1
7   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:76
>>SIZE:48
>>DEPT:9
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010f99835e _ZN3ndn8encoding15makeBinaryBlockEjPKcm + 14
4   libndn-cxx.0.7.1.dylib              0x000000010fa29652 _ZN3ndn4name9ComponentC1EPKc + 50
5   encode-03                           0x000000010f95de6d _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 189
6   encode-03                           0x000000010f95e2cd main + 189
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:77
>>SIZE:7
>>DEPT:9
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010f99835e _ZN3ndn8encoding15makeBinaryBlockEjPKcm + 14
4   libndn-cxx.0.7.1.dylib              0x000000010fa29652 _ZN3ndn4name9ComponentC1EPKc + 50
5   encode-03                           0x000000010f95de6d _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 189
6   encode-03                           0x000000010f95e2cd main + 189
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:78
>>SIZE:48
>>DEPT:8
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a46f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010f97b35c _ZN3ndn4Data10setContentEPKhm + 60
4   encode-03                           0x000000010f95e040 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 656
5   encode-03                           0x000000010f95e2cd main + 189
6   libdyld.dylib                       0x00007fff20544f3d start + 1
7   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:79
>>SIZE:12
>>DEPT:8
>>ADDR:
0   encode-03                           0x000000010f9600a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f9a4741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f998966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010f97b35c _ZN3ndn4Data10setContentEPKhm + 60
4   encode-03                           0x000000010f95e040 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 656
5   encode-03                           0x000000010f95e2cd main + 189
6   libdyld.dylib                       0x00007fff20544f3d start + 1
7   ???                                 0x0000000000000001 0x0 + 1
<<
```

### encode-01 (after encoding)

```text
sizeof(data) = 1016
Referred to Size = 11076

>>TIME:11
>>SIZE:264
>>DEPT:8
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d37e5 _ZNSt3__16vectorIN3ndn5BlockENS_9allocatorIS2_EEEC2ERKS5_ + 101
2   libndn-cxx.0.7.1.dylib              0x000000010f4cb63f _ZN3ndn5BlockC1ERKS0_ + 95
3   libndn-cxx.0.7.1.dylib              0x000000010f4a9fb7 _ZN3ndn4DataC2ERKNS_4NameE + 71
4   encode-01                           0x000000010f48ce5a _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 106
5   encode-01                           0x000000010f48d30d main + 189
6   libdyld.dylib                       0x00007fff20544f3d start + 1
7   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:16
>>SIZE:352
>>DEPT:11
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010fc4c4e4 _ZN5boost3log11v2_mt_posix13attribute_setC1Ev + 20
2   libndn-cxx.0.7.1.dylib              0x000000010f6598cc _ZN5boost3log11v2_mt_posix7sources12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEC2Ev + 44
3   libndn-cxx.0.7.1.dylib              0x000000010f6597c2 _ZN5boost3log11v2_mt_posix7sources21basic_severity_loggerINS2_12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEES8_EC2Ev + 18
4   libndn-cxx.0.7.1.dylib              0x000000010f658b2f _ZN3ndn4util6LoggerC2EPKc + 31
5   libndn-cxx.0.7.1.dylib              0x000000010f5b0779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
6   libndn-cxx.0.7.1.dylib              0x000000010f5af86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
7   encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
8   encode-01                           0x000000010f48d30d main + 189
9   libdyld.dylib                       0x00007fff20544f3d start + 1
10  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:17
>>SIZE:32
>>DEPT:10
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010fc4c89a _ZN5boost3log11v2_mt_posix13attribute_set14implementation6insertENS1_14attribute_nameERKNS1_9attributeE + 266
2   libndn-cxx.0.7.1.dylib              0x000000010f659816 _ZN5boost3log11v2_mt_posix7sources21basic_severity_loggerINS2_12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEES8_EC2Ev + 102
3   libndn-cxx.0.7.1.dylib              0x000000010f658b2f _ZN3ndn4util6LoggerC2EPKc + 31
4   libndn-cxx.0.7.1.dylib              0x000000010f5b0779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
5   libndn-cxx.0.7.1.dylib              0x000000010f5af86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
6   encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
7   encode-01                           0x000000010f48d30d main + 189
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:18
>>SIZE:32
>>DEPT:9
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010fc4c89a _ZN5boost3log11v2_mt_posix13attribute_set14implementation6insertENS1_14attribute_nameERKNS1_9attributeE + 266
2   libndn-cxx.0.7.1.dylib              0x000000010f658c40 _ZN3ndn4util6LoggerC2EPKc + 304
3   libndn-cxx.0.7.1.dylib              0x000000010f5b0779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
4   libndn-cxx.0.7.1.dylib              0x000000010f5af86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
5   encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
6   encode-01                           0x000000010f48d30d main + 189
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:19
>>SIZE:48
>>DEPT:10
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f65c1c8 _ZNSt3__112__hash_tableINS_17__hash_value_typeINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEEPN3ndn4util6LoggerEEENS_22__unordered_map_hasherIS7_SC_NS_4hashIS7_EELb1EEENS_21__unordered_map_equalIS7_SC_NS_8equal_toIS7_EELb1EEENS5_ISC_EEE16__construct_nodeIJRKS7_SB_EEENS_10unique_ptrINS_11__hash_nodeISC_PvEENS_22__hash_node_destructorINS5_IST_EEEEEEDpOT_ + 40
2   libndn-cxx.0.7.1.dylib              0x000000010f65adcc _ZN3ndn4util7Logging13addLoggerImplERNS0_6LoggerE + 60
3   libndn-cxx.0.7.1.dylib              0x000000010f658c79 _ZN3ndn4util6LoggerC2EPKc + 361
4   libndn-cxx.0.7.1.dylib              0x000000010f5b0779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
5   libndn-cxx.0.7.1.dylib              0x000000010f5af86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
6   encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
7   encode-01                           0x000000010f48d30d main + 189
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:21
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d66f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f4ca966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010f55b4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010f55bfe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010f55b8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010f564050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010f563c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010f563e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010f5cbb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010f5b0579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010f5af86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-01                           0x000000010f48d30d main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:22
>>SIZE:11
>>DEPT:16
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d6741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f4ca966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010f55b4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010f55bfe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010f55b8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010f564050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010f563c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010f563e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010f5cbb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010f5b0579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010f5af86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-01                           0x000000010f48d30d main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:24
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d66f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f4ca966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010f55b4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010f55bfe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010f55b8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010f564050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010f563c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010f563e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010f5cbb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010f5b0579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010f5af86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-01                           0x000000010f48d30d main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:25
>>SIZE:10
>>DEPT:16
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d6741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f4ca966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010f55b4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010f55bfe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010f55b8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010f564050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010f563c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010f563e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010f5cbb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010f5b0579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010f5af86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-01                           0x000000010f48d30d main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:27
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d66f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f4ca966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010f55b4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010f55bfe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010f55b8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010f564050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010f563c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010f563e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010f5cbb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010f5b0579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010f5af86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-01                           0x000000010f48d30d main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:28
>>SIZE:15
>>DEPT:16
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d6741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f4ca966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010f55b4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010f55bfe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010f55b8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010f564050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010f563c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010f563e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010f5cbb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010f5b0579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010f5af86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-01                           0x000000010f48d30d main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:29
>>SIZE:352
>>DEPT:11
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d1fba _ZN3ndn5Block9push_backERKS0_ + 346
2   libndn-cxx.0.7.1.dylib              0x000000010f563c7d _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 509
3   libndn-cxx.0.7.1.dylib              0x000000010f563e66 _ZN3ndn4NameC1EPKc + 134
4   libndn-cxx.0.7.1.dylib              0x000000010f5cbb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
5   libndn-cxx.0.7.1.dylib              0x000000010f5b0579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
6   libndn-cxx.0.7.1.dylib              0x000000010f5af86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
7   encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
8   encode-01                           0x000000010f48d30d main + 189
9   libdyld.dylib                       0x00007fff20544f3d start + 1
10  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:31
>>SIZE:48
>>DEPT:7
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d66f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010f5af941 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 305
3   encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
4   encode-01                           0x000000010f48d30d main + 189
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:32
>>SIZE:8800
>>DEPT:7
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d6741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010f5af941 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 305
3   encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
4   encode-01                           0x000000010f48d30d main + 189
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:46
>>SIZE:704
>>DEPT:9
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d0b51 _ZNK3ndn5Block5parseEv + 481
2   libndn-cxx.0.7.1.dylib              0x000000010f4aa3ec _ZN3ndn4Data10wireDecodeERKNS_5BlockE + 76
3   libndn-cxx.0.7.1.dylib              0x000000010f4ac3e0 _ZNK3ndn4Data10wireEncodeERNS_8encoding12EncodingImplILNS1_3TagE1EEERKNS_5BlockE + 96
4   libndn-cxx.0.7.1.dylib              0x000000010f5afa41 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 561
5   encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
6   encode-01                           0x000000010f48d30d main + 189
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:48
>>SIZE:176
>>DEPT:10
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d0b51 _ZNK3ndn5Block5parseEv + 481
2   libndn-cxx.0.7.1.dylib              0x000000010f518069 _ZN3ndn8MetaInfo10wireDecodeERKNS_5BlockE + 41
3   libndn-cxx.0.7.1.dylib              0x000000010f4aa892 _ZN3ndn4Data10wireDecodeERKNS_5BlockE + 1266
4   libndn-cxx.0.7.1.dylib              0x000000010f4ac3e0 _ZNK3ndn4Data10wireEncodeERNS_8encoding12EncodingImplILNS1_3TagE1EEERKNS_5BlockE + 96
5   libndn-cxx.0.7.1.dylib              0x000000010f5afa41 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 561
6   encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
7   encode-01                           0x000000010f48d30d main + 189
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:49
>>SIZE:88
>>DEPT:10
>>ADDR:
0   encode-01                           0x000000010f48f0e3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010f4d0b51 _ZNK3ndn5Block5parseEv + 481
2   libndn-cxx.0.7.1.dylib              0x000000010f62e6b4 _ZN3ndn13SignatureInfo10wireDecodeERKNS_5BlockENS0_4TypeE + 100
3   libndn-cxx.0.7.1.dylib              0x000000010f4aa8bd _ZN3ndn4Data10wireDecodeERKNS_5BlockE + 1309
4   libndn-cxx.0.7.1.dylib              0x000000010f4ac3e0 _ZNK3ndn4Data10wireEncodeERNS_8encoding12EncodingImplILNS1_3TagE1EEERKNS_5BlockE + 96
5   libndn-cxx.0.7.1.dylib              0x000000010f5afa41 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 561
6   encode-01                           0x000000010f48d0a6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
7   encode-01                           0x000000010f48d30d main + 189
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
```

### encode-02 (after encoding)

```text
sizeof(data) = 1016
Referred to Size = 11076

>>TIME:11
>>SIZE:264
>>DEPT:8
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d5607e5 _ZNSt3__16vectorIN3ndn5BlockENS_9allocatorIS2_EEEC2ERKS5_ + 101
2   libndn-cxx.0.7.1.dylib              0x000000010d55863f _ZN3ndn5BlockC1ERKS0_ + 95
3   libndn-cxx.0.7.1.dylib              0x000000010d536fb7 _ZN3ndn4DataC2ERKNS_4NameE + 71
4   encode-02                           0x000000010d518e42 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 114
5   encode-02                           0x000000010d5192fd main + 189
6   libdyld.dylib                       0x00007fff20544f3d start + 1
7   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:16
>>SIZE:352
>>DEPT:11
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010dce14e4 _ZN5boost3log11v2_mt_posix13attribute_setC1Ev + 20
2   libndn-cxx.0.7.1.dylib              0x000000010d6e68cc _ZN5boost3log11v2_mt_posix7sources12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEC2Ev + 44
3   libndn-cxx.0.7.1.dylib              0x000000010d6e67c2 _ZN5boost3log11v2_mt_posix7sources21basic_severity_loggerINS2_12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEES8_EC2Ev + 18
4   libndn-cxx.0.7.1.dylib              0x000000010d6e5b2f _ZN3ndn4util6LoggerC2EPKc + 31
5   libndn-cxx.0.7.1.dylib              0x000000010d63d779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
6   libndn-cxx.0.7.1.dylib              0x000000010d63c86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
7   encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
8   encode-02                           0x000000010d5192fd main + 189
9   libdyld.dylib                       0x00007fff20544f3d start + 1
10  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:17
>>SIZE:32
>>DEPT:10
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010dce189a _ZN5boost3log11v2_mt_posix13attribute_set14implementation6insertENS1_14attribute_nameERKNS1_9attributeE + 266
2   libndn-cxx.0.7.1.dylib              0x000000010d6e6816 _ZN5boost3log11v2_mt_posix7sources21basic_severity_loggerINS2_12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEES8_EC2Ev + 102
3   libndn-cxx.0.7.1.dylib              0x000000010d6e5b2f _ZN3ndn4util6LoggerC2EPKc + 31
4   libndn-cxx.0.7.1.dylib              0x000000010d63d779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
5   libndn-cxx.0.7.1.dylib              0x000000010d63c86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
6   encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
7   encode-02                           0x000000010d5192fd main + 189
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:18
>>SIZE:32
>>DEPT:9
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010dce189a _ZN5boost3log11v2_mt_posix13attribute_set14implementation6insertENS1_14attribute_nameERKNS1_9attributeE + 266
2   libndn-cxx.0.7.1.dylib              0x000000010d6e5c40 _ZN3ndn4util6LoggerC2EPKc + 304
3   libndn-cxx.0.7.1.dylib              0x000000010d63d779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
4   libndn-cxx.0.7.1.dylib              0x000000010d63c86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
5   encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
6   encode-02                           0x000000010d5192fd main + 189
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:19
>>SIZE:48
>>DEPT:10
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d6e91c8 _ZNSt3__112__hash_tableINS_17__hash_value_typeINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEEPN3ndn4util6LoggerEEENS_22__unordered_map_hasherIS7_SC_NS_4hashIS7_EELb1EEENS_21__unordered_map_equalIS7_SC_NS_8equal_toIS7_EELb1EEENS5_ISC_EEE16__construct_nodeIJRKS7_SB_EEENS_10unique_ptrINS_11__hash_nodeISC_PvEENS_22__hash_node_destructorINS5_IST_EEEEEEDpOT_ + 40
2   libndn-cxx.0.7.1.dylib              0x000000010d6e7dcc _ZN3ndn4util7Logging13addLoggerImplERNS0_6LoggerE + 60
3   libndn-cxx.0.7.1.dylib              0x000000010d6e5c79 _ZN3ndn4util6LoggerC2EPKc + 361
4   libndn-cxx.0.7.1.dylib              0x000000010d63d779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
5   libndn-cxx.0.7.1.dylib              0x000000010d63c86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
6   encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
7   encode-02                           0x000000010d5192fd main + 189
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:21
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d5636f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010d557966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010d5e84bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010d5e8fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010d5e88c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010d5f1050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010d5f0c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010d5f0e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010d658b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010d63d579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010d63c86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
13  encode-02                           0x000000010d5192fd main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:22
>>SIZE:11
>>DEPT:16
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d563741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010d557966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010d5e84bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010d5e8fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010d5e88c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010d5f1050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010d5f0c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010d5f0e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010d658b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010d63d579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010d63c86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
13  encode-02                           0x000000010d5192fd main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:24
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d5636f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010d557966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010d5e84bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010d5e8fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010d5e88c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010d5f1050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010d5f0c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010d5f0e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010d658b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010d63d579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010d63c86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
13  encode-02                           0x000000010d5192fd main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:25
>>SIZE:10
>>DEPT:16
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d563741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010d557966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010d5e84bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010d5e8fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010d5e88c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010d5f1050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010d5f0c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010d5f0e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010d658b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010d63d579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010d63c86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
13  encode-02                           0x000000010d5192fd main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:27
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d5636f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010d557966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010d5e84bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010d5e8fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010d5e88c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010d5f1050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010d5f0c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010d5f0e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010d658b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010d63d579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010d63c86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
13  encode-02                           0x000000010d5192fd main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:28
>>SIZE:15
>>DEPT:16
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d563741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010d557966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010d5e84bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010d5e8fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010d5e88c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010d5f1050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010d5f0c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010d5f0e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010d658b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010d63d579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010d63c86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
13  encode-02                           0x000000010d5192fd main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:29
>>SIZE:352
>>DEPT:11
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d55efba _ZN3ndn5Block9push_backERKS0_ + 346
2   libndn-cxx.0.7.1.dylib              0x000000010d5f0c7d _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 509
3   libndn-cxx.0.7.1.dylib              0x000000010d5f0e66 _ZN3ndn4NameC1EPKc + 134
4   libndn-cxx.0.7.1.dylib              0x000000010d658b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
5   libndn-cxx.0.7.1.dylib              0x000000010d63d579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
6   libndn-cxx.0.7.1.dylib              0x000000010d63c86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
7   encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
8   encode-02                           0x000000010d5192fd main + 189
9   libdyld.dylib                       0x00007fff20544f3d start + 1
10  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:31
>>SIZE:48
>>DEPT:7
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d5636f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010d63c941 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 305
3   encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
4   encode-02                           0x000000010d5192fd main + 189
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:32
>>SIZE:8800
>>DEPT:7
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d563741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010d63c941 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 305
3   encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
4   encode-02                           0x000000010d5192fd main + 189
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:46
>>SIZE:704
>>DEPT:9
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d55db51 _ZNK3ndn5Block5parseEv + 481
2   libndn-cxx.0.7.1.dylib              0x000000010d5373ec _ZN3ndn4Data10wireDecodeERKNS_5BlockE + 76
3   libndn-cxx.0.7.1.dylib              0x000000010d5393e0 _ZNK3ndn4Data10wireEncodeERNS_8encoding12EncodingImplILNS1_3TagE1EEERKNS_5BlockE + 96
4   libndn-cxx.0.7.1.dylib              0x000000010d63ca41 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 561
5   encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
6   encode-02                           0x000000010d5192fd main + 189
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:48
>>SIZE:176
>>DEPT:10
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d55db51 _ZNK3ndn5Block5parseEv + 481
2   libndn-cxx.0.7.1.dylib              0x000000010d5a5069 _ZN3ndn8MetaInfo10wireDecodeERKNS_5BlockE + 41
3   libndn-cxx.0.7.1.dylib              0x000000010d537892 _ZN3ndn4Data10wireDecodeERKNS_5BlockE + 1266
4   libndn-cxx.0.7.1.dylib              0x000000010d5393e0 _ZNK3ndn4Data10wireEncodeERNS_8encoding12EncodingImplILNS1_3TagE1EEERKNS_5BlockE + 96
5   libndn-cxx.0.7.1.dylib              0x000000010d63ca41 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 561
6   encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
7   encode-02                           0x000000010d5192fd main + 189
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:49
>>SIZE:88
>>DEPT:10
>>ADDR:
0   encode-02                           0x000000010d51b0d3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d55db51 _ZNK3ndn5Block5parseEv + 481
2   libndn-cxx.0.7.1.dylib              0x000000010d6bb6b4 _ZN3ndn13SignatureInfo10wireDecodeERKNS_5BlockENS0_4TypeE + 100
3   libndn-cxx.0.7.1.dylib              0x000000010d5378bd _ZN3ndn4Data10wireDecodeERKNS_5BlockE + 1309
4   libndn-cxx.0.7.1.dylib              0x000000010d5393e0 _ZNK3ndn4Data10wireEncodeERNS_8encoding12EncodingImplILNS1_3TagE1EEERKNS_5BlockE + 96
5   libndn-cxx.0.7.1.dylib              0x000000010d63ca41 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 561
6   encode-02                           0x000000010d519096 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 710
7   encode-02                           0x000000010d5192fd main + 189
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
```

### encode-03 (after encoding)

```text
sizeof(data) = 1016
Referred to Size = 13716

>>TIME:75
>>SIZE:2904
>>DEPT:8
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac217e5 _ZNSt3__16vectorIN3ndn5BlockENS_9allocatorIS2_EEEC2ERKS5_ + 101
2   libndn-cxx.0.7.1.dylib              0x000000010ac1963f _ZN3ndn5BlockC1ERKS0_ + 95
3   libndn-cxx.0.7.1.dylib              0x000000010abf7fb7 _ZN3ndn4DataC2ERKNS_4NameE + 71
4   encode-03                           0x000000010abd9e1a _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 106
5   encode-03                           0x000000010abda2cd main + 189
6   libdyld.dylib                       0x00007fff20544f3d start + 1
7   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:80
>>SIZE:352
>>DEPT:11
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010b39b4e4 _ZN5boost3log11v2_mt_posix13attribute_setC1Ev + 20
2   libndn-cxx.0.7.1.dylib              0x000000010ada78cc _ZN5boost3log11v2_mt_posix7sources12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEC2Ev + 44
3   libndn-cxx.0.7.1.dylib              0x000000010ada77c2 _ZN5boost3log11v2_mt_posix7sources21basic_severity_loggerINS2_12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEES8_EC2Ev + 18
4   libndn-cxx.0.7.1.dylib              0x000000010ada6b2f _ZN3ndn4util6LoggerC2EPKc + 31
5   libndn-cxx.0.7.1.dylib              0x000000010acfe779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
6   libndn-cxx.0.7.1.dylib              0x000000010acfd86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
7   encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
8   encode-03                           0x000000010abda2cd main + 189
9   libdyld.dylib                       0x00007fff20544f3d start + 1
10  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:81
>>SIZE:32
>>DEPT:10
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010b39b89a _ZN5boost3log11v2_mt_posix13attribute_set14implementation6insertENS1_14attribute_nameERKNS1_9attributeE + 266
2   libndn-cxx.0.7.1.dylib              0x000000010ada7816 _ZN5boost3log11v2_mt_posix7sources21basic_severity_loggerINS2_12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEES8_EC2Ev + 102
3   libndn-cxx.0.7.1.dylib              0x000000010ada6b2f _ZN3ndn4util6LoggerC2EPKc + 31
4   libndn-cxx.0.7.1.dylib              0x000000010acfe779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
5   libndn-cxx.0.7.1.dylib              0x000000010acfd86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
6   encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
7   encode-03                           0x000000010abda2cd main + 189
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:82
>>SIZE:32
>>DEPT:9
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010b39b89a _ZN5boost3log11v2_mt_posix13attribute_set14implementation6insertENS1_14attribute_nameERKNS1_9attributeE + 266
2   libndn-cxx.0.7.1.dylib              0x000000010ada6c40 _ZN3ndn4util6LoggerC2EPKc + 304
3   libndn-cxx.0.7.1.dylib              0x000000010acfe779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
4   libndn-cxx.0.7.1.dylib              0x000000010acfd86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
5   encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
6   encode-03                           0x000000010abda2cd main + 189
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:83
>>SIZE:48
>>DEPT:10
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010adaa1c8 _ZNSt3__112__hash_tableINS_17__hash_value_typeINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEEPN3ndn4util6LoggerEEENS_22__unordered_map_hasherIS7_SC_NS_4hashIS7_EELb1EEENS_21__unordered_map_equalIS7_SC_NS_8equal_toIS7_EELb1EEENS5_ISC_EEE16__construct_nodeIJRKS7_SB_EEENS_10unique_ptrINS_11__hash_nodeISC_PvEENS_22__hash_node_destructorINS5_IST_EEEEEEDpOT_ + 40
2   libndn-cxx.0.7.1.dylib              0x000000010ada8dcc _ZN3ndn4util7Logging13addLoggerImplERNS0_6LoggerE + 60
3   libndn-cxx.0.7.1.dylib              0x000000010ada6c79 _ZN3ndn4util6LoggerC2EPKc + 361
4   libndn-cxx.0.7.1.dylib              0x000000010acfe779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
5   libndn-cxx.0.7.1.dylib              0x000000010acfd86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
6   encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
7   encode-03                           0x000000010abda2cd main + 189
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:85
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac246f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010ac18966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010aca94bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010aca9fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010aca98c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010acb2050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010acb1c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010acb1e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010ad19b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010acfe579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010acfd86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-03                           0x000000010abda2cd main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:86
>>SIZE:11
>>DEPT:16
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac24741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010ac18966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010aca94bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010aca9fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010aca98c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010acb2050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010acb1c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010acb1e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010ad19b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010acfe579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010acfd86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-03                           0x000000010abda2cd main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:88
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac246f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010ac18966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010aca94bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010aca9fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010aca98c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010acb2050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010acb1c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010acb1e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010ad19b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010acfe579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010acfd86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-03                           0x000000010abda2cd main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:89
>>SIZE:10
>>DEPT:16
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac24741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010ac18966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010aca94bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010aca9fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010aca98c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010acb2050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010acb1c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010acb1e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010ad19b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010acfe579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010acfd86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-03                           0x000000010abda2cd main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:91
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac246f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010ac18966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010aca94bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010aca9fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010aca98c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010acb2050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010acb1c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010acb1e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010ad19b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010acfe579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010acfd86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-03                           0x000000010abda2cd main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:92
>>SIZE:15
>>DEPT:16
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac24741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010ac18966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010aca94bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010aca9fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010aca98c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010acb2050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010acb1c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010acb1e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010ad19b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010acfe579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010acfd86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-03                           0x000000010abda2cd main + 189
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:93
>>SIZE:352
>>DEPT:11
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac1ffba _ZN3ndn5Block9push_backERKS0_ + 346
2   libndn-cxx.0.7.1.dylib              0x000000010acb1c7d _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 509
3   libndn-cxx.0.7.1.dylib              0x000000010acb1e66 _ZN3ndn4NameC1EPKc + 134
4   libndn-cxx.0.7.1.dylib              0x000000010ad19b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
5   libndn-cxx.0.7.1.dylib              0x000000010acfe579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
6   libndn-cxx.0.7.1.dylib              0x000000010acfd86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
7   encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
8   encode-03                           0x000000010abda2cd main + 189
9   libdyld.dylib                       0x00007fff20544f3d start + 1
10  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:95
>>SIZE:48
>>DEPT:7
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac246f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010acfd941 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 305
3   encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
4   encode-03                           0x000000010abda2cd main + 189
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:96
>>SIZE:8800
>>DEPT:7
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac24741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010acfd941 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 305
3   encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
4   encode-03                           0x000000010abda2cd main + 189
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:110
>>SIZE:704
>>DEPT:9
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac1eb51 _ZNK3ndn5Block5parseEv + 481
2   libndn-cxx.0.7.1.dylib              0x000000010abf83ec _ZN3ndn4Data10wireDecodeERKNS_5BlockE + 76
3   libndn-cxx.0.7.1.dylib              0x000000010abfa3e0 _ZNK3ndn4Data10wireEncodeERNS_8encoding12EncodingImplILNS1_3TagE1EEERKNS_5BlockE + 96
4   libndn-cxx.0.7.1.dylib              0x000000010acfda41 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 561
5   encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
6   encode-03                           0x000000010abda2cd main + 189
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:112
>>SIZE:176
>>DEPT:10
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac1eb51 _ZNK3ndn5Block5parseEv + 481
2   libndn-cxx.0.7.1.dylib              0x000000010ac66069 _ZN3ndn8MetaInfo10wireDecodeERKNS_5BlockE + 41
3   libndn-cxx.0.7.1.dylib              0x000000010abf8892 _ZN3ndn4Data10wireDecodeERKNS_5BlockE + 1266
4   libndn-cxx.0.7.1.dylib              0x000000010abfa3e0 _ZNK3ndn4Data10wireEncodeERNS_8encoding12EncodingImplILNS1_3TagE1EEERKNS_5BlockE + 96
5   libndn-cxx.0.7.1.dylib              0x000000010acfda41 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 561
6   encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
7   encode-03                           0x000000010abda2cd main + 189
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:113
>>SIZE:88
>>DEPT:10
>>ADDR:
0   encode-03                           0x000000010abdc0a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010ac1eb51 _ZNK3ndn5Block5parseEv + 481
2   libndn-cxx.0.7.1.dylib              0x000000010ad7c6b4 _ZN3ndn13SignatureInfo10wireDecodeERKNS_5BlockENS0_4TypeE + 100
3   libndn-cxx.0.7.1.dylib              0x000000010abf88bd _ZN3ndn4Data10wireDecodeERKNS_5BlockE + 1309
4   libndn-cxx.0.7.1.dylib              0x000000010abfa3e0 _ZNK3ndn4Data10wireEncodeERNS_8encoding12EncodingImplILNS1_3TagE1EEERKNS_5BlockE + 96
5   libndn-cxx.0.7.1.dylib              0x000000010acfda41 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 561
6   encode-03                           0x000000010abda066 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
7   encode-03                           0x000000010abda2cd main + 189
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
```

### encode-01 (wire)

```text
sizeof(wire) = 88
Referred to Size = 10284

>>TIME:17
>>SIZE:352
>>DEPT:11
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010d7944e4 _ZN5boost3log11v2_mt_posix13attribute_setC1Ev + 20
2   libndn-cxx.0.7.1.dylib              0x000000010d1948cc _ZN5boost3log11v2_mt_posix7sources12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEC2Ev + 44
3   libndn-cxx.0.7.1.dylib              0x000000010d1947c2 _ZN5boost3log11v2_mt_posix7sources21basic_severity_loggerINS2_12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEES8_EC2Ev + 18
4   libndn-cxx.0.7.1.dylib              0x000000010d193b2f _ZN3ndn4util6LoggerC2EPKc + 31
5   libndn-cxx.0.7.1.dylib              0x000000010d0eb779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
6   libndn-cxx.0.7.1.dylib              0x000000010d0ea86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
7   encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
8   encode-01                           0x000000010cfcb0a5 main + 229
9   libdyld.dylib                       0x00007fff20544f3d start + 1
10  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:18
>>SIZE:32
>>DEPT:10
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010d79489a _ZN5boost3log11v2_mt_posix13attribute_set14implementation6insertENS1_14attribute_nameERKNS1_9attributeE + 266
2   libndn-cxx.0.7.1.dylib              0x000000010d194816 _ZN5boost3log11v2_mt_posix7sources21basic_severity_loggerINS2_12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEES8_EC2Ev + 102
3   libndn-cxx.0.7.1.dylib              0x000000010d193b2f _ZN3ndn4util6LoggerC2EPKc + 31
4   libndn-cxx.0.7.1.dylib              0x000000010d0eb779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
5   libndn-cxx.0.7.1.dylib              0x000000010d0ea86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
6   encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
7   encode-01                           0x000000010cfcb0a5 main + 229
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:19
>>SIZE:32
>>DEPT:9
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010d79489a _ZN5boost3log11v2_mt_posix13attribute_set14implementation6insertENS1_14attribute_nameERKNS1_9attributeE + 266
2   libndn-cxx.0.7.1.dylib              0x000000010d193c40 _ZN3ndn4util6LoggerC2EPKc + 304
3   libndn-cxx.0.7.1.dylib              0x000000010d0eb779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
4   libndn-cxx.0.7.1.dylib              0x000000010d0ea86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
5   encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
6   encode-01                           0x000000010cfcb0a5 main + 229
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:20
>>SIZE:48
>>DEPT:10
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d1971c8 _ZNSt3__112__hash_tableINS_17__hash_value_typeINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEEPN3ndn4util6LoggerEEENS_22__unordered_map_hasherIS7_SC_NS_4hashIS7_EELb1EEENS_21__unordered_map_equalIS7_SC_NS_8equal_toIS7_EELb1EEENS5_ISC_EEE16__construct_nodeIJRKS7_SB_EEENS_10unique_ptrINS_11__hash_nodeISC_PvEENS_22__hash_node_destructorINS5_IST_EEEEEEDpOT_ + 40
2   libndn-cxx.0.7.1.dylib              0x000000010d195dcc _ZN3ndn4util7Logging13addLoggerImplERNS0_6LoggerE + 60
3   libndn-cxx.0.7.1.dylib              0x000000010d193c79 _ZN3ndn4util6LoggerC2EPKc + 361
4   libndn-cxx.0.7.1.dylib              0x000000010d0eb779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
5   libndn-cxx.0.7.1.dylib              0x000000010d0ea86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
6   encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
7   encode-01                           0x000000010cfcb0a5 main + 229
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:22
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d0116f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010d005966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010d0964bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010d096fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010d0968c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010d09f050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010d09ec71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010d09ee66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010d106b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010d0eb579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010d0ea86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
13  encode-01                           0x000000010cfcb0a5 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:23
>>SIZE:11
>>DEPT:16
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d011741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010d005966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010d0964bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010d096fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010d0968c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010d09f050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010d09ec71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010d09ee66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010d106b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010d0eb579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010d0ea86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
13  encode-01                           0x000000010cfcb0a5 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:25
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d0116f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010d005966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010d0964bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010d096fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010d0968c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010d09f050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010d09ec71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010d09ee66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010d106b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010d0eb579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010d0ea86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
13  encode-01                           0x000000010cfcb0a5 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:26
>>SIZE:10
>>DEPT:16
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d011741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010d005966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010d0964bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010d096fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010d0968c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010d09f050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010d09ec71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010d09ee66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010d106b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010d0eb579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010d0ea86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
13  encode-01                           0x000000010cfcb0a5 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:28
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d0116f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010d005966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010d0964bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010d096fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010d0968c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010d09f050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010d09ec71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010d09ee66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010d106b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010d0eb579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010d0ea86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
13  encode-01                           0x000000010cfcb0a5 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:29
>>SIZE:15
>>DEPT:16
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d011741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010d005966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x000000010d0964bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x000000010d096fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x000000010d0968c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010d09f050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010d09ec71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010d09ee66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x000000010d106b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x000000010d0eb579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010d0ea86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
13  encode-01                           0x000000010cfcb0a5 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:30
>>SIZE:352
>>DEPT:11
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d00cfba _ZN3ndn5Block9push_backERKS0_ + 346
2   libndn-cxx.0.7.1.dylib              0x000000010d09ec7d _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 509
3   libndn-cxx.0.7.1.dylib              0x000000010d09ee66 _ZN3ndn4NameC1EPKc + 134
4   libndn-cxx.0.7.1.dylib              0x000000010d106b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
5   libndn-cxx.0.7.1.dylib              0x000000010d0eb579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
6   libndn-cxx.0.7.1.dylib              0x000000010d0ea86b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
7   encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
8   encode-01                           0x000000010cfcb0a5 main + 229
9   libdyld.dylib                       0x00007fff20544f3d start + 1
10  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:32
>>SIZE:48
>>DEPT:7
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d0116f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x000000010d0ea941 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 305
3   encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
4   encode-01                           0x000000010cfcb0a5 main + 229
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:33
>>SIZE:8800
>>DEPT:7
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d011741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x000000010d0ea941 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 305
3   encode-01                           0x000000010cfcaea6 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
4   encode-01                           0x000000010cfcb0a5 main + 229
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:51
>>SIZE:440
>>DEPT:7
>>ADDR:
0   encode-01                           0x000000010cfcd0b3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x000000010d00e7e5 _ZNSt3__16vectorIN3ndn5BlockENS_9allocatorIS2_EEEC2ERKS5_ + 101
2   libndn-cxx.0.7.1.dylib              0x000000010d00663f _ZN3ndn5BlockC1ERKS0_ + 95
3   encode-01                           0x000000010cfcaec9 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 713
4   encode-01                           0x000000010cfcb0a5 main + 229
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
```

### encode-02 (wire)

```text
sizeof(wire) = 88
Referred to Size = 10284

>>TIME:17
>>SIZE:352
>>DEPT:11
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x0000000100c3d4e4 _ZN5boost3log11v2_mt_posix13attribute_setC1Ev + 20
2   libndn-cxx.0.7.1.dylib              0x00000001006408cc _ZN5boost3log11v2_mt_posix7sources12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEC2Ev + 44
3   libndn-cxx.0.7.1.dylib              0x00000001006407c2 _ZN5boost3log11v2_mt_posix7sources21basic_severity_loggerINS2_12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEES8_EC2Ev + 18
4   libndn-cxx.0.7.1.dylib              0x000000010063fb2f _ZN3ndn4util6LoggerC2EPKc + 31
5   libndn-cxx.0.7.1.dylib              0x0000000100597779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
6   libndn-cxx.0.7.1.dylib              0x000000010059686b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
7   encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
8   encode-02                           0x0000000100475095 main + 229
9   libdyld.dylib                       0x00007fff20544f3d start + 1
10  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:18
>>SIZE:32
>>DEPT:10
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x0000000100c3d89a _ZN5boost3log11v2_mt_posix13attribute_set14implementation6insertENS1_14attribute_nameERKNS1_9attributeE + 266
2   libndn-cxx.0.7.1.dylib              0x0000000100640816 _ZN5boost3log11v2_mt_posix7sources21basic_severity_loggerINS2_12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEES8_EC2Ev + 102
3   libndn-cxx.0.7.1.dylib              0x000000010063fb2f _ZN3ndn4util6LoggerC2EPKc + 31
4   libndn-cxx.0.7.1.dylib              0x0000000100597779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
5   libndn-cxx.0.7.1.dylib              0x000000010059686b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
6   encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
7   encode-02                           0x0000000100475095 main + 229
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:19
>>SIZE:32
>>DEPT:9
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x0000000100c3d89a _ZN5boost3log11v2_mt_posix13attribute_set14implementation6insertENS1_14attribute_nameERKNS1_9attributeE + 266
2   libndn-cxx.0.7.1.dylib              0x000000010063fc40 _ZN3ndn4util6LoggerC2EPKc + 304
3   libndn-cxx.0.7.1.dylib              0x0000000100597779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
4   libndn-cxx.0.7.1.dylib              0x000000010059686b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
5   encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
6   encode-02                           0x0000000100475095 main + 229
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:20
>>SIZE:48
>>DEPT:10
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x00000001006431c8 _ZNSt3__112__hash_tableINS_17__hash_value_typeINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEEPN3ndn4util6LoggerEEENS_22__unordered_map_hasherIS7_SC_NS_4hashIS7_EELb1EEENS_21__unordered_map_equalIS7_SC_NS_8equal_toIS7_EELb1EEENS5_ISC_EEE16__construct_nodeIJRKS7_SB_EEENS_10unique_ptrINS_11__hash_nodeISC_PvEENS_22__hash_node_destructorINS5_IST_EEEEEEDpOT_ + 40
2   libndn-cxx.0.7.1.dylib              0x0000000100641dcc _ZN3ndn4util7Logging13addLoggerImplERNS0_6LoggerE + 60
3   libndn-cxx.0.7.1.dylib              0x000000010063fc79 _ZN3ndn4util6LoggerC2EPKc + 361
4   libndn-cxx.0.7.1.dylib              0x0000000100597779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
5   libndn-cxx.0.7.1.dylib              0x000000010059686b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
6   encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
7   encode-02                           0x0000000100475095 main + 229
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:22
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x00000001004bd6f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x00000001004b1966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x00000001005424bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x0000000100542fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x00000001005428c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010054b050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010054ac71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010054ae66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x00000001005b2b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x0000000100597579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010059686b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-02                           0x0000000100475095 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:23
>>SIZE:11
>>DEPT:16
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x00000001004bd741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x00000001004b1966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x00000001005424bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x0000000100542fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x00000001005428c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010054b050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010054ac71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010054ae66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x00000001005b2b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x0000000100597579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010059686b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-02                           0x0000000100475095 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:25
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x00000001004bd6f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x00000001004b1966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x00000001005424bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x0000000100542fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x00000001005428c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010054b050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010054ac71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010054ae66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x00000001005b2b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x0000000100597579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010059686b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-02                           0x0000000100475095 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:26
>>SIZE:10
>>DEPT:16
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x00000001004bd741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x00000001004b1966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x00000001005424bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x0000000100542fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x00000001005428c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010054b050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010054ac71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010054ae66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x00000001005b2b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x0000000100597579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010059686b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-02                           0x0000000100475095 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:28
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x00000001004bd6f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x00000001004b1966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x00000001005424bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x0000000100542fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x00000001005428c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010054b050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010054ac71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010054ae66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x00000001005b2b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x0000000100597579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010059686b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-02                           0x0000000100475095 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:29
>>SIZE:15
>>DEPT:16
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x00000001004bd741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x00000001004b1966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x00000001005424bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x0000000100542fe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x00000001005428c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x000000010054b050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x000000010054ac71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x000000010054ae66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x00000001005b2b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x0000000100597579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x000000010059686b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
13  encode-02                           0x0000000100475095 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:30
>>SIZE:352
>>DEPT:11
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x00000001004b8fba _ZN3ndn5Block9push_backERKS0_ + 346
2   libndn-cxx.0.7.1.dylib              0x000000010054ac7d _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 509
3   libndn-cxx.0.7.1.dylib              0x000000010054ae66 _ZN3ndn4NameC1EPKc + 134
4   libndn-cxx.0.7.1.dylib              0x00000001005b2b21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
5   libndn-cxx.0.7.1.dylib              0x0000000100597579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
6   libndn-cxx.0.7.1.dylib              0x000000010059686b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
7   encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
8   encode-02                           0x0000000100475095 main + 229
9   libdyld.dylib                       0x00007fff20544f3d start + 1
10  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:32
>>SIZE:48
>>DEPT:7
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x00000001004bd6f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x0000000100596941 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 305
3   encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
4   encode-02                           0x0000000100475095 main + 229
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:33
>>SIZE:8800
>>DEPT:7
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x00000001004bd741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x0000000100596941 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 305
3   encode-02                           0x0000000100474e96 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 694
4   encode-02                           0x0000000100475095 main + 229
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:51
>>SIZE:440
>>DEPT:7
>>ADDR:
0   encode-02                           0x00000001004770a3 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x00000001004ba7e5 _ZNSt3__16vectorIN3ndn5BlockENS_9allocatorIS2_EEEC2ERKS5_ + 101
2   libndn-cxx.0.7.1.dylib              0x00000001004b263f _ZN3ndn5BlockC1ERKS0_ + 95
3   encode-02                           0x0000000100474eb9 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 729
4   encode-02                           0x0000000100475095 main + 229
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
```

### encode-03 (wire)

```text
sizeof(wire) = 88
Referred to Size = 10284

>>TIME:81
>>SIZE:352
>>DEPT:11
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x00000001072804e4 _ZN5boost3log11v2_mt_posix13attribute_setC1Ev + 20
2   libndn-cxx.0.7.1.dylib              0x0000000106c8a8cc _ZN5boost3log11v2_mt_posix7sources12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEC2Ev + 44
3   libndn-cxx.0.7.1.dylib              0x0000000106c8a7c2 _ZN5boost3log11v2_mt_posix7sources21basic_severity_loggerINS2_12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEES8_EC2Ev + 18
4   libndn-cxx.0.7.1.dylib              0x0000000106c89b2f _ZN3ndn4util6LoggerC2EPKc + 31
5   libndn-cxx.0.7.1.dylib              0x0000000106be1779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
6   libndn-cxx.0.7.1.dylib              0x0000000106be086b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
7   encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
8   encode-03                           0x0000000106abe065 main + 229
9   libdyld.dylib                       0x00007fff20544f3d start + 1
10  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:82
>>SIZE:32
>>DEPT:10
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010728089a _ZN5boost3log11v2_mt_posix13attribute_set14implementation6insertENS1_14attribute_nameERKNS1_9attributeE + 266
2   libndn-cxx.0.7.1.dylib              0x0000000106c8a816 _ZN5boost3log11v2_mt_posix7sources21basic_severity_loggerINS2_12basic_loggerIcNS2_18severity_logger_mtIN3ndn4util8LogLevelEEENS2_18multi_thread_modelINS1_3aux14light_rw_mutexEEEEES8_EC2Ev + 102
3   libndn-cxx.0.7.1.dylib              0x0000000106c89b2f _ZN3ndn4util6LoggerC2EPKc + 31
4   libndn-cxx.0.7.1.dylib              0x0000000106be1779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
5   libndn-cxx.0.7.1.dylib              0x0000000106be086b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
6   encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
7   encode-03                           0x0000000106abe065 main + 229
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:83
>>SIZE:32
>>DEPT:9
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libboost_log-mt.dylib               0x000000010728089a _ZN5boost3log11v2_mt_posix13attribute_set14implementation6insertENS1_14attribute_nameERKNS1_9attributeE + 266
2   libndn-cxx.0.7.1.dylib              0x0000000106c89c40 _ZN3ndn4util6LoggerC2EPKc + 304
3   libndn-cxx.0.7.1.dylib              0x0000000106be1779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
4   libndn-cxx.0.7.1.dylib              0x0000000106be086b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
5   encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
6   encode-03                           0x0000000106abe065 main + 229
7   libdyld.dylib                       0x00007fff20544f3d start + 1
8   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:84
>>SIZE:48
>>DEPT:10
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x0000000106c8d1c8 _ZNSt3__112__hash_tableINS_17__hash_value_typeINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEEPN3ndn4util6LoggerEEENS_22__unordered_map_hasherIS7_SC_NS_4hashIS7_EELb1EEENS_21__unordered_map_equalIS7_SC_NS_8equal_toIS7_EELb1EEENS5_ISC_EEE16__construct_nodeIJRKS7_SB_EEENS_10unique_ptrINS_11__hash_nodeISC_PvEENS_22__hash_node_destructorINS5_IST_EEEEEEDpOT_ + 40
2   libndn-cxx.0.7.1.dylib              0x0000000106c8bdcc _ZN3ndn4util7Logging13addLoggerImplERNS0_6LoggerE + 60
3   libndn-cxx.0.7.1.dylib              0x0000000106c89c79 _ZN3ndn4util6LoggerC2EPKc + 361
4   libndn-cxx.0.7.1.dylib              0x0000000106be1779 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2825
5   libndn-cxx.0.7.1.dylib              0x0000000106be086b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
6   encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
7   encode-03                           0x0000000106abe065 main + 229
8   libdyld.dylib                       0x00007fff20544f3d start + 1
9   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:86
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x0000000106b076f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x0000000106afb966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x0000000106b8c4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x0000000106b8cfe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x0000000106b8c8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x0000000106b95050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x0000000106b94c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x0000000106b94e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x0000000106bfcb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x0000000106be1579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x0000000106be086b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
13  encode-03                           0x0000000106abe065 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:87
>>SIZE:11
>>DEPT:16
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x0000000106b07741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x0000000106afb966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x0000000106b8c4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x0000000106b8cfe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x0000000106b8c8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x0000000106b95050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x0000000106b94c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x0000000106b94e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x0000000106bfcb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x0000000106be1579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x0000000106be086b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
13  encode-03                           0x0000000106abe065 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:89
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x0000000106b076f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x0000000106afb966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x0000000106b8c4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x0000000106b8cfe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x0000000106b8c8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x0000000106b95050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x0000000106b94c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x0000000106b94e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x0000000106bfcb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x0000000106be1579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x0000000106be086b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
13  encode-03                           0x0000000106abe065 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:90
>>SIZE:10
>>DEPT:16
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x0000000106b07741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x0000000106afb966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x0000000106b8c4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x0000000106b8cfe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x0000000106b8c8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x0000000106b95050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x0000000106b94c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x0000000106b94e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x0000000106bfcb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x0000000106be1579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x0000000106be086b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
13  encode-03                           0x0000000106abe065 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:92
>>SIZE:48
>>DEPT:16
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x0000000106b076f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x0000000106afb966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x0000000106b8c4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x0000000106b8cfe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x0000000106b8c8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x0000000106b95050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x0000000106b94c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x0000000106b94e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x0000000106bfcb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x0000000106be1579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x0000000106be086b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
13  encode-03                           0x0000000106abe065 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:93
>>SIZE:15
>>DEPT:16
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x0000000106b07741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x0000000106afb966 _ZN3ndn8encoding15makeBinaryBlockEjPKhm + 54
3   libndn-cxx.0.7.1.dylib              0x0000000106b8c4bc _ZN3ndn4name9ComponentC2EjPKhm + 28
4   libndn-cxx.0.7.1.dylib              0x0000000106b8cfe1 _ZN3ndn4nameL20parseUriEscapedValueEjPKcm + 337
5   libndn-cxx.0.7.1.dylib              0x0000000106b8c8c9 _ZN3ndn4name9Component17fromEscapedStringERKNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEE + 233
6   libndn-cxx.0.7.1.dylib              0x0000000106b95050 _ZN3ndn4name9Component17fromEscapedStringEPKcmm + 400
7   libndn-cxx.0.7.1.dylib              0x0000000106b94c71 _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 497
8   libndn-cxx.0.7.1.dylib              0x0000000106b94e66 _ZN3ndn4NameC1EPKc + 134
9   libndn-cxx.0.7.1.dylib              0x0000000106bfcb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
10  libndn-cxx.0.7.1.dylib              0x0000000106be1579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
11  libndn-cxx.0.7.1.dylib              0x0000000106be086b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
12  encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
13  encode-03                           0x0000000106abe065 main + 229
14  libdyld.dylib                       0x00007fff20544f3d start + 1
15  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:94
>>SIZE:352
>>DEPT:11
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x0000000106b02fba _ZN3ndn5Block9push_backERKS0_ + 346
2   libndn-cxx.0.7.1.dylib              0x0000000106b94c7d _ZN3ndn4NameC2ENSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEE + 509
3   libndn-cxx.0.7.1.dylib              0x0000000106b94e66 _ZN3ndn4NameC1EPKc + 134
4   libndn-cxx.0.7.1.dylib              0x0000000106bfcb21 _ZN3ndn8security11SigningInfo23getDigestSha256IdentityEv + 65
5   libndn-cxx.0.7.1.dylib              0x0000000106be1579 _ZN3ndn8security2v28KeyChain20prepareSignatureInfoERKNS0_11SigningInfoE + 2313
6   libndn-cxx.0.7.1.dylib              0x0000000106be086b _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 91
7   encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
8   encode-03                           0x0000000106abe065 main + 229
9   libdyld.dylib                       0x00007fff20544f3d start + 1
10  ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:96
>>SIZE:48
>>DEPT:7
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x0000000106b076f1 _ZN3ndn8encoding7EncoderC2Emm + 33
2   libndn-cxx.0.7.1.dylib              0x0000000106be0941 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 305
3   encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
4   encode-03                           0x0000000106abe065 main + 229
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:97
>>SIZE:8800
>>DEPT:7
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x0000000106b07741 _ZN3ndn8encoding7EncoderC2Emm + 113
2   libndn-cxx.0.7.1.dylib              0x0000000106be0941 _ZN3ndn8security2v28KeyChain4signERNS_4DataERKNS0_11SigningInfoE + 305
3   encode-03                           0x0000000106abde66 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 678
4   encode-03                           0x0000000106abe065 main + 229
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
>>TIME:115
>>SIZE:440
>>DEPT:7
>>ADDR:
0   encode-03                           0x0000000106ac0073 _ZN7monitor11MonitorImpl6getMemEm + 51
1   libndn-cxx.0.7.1.dylib              0x0000000106b047e5 _ZNSt3__16vectorIN3ndn5BlockENS_9allocatorIS2_EEEC2ERKS5_ + 101
2   libndn-cxx.0.7.1.dylib              0x0000000106afc63f _ZN3ndn5BlockC1ERKS0_ + 95
3   encode-03                           0x0000000106abde89 _Z6encodeNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEE + 713
4   encode-03                           0x0000000106abe065 main + 229
5   libdyld.dylib                       0x00007fff20544f3d start + 1
6   ???                                 0x0000000000000001 0x0 + 1
<<
```

Memory - cpp-template
---------------------

### encode-01

```text
sizeof(data) = 208
Referred to Size = 373
sizeof(wire) = 40
Content Size = 189

>>TIME:1
>>SIZE:120
>>DEPT:1
>>ADDR:
0   encode-01                           0x00000001051002a8 _Znwm + 104
<<
>>TIME:2
>>SIZE:8
>>DEPT:1
>>ADDR:
0   encode-01                           0x00000001051002a8 _Znwm + 104
<<
>>TIME:3
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-01                           0x00000001051002a8 _Znwm + 104
<<
>>TIME:4
>>SIZE:6
>>DEPT:1
>>ADDR:
0   encode-01                           0x00000001051002a8 _Znwm + 104
<<
>>TIME:5
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-01                           0x00000001051002a8 _Znwm + 104
<<
>>TIME:6
>>SIZE:12
>>DEPT:1
>>ADDR:
0   encode-01                           0x00000001051002a8 _Znwm + 104
<<
>>TIME:7
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-01                           0x00000001051002a8 _Znwm + 104
<<
>>TIME:8
>>SIZE:7
>>DEPT:1
>>ADDR:
0   encode-01                           0x00000001051002a8 _Znwm + 104
<<
>>TIME:9
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-01                           0x00000001051002a8 _Znwm + 104
<<
>>TIME:10
>>SIZE:100
>>DEPT:1
>>ADDR:
0   encode-01                           0x00000001051002a8 _Znwm + 104
<<
>>TIME:11
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-01                           0x00000001051002a8 _Znwm + 104
<<
```


### encode-02

```text
sizeof(data) = 208
Referred to Size = 4273
sizeof(wire) = 40
Content Size = 4093

>>TIME:1
>>SIZE:120
>>DEPT:1
>>ADDR:
0   encode-02                           0x000000010e8362a8 _Znwm + 104
<<
>>TIME:2
>>SIZE:8
>>DEPT:1
>>ADDR:
0   encode-02                           0x000000010e8362a8 _Znwm + 104
<<
>>TIME:3
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-02                           0x000000010e8362a8 _Znwm + 104
<<
>>TIME:4
>>SIZE:6
>>DEPT:1
>>ADDR:
0   encode-02                           0x000000010e8362a8 _Znwm + 104
<<
>>TIME:5
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-02                           0x000000010e8362a8 _Znwm + 104
<<
>>TIME:6
>>SIZE:12
>>DEPT:1
>>ADDR:
0   encode-02                           0x000000010e8362a8 _Znwm + 104
<<
>>TIME:7
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-02                           0x000000010e8362a8 _Znwm + 104
<<
>>TIME:8
>>SIZE:7
>>DEPT:1
>>ADDR:
0   encode-02                           0x000000010e8362a8 _Znwm + 104
<<
>>TIME:9
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-02                           0x000000010e8362a8 _Znwm + 104
<<
>>TIME:10
>>SIZE:4000
>>DEPT:1
>>ADDR:
0   encode-02                           0x000000010e8362a8 _Znwm + 104
<<
>>TIME:11
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-02                           0x000000010e8362a8 _Znwm + 104
<<
```

### encode-03

```text
sizeof(data) = 208
Referred to Size = 2293
sizeof(wire) = 40
Content Size = 189

>>TIME:1
>>SIZE:1320
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:2
>>SIZE:8
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:3
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:4
>>SIZE:6
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:5
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:6
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:7
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:8
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:9
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:10
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:11
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:12
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:13
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:14
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:15
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:16
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:17
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:18
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:19
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:20
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:21
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:22
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:23
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:24
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:25
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:26
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:27
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:28
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:29
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:30
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:31
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:32
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:33
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:34
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:35
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:36
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:37
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:38
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:39
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:40
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:41
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:42
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:43
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:44
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:45
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:46
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:47
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:48
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:49
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:50
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:51
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:52
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:53
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:54
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:55
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:56
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:57
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:58
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:59
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:60
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:61
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:62
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:63
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:64
>>SIZE:3
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:65
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:66
>>SIZE:12
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:67
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:68
>>SIZE:7
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:69
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:70
>>SIZE:10
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
>>TIME:71
>>SIZE:24
>>DEPT:1
>>ADDR:
0   encode-03                           0x00000001071f3278 _Znwm + 104
<<
```
