Performance Data
================

Environments
------------

- CPU: Intel Core i7-5820K 3.3GHz 6 Core 12 Threads
- Memory: 64 GiB
- Operating System: Ubuntu 18.04
- Compilers
  - gcc 10.3.0
  - go 1.16
  - Python 3.6.8
- Method: 5 times, take average.
- Library Versions
  - ndn-cxx: [dbe645f](https://github.com/named-data/ndn-cxx/commit/dbe645f879e6906945a07935b5da344f8801afd4)
  - YaNFD: [b4baed1](https://github.com/zjkmxy/YaNFD/commit/b4baed13e82856b7ab02be7d6625757401e8f2c1)
  - python-ndn: [5244a8e](https://github.com/named-data/python-ndn/commit/5244a8e2d4f0a1e6a0562683960cf3d36f83add9)

Summary
-------

|              | Enc 1       | Enc 2       | Enc 3       | Dec 1      | Dec 2      | Dec 3       |
|--------------|-------------|-------------|-------------|------------|------------|-------------|
| N Packets    |      100000 |             |             |            |            |             |
| Total Bytes  |       189MB |      4093MB |       189MB |      189MB |     4093MB |       189MB |
| Name Length  |           3 |           3 |          33 |          3 |          3 |          33 |
| YaNFD        |  4.90±0.02s | 19.51±0.03s | 15.15±0.08s | 3.17±0.01s | 4.64±0.07s | 13.70±0.06s |
| Reflect      |  3.80±0.01s | 14.43±0.02s |  5.83±0.03s | 3.27±0.01s | 3.24±0.01s |  4.99±0.02s |
| CodeGen      |  1.13±0.01s | 11.89±0.03s |  3.14±0.02s | 0.48±0.01s | 0.56±0.01s |  1.93±0.02s |
| ndn-cxx      |    6.4±0.1s |   15.1±0.2s |   25.8±0.3s | 1.12±0.01s | 1.13±0.01s |  2.68±0.02s |
| C++ template |  1.06±0.01s |  9.86±0.08s |  3.83±0.02s | 0.39±0.01s | 0.47±0.01s |  1.33±0.01s |
| python-ndn   |  118.0±0.6s |  141.5±0.9s |  221.7±0.7s |  37.5±0.3s |  39.7±0.2s |   54.7±0.3s |


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
block: 		3.189096366s
reflection: 	3.291573173s
codegen: 	460.77955ms
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		4.439581405s
reflection: 	3.243882034s
codegen: 	556.047842ms
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		13.653388577s
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
block: 		3.149444212s
reflection: 	3.236745777s
codegen: 	470.479219ms
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		4.653946675s
reflection: 	3.251319296s
codegen: 	562.28719ms
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		13.867353821s
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
block: 		3.167998289s
reflection: 	3.27619214s
codegen: 	477.900987ms
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		4.639645857s
reflection: 	3.237357446s
codegen: 	566.837059ms
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		13.580961856s
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
block: 		3.158894297s
reflection: 	3.277122611s
codegen: 	481.221132ms
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		4.863815989s
reflection: 	3.246431148s
codegen: 	557.789032ms
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		13.775798189s
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
block: 		3.171213212s
reflection: 	3.285752864s
codegen: 	493.424117ms
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		4.595272187s
reflection: 	3.245891018s
codegen: 	557.266169ms
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		13.599677625s
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