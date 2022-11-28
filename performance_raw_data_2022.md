Performance Raw Data
====================

Raw Data - Go
-------------

### T1
```text
========== ENCODING ==========
#1: 1000000 x 100B
block: 		4.67663692s
reflection: 	3.157657245s
codegen: 	395.029188ms
go-ndn: 	1.190615811s
=== Total Bytes: 157000000 ===

#2: 1000000 x 4000B
block: 		10.165547017s
reflection: 	4.388770015s
codegen: 	1.293207849s
go-ndn: 	1.225761454s
=== Total Bytes: 4061000000 ===

#3: 1000000 x 10B, very long name
block: 		18.731163322s
reflection: 	4.693285616s
codegen: 	1.699032847s
go-ndn: 	3.667540899s
=== Total Bytes: 157000000 ===


========== DECODING ==========
#1: 1000000 x 100B
block: 		2.860157795s
reflection: 	3.500351787s
codegen: 	480.816891ms
go-ndn: 	1.487439062s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		3.333955151s
reflection: 	3.817095418s
codegen: 	584.847617ms
go-ndn: 	1.501810026s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		11.288632567s
reflection: 	4.806839129s
codegen: 	1.624469683s
go-ndn: 	2.61328961s
=== Total Bytes: 189000000 ===
```

### T2
```text
========== ENCODING ==========
#1: 1000000 x 100B
block: 		4.582129134s
reflection: 	3.145796293s
codegen: 	412.242271ms
go-ndn: 	1.196418429s
=== Total Bytes: 157000000 ===

#2: 1000000 x 4000B
block: 		9.718630792s
reflection: 	4.326103729s
codegen: 	1.25041934s
go-ndn: 	1.22986069s
=== Total Bytes: 4061000000 ===

#3: 1000000 x 10B, very long name
block: 		17.96821182s
reflection: 	4.699999557s
codegen: 	1.698533895s
go-ndn: 	3.646105196s
=== Total Bytes: 157000000 ===


========== DECODING ==========
#1: 1000000 x 100B
block: 		2.861431878s
reflection: 	3.530654818s
codegen: 	469.351441ms
go-ndn: 	1.477805705s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		3.255373885s
reflection: 	3.798165811s
codegen: 	469.669904ms
go-ndn: 	1.563610921s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		11.538812214s
reflection: 	4.81665073s
codegen: 	1.690652069s
go-ndn: 	2.658300768s
=== Total Bytes: 189000000 ===
```

### T3
```text
========== ENCODING ==========
#1: 1000000 x 100B
block: 		5.474662919s
reflection: 	3.577971846s
codegen: 	450.367407ms
go-ndn: 	1.373183654s
=== Total Bytes: 157000000 ===

#2: 1000000 x 4000B
block: 		10.664963821s
reflection: 	5.012251785s
codegen: 	1.312554619s
go-ndn: 	1.345747491s
=== Total Bytes: 4061000000 ===

#3: 1000000 x 10B, very long name
block: 		19.740851257s
reflection: 	5.090469479s
codegen: 	1.757999314s
go-ndn: 	3.824067977s
=== Total Bytes: 157000000 ===


========== DECODING ==========
#1: 1000000 x 100B
block: 		3.391460814s
reflection: 	3.777783975s
codegen: 	570.674198ms
go-ndn: 	1.653739367s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		3.557387109s
reflection: 	4.007941466s
codegen: 	674.626103ms
go-ndn: 	2.043325283s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		12.96831358s
reflection: 	5.098410121s
codegen: 	1.778530534s
go-ndn: 	2.796506643s
=== Total Bytes: 189000000 ===
```

### T4
```text
========== ENCODING ==========
#1: 1000000 x 100B
block: 		4.709141327s
reflection: 	3.168203245s
codegen: 	409.810874ms
go-ndn: 	1.239392843s
=== Total Bytes: 157000000 ===

#2: 1000000 x 4000B
block: 		10.120164771s
reflection: 	4.372149564s
codegen: 	1.299950136s
go-ndn: 	1.185275807s
=== Total Bytes: 4061000000 ===

#3: 1000000 x 10B, very long name
block: 		18.533588836s
reflection: 	4.748805008s
codegen: 	1.67949514s
go-ndn: 	3.661470608s
=== Total Bytes: 157000000 ===


========== DECODING ==========
#1: 1000000 x 100B
block: 		2.831127777s
reflection: 	3.543941292s
codegen: 	456.835738ms
go-ndn: 	1.509674352s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		3.176879953s
reflection: 	3.757457115s
codegen: 	472.995216ms
go-ndn: 	1.579113511s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		11.534729715s
reflection: 	4.802152714s
codegen: 	1.618808941s
go-ndn: 	2.726003154s
=== Total Bytes: 189000000 ===
```

### T5
```text
========== ENCODING ==========
#1: 1000000 x 100B
block: 		4.756385213s
reflection: 	3.151173713s
codegen: 	407.405585ms
go-ndn: 	1.204561764s
=== Total Bytes: 157000000 ===

#2: 1000000 x 4000B
block: 		9.893059001s
reflection: 	4.58520191s
codegen: 	1.220856638s
go-ndn: 	1.202472278s
=== Total Bytes: 4061000000 ===

#3: 1000000 x 10B, very long name
block: 		18.481247954s
reflection: 	4.670820359s
codegen: 	1.731437394s
go-ndn: 	3.601708637s
=== Total Bytes: 157000000 ===


========== DECODING ==========
#1: 1000000 x 100B
block: 		2.909725209s
reflection: 	3.523114509s
codegen: 	453.671962ms
go-ndn: 	1.477389402s
=== Total Bytes: 189000000 ===

#2: 1000000 x 4000B
block: 		3.155751748s
reflection: 	3.775435955s
codegen: 	590.404335ms
go-ndn: 	1.804836655s
=== Total Bytes: 4093000000 ===

#3: 1000000 x 10B, very long name
block: 		11.707822883s
reflection: 	4.835008836s
codegen: 	1.675320903s
go-ndn: 	2.699001139s
=== Total Bytes: 189000000 ===
```

Raw Data - ndn-cxx
------------------

### T1
```text
time ./encode-01
4.76user 0.00system 0:04.79elapsed 99%CPU (0avgtext+0avgdata 10820maxresident)k
1912inputs+0outputs (12major+448minor)pagefaults 0swaps
time ./encode-02
5.06user 0.00system 0:05.07elapsed 99%CPU (0avgtext+0avgdata 10916maxresident)k
0inputs+0outputs (0major+455minor)pagefaults 0swaps
time ./encode-03
24.40user 0.00system 0:24.41elapsed 99%CPU (0avgtext+0avgdata 10968maxresident)k
0inputs+0outputs (0major+456minor)pagefaults 0swaps
time ./decode-01
0.97user 0.00system 0:00.97elapsed 99%CPU (0avgtext+0avgdata 10784maxresident)k
0inputs+0outputs (0major+449minor)pagefaults 0swaps
time ./decode-02
0.98user 0.00system 0:00.99elapsed 99%CPU (0avgtext+0avgdata 10876maxresident)k
0inputs+0outputs (0major+451minor)pagefaults 0swaps
time ./decode-03
2.22user 0.00system 0:02.23elapsed 99%CPU (0avgtext+0avgdata 10816maxresident)k
0inputs+0outputs (0major+451minor)pagefaults 0swaps
```

### T2
```text
time ./encode-01
4.72user 0.00system 0:04.73elapsed 99%CPU (0avgtext+0avgdata 10872maxresident)k
0inputs+0outputs (0major+452minor)pagefaults 0swaps
time ./encode-02
4.99user 0.00system 0:04.99elapsed 99%CPU (0avgtext+0avgdata 10732maxresident)k
0inputs+0outputs (0major+453minor)pagefaults 0swaps
time ./encode-03
24.61user 0.00system 0:24.62elapsed 99%CPU (0avgtext+0avgdata 10924maxresident)k
0inputs+0outputs (0major+455minor)pagefaults 0swaps
time ./decode-01
0.98user 0.00system 0:00.99elapsed 100%CPU (0avgtext+0avgdata 10784maxresident)k
0inputs+0outputs (0major+448minor)pagefaults 0swaps
time ./decode-02
0.98user 0.00system 0:00.99elapsed 99%CPU (0avgtext+0avgdata 10736maxresident)k
0inputs+0outputs (0major+449minor)pagefaults 0swaps
time ./decode-03
2.19user 0.00system 0:02.19elapsed 99%CPU (0avgtext+0avgdata 10772maxresident)k
0inputs+0outputs (0major+452minor)pagefaults 0swaps
```

### T3
```text
time ./encode-01
4.67user 0.00system 0:04.68elapsed 99%CPU (0avgtext+0avgdata 10788maxresident)k
0inputs+0outputs (0major+452minor)pagefaults 0swaps
time ./encode-02
5.07user 0.00system 0:05.07elapsed 99%CPU (0avgtext+0avgdata 11000maxresident)k
0inputs+0outputs (0major+452minor)pagefaults 0swaps
time ./encode-03
24.76user 0.00system 0:24.77elapsed 99%CPU (0avgtext+0avgdata 10796maxresident)k
0inputs+0outputs (0major+454minor)pagefaults 0swaps
time ./decode-01
0.99user 0.00system 0:01.00elapsed 99%CPU (0avgtext+0avgdata 10772maxresident)k
0inputs+0outputs (0major+450minor)pagefaults 0swaps
time ./decode-02
0.99user 0.00system 0:00.99elapsed 99%CPU (0avgtext+0avgdata 10576maxresident)k
0inputs+0outputs (0major+446minor)pagefaults 0swaps
time ./decode-03
2.21user 0.00system 0:02.21elapsed 100%CPU (0avgtext+0avgdata 10820maxresident)k
0inputs+0outputs (0major+453minor)pagefaults 0swaps
```

### T4
```text
time ./encode-01
4.69user 0.00system 0:04.69elapsed 99%CPU (0avgtext+0avgdata 10996maxresident)k
0inputs+0outputs (0major+452minor)pagefaults 0swaps
time ./encode-02
5.00user 0.00system 0:05.00elapsed 99%CPU (0avgtext+0avgdata 10896maxresident)k
0inputs+0outputs (0major+454minor)pagefaults 0swaps
time ./encode-03
24.32user 0.00system 0:24.32elapsed 99%CPU (0avgtext+0avgdata 10924maxresident)k
0inputs+0outputs (0major+454minor)pagefaults 0swaps
time ./decode-01
0.98user 0.00system 0:00.99elapsed 99%CPU (0avgtext+0avgdata 10568maxresident)k
0inputs+0outputs (0major+445minor)pagefaults 0swaps
time ./decode-02
1.00user 0.00system 0:01.00elapsed 100%CPU (0avgtext+0avgdata 10792maxresident)k
0inputs+0outputs (0major+450minor)pagefaults 0swaps
time ./decode-03
2.23user 0.00system 0:02.23elapsed 99%CPU (0avgtext+0avgdata 10876maxresident)k
0inputs+0outputs (0major+451minor)pagefaults 0swaps
```

### T5
```text
time ./encode-01
4.72user 0.01system 0:04.80elapsed 98%CPU (0avgtext+0avgdata 10960maxresident)k
0inputs+0outputs (0major+451minor)pagefaults 0swaps
time ./encode-02
4.97user 0.00system 0:04.97elapsed 99%CPU (0avgtext+0avgdata 10972maxresident)k
0inputs+0outputs (0major+454minor)pagefaults 0swaps
time ./encode-03
25.43user 0.01system 0:25.46elapsed 99%CPU (0avgtext+0avgdata 10972maxresident)k
0inputs+0outputs (0major+455minor)pagefaults 0swaps
time ./decode-01
0.96user 0.00system 0:00.97elapsed 99%CPU (0avgtext+0avgdata 10868maxresident)k
0inputs+0outputs (0major+450minor)pagefaults 0swaps
time ./decode-02
0.98user 0.00system 0:00.99elapsed 100%CPU (0avgtext+0avgdata 10772maxresident)k
0inputs+0outputs (0major+451minor)pagefaults 0swaps
time ./decode-03
2.16user 0.00system 0:02.16elapsed 99%CPU (0avgtext+0avgdata 10844maxresident)k
0inputs+0outputs (0major+453minor)pagefaults 0swaps
```

Raw Data - cpp-template
-----------------------

### T1
```text
time ./encode-01
0.65user 0.00system 0:00.65elapsed 100%CPU (0avgtext+0avgdata 4076maxresident)k
0inputs+0outputs (0major+202minor)pagefaults 0swaps
time ./encode-02
0.86user 0.00system 0:00.87elapsed 100%CPU (0avgtext+0avgdata 4056maxresident)k
0inputs+0outputs (0major+204minor)pagefaults 0swaps
time ./encode-03
3.93user 0.00system 0:03.93elapsed 99%CPU (0avgtext+0avgdata 4048maxresident)k
0inputs+0outputs (0major+202minor)pagefaults 0swaps
time ./decode-01
0.44user 0.00system 0:00.44elapsed 99%CPU (0avgtext+0avgdata 2696maxresident)k
0inputs+0outputs (0major+117minor)pagefaults 0swaps
time ./decode-02
0.55user 0.00system 0:00.55elapsed 99%CPU (0avgtext+0avgdata 2656maxresident)k
0inputs+0outputs (0major+118minor)pagefaults 0swaps
time ./decode-03
1.45user 0.00system 0:01.45elapsed 99%CPU (0avgtext+0avgdata 2640maxresident)k
0inputs+0outputs (0major+118minor)pagefaults 0swaps
```

### T2
```text
time ./encode-01
0.64user 0.00system 0:00.64elapsed 100%CPU (0avgtext+0avgdata 4072maxresident)k
0inputs+0outputs (0major+201minor)pagefaults 0swaps
time ./encode-02
0.87user 0.00system 0:00.87elapsed 99%CPU (0avgtext+0avgdata 4196maxresident)k
0inputs+0outputs (0major+206minor)pagefaults 0swaps
time ./encode-03
3.80user 0.00system 0:03.81elapsed 99%CPU (0avgtext+0avgdata 4076maxresident)k
0inputs+0outputs (0major+203minor)pagefaults 0swaps
time ./decode-01
0.44user 0.00system 0:00.45elapsed 99%CPU (0avgtext+0avgdata 2828maxresident)k
0inputs+0outputs (0major+119minor)pagefaults 0swaps
time ./decode-02
0.55user 0.00system 0:00.55elapsed 99%CPU (0avgtext+0avgdata 2780maxresident)k
0inputs+0outputs (0major+118minor)pagefaults 0swaps
time ./decode-03
1.43user 0.00system 0:01.43elapsed 99%CPU (0avgtext+0avgdata 2644maxresident)k
0inputs+0outputs (0major+119minor)pagefaults 0swaps
```

### T3
```text
time ./encode-01
0.65user 0.00system 0:00.65elapsed 100%CPU (0avgtext+0avgdata 4120maxresident)k
0inputs+0outputs (0major+204minor)pagefaults 0swaps
time ./encode-02
0.86user 0.00system 0:00.87elapsed 99%CPU (0avgtext+0avgdata 4080maxresident)k
0inputs+0outputs (0major+205minor)pagefaults 0swaps
time ./encode-03
3.84user 0.00system 0:03.84elapsed 99%CPU (0avgtext+0avgdata 4072maxresident)k
0inputs+0outputs (0major+201minor)pagefaults 0swaps
time ./decode-01
0.44user 0.00system 0:00.44elapsed 100%CPU (0avgtext+0avgdata 2700maxresident)k
0inputs+0outputs (0major+118minor)pagefaults 0swaps
time ./decode-02
0.53user 0.00system 0:00.54elapsed 99%CPU (0avgtext+0avgdata 2724maxresident)k
0inputs+0outputs (0major+119minor)pagefaults 0swaps
time ./decode-03
1.45user 0.00system 0:01.45elapsed 99%CPU (0avgtext+0avgdata 2760maxresident)k
0inputs+0outputs (0major+120minor)pagefaults 0swaps
```

### T4
```text
time ./encode-01
0.66user 0.00system 0:00.66elapsed 100%CPU (0avgtext+0avgdata 4140maxresident)k
0inputs+0outputs (0major+204minor)pagefaults 0swaps
time ./encode-02
0.87user 0.00system 0:00.87elapsed 100%CPU (0avgtext+0avgdata 4276maxresident)k
0inputs+0outputs (0major+207minor)pagefaults 0swaps
time ./encode-03
3.83user 0.00system 0:03.83elapsed 99%CPU (0avgtext+0avgdata 4076maxresident)k
0inputs+0outputs (0major+202minor)pagefaults 0swaps
time ./decode-01
0.44user 0.00system 0:00.44elapsed 99%CPU (0avgtext+0avgdata 2636maxresident)k
0inputs+0outputs (0major+116minor)pagefaults 0swaps
time ./decode-02
0.56user 0.00system 0:00.56elapsed 99%CPU (0avgtext+0avgdata 2784maxresident)k
0inputs+0outputs (0major+120minor)pagefaults 0swaps
time ./decode-03
1.43user 0.00system 0:01.43elapsed 99%CPU (0avgtext+0avgdata 2924maxresident)k
0inputs+0outputs (0major+121minor)pagefaults 0swaps
```

### T5
```text
time ./encode-01
0.65user 0.00system 0:00.65elapsed 99%CPU (0avgtext+0avgdata 4392maxresident)k
0inputs+0outputs (0major+207minor)pagefaults 0swaps
time ./encode-02
0.86user 0.00system 0:00.86elapsed 99%CPU (0avgtext+0avgdata 4080maxresident)k
0inputs+0outputs (0major+203minor)pagefaults 0swaps
time ./encode-03
3.84user 0.00system 0:03.84elapsed 99%CPU (0avgtext+0avgdata 4076maxresident)k
0inputs+0outputs (0major+203minor)pagefaults 0swaps
time ./decode-01
0.44user 0.00system 0:00.44elapsed 100%CPU (0avgtext+0avgdata 2692maxresident)k
0inputs+0outputs (0major+118minor)pagefaults 0swaps
time ./decode-02
0.54user 0.00system 0:00.54elapsed 100%CPU (0avgtext+0avgdata 2788maxresident)k
0inputs+0outputs (0major+119minor)pagefaults 0swaps
time ./decode-03
1.45user 0.00system 0:01.45elapsed 100%CPU (0avgtext+0avgdata 2748maxresident)k
0inputs+0outputs (0major+120minor)pagefaults 0swaps
```

Raw Data - python-ndn
---------------------

### T1
```text
Total time: 62.2852 seconds
Total time: 64.2509 seconds
Total time: 148.7726 seconds
Total time: 28.9768 seconds
Total time: 31.7006 seconds
Total time: 43.9430 seconds
```

### T2
```text
Total time: 62.1275 seconds
Total time: 63.8741 seconds
Total time: 150.2971 seconds
Total time: 29.4303 seconds
Total time: 31.0668 seconds
Total time: 44.3359 seconds
```

### T3
```text
Total time: 62.2678 seconds
Total time: 63.3428 seconds
Total time: 148.6627 seconds
Total time: 29.4934 seconds
Total time: 32.0789 seconds
Total time: 43.4546 seconds
```

### T4
```text
Total time: 63.2147 seconds
Total time: 64.3971 seconds
Total time: 149.9341 seconds
Total time: 30.1219 seconds
Total time: 31.7342 seconds
Total time: 44.4142 seconds
```

### T5
```text
Total time: 62.0192 seconds
Total time: 63.1404 seconds
Total time: 150.6032 seconds
Total time: 30.3081 seconds
Total time: 32.2927 seconds
Total time: 44.4706 seconds
```

Raw Data - python-ndn (pypy3)
---------------------

### T1
```text
pypy3.9 encode-01.py && pypy3.9 encode-02.py && pypy3.9 encode-03.py && pypy3.9 decode-01.py && pypy3.9 decode-02.py && pypy3.9 decode-03.py
Total time: 11.9140 seconds
Total time: 14.0051 seconds
Total time: 31.4962 seconds
Total time: 3.9848 seconds
Total time: 3.9903 seconds
Total time: 4.4731 seconds
```

### T2
```text
Total time: 11.7807 seconds
Total time: 14.5594 seconds
Total time: 31.0613 seconds
Total time: 3.9843 seconds
Total time: 4.1010 seconds
Total time: 4.4938 seconds
```

### T3
```text
Total time: 11.8982 seconds
Total time: 14.1013 seconds
Total time: 31.6415 seconds
Total time: 3.9489 seconds
Total time: 4.1406 seconds
Total time: 4.5027 seconds
```

### T4
```text
Total time: 12.2620 seconds
Total time: 14.6467 seconds
Total time: 28.8061 seconds
Total time: 3.9287 seconds
Total time: 4.2668 seconds
Total time: 4.4095 seconds
```

### T5
```text
Total time: 11.7832 seconds
Total time: 13.8933 seconds
Total time: 28.6914 seconds
Total time: 3.9962 seconds
Total time: 4.0430 seconds
Total time: 4.3505 seconds
```
