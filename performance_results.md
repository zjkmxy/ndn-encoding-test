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
  - Pypy 3.7
- Method: 5 times, take average.
- Library Versions
  - ndn-cxx: [dbe645f](https://github.com/named-data/ndn-cxx/commit/dbe645f879e6906945a07935b5da344f8801afd4)
  - YaNFD: [c32e9ea](https://github.com/zjkmxy/YaNFD/commit/c32e9ea87b2f8d8805d5299b3347f5e498020260)
  - python-ndn: [5244a8e](https://github.com/named-data/python-ndn/commit/5244a8e2d4f0a1e6a0562683960cf3d36f83add9)

Summary
-------

### Performance Time

Total time for `10^6` packets.

|              | Enc 1       | Enc 2       | Enc 3       | Dec 1      | Dec 2      | Dec 3       |
|--------------|-------------|-------------|-------------|------------|------------|-------------|
| N Packets    |     1000000 |             |             |            |            |             |
| Total Bytes  |       189MB |      4093MB |       189MB |      189MB |     4093MB |       189MB |
| Name Length  |           3 |           3 |          33 |          3 |          3 |          33 |
| YaNFD        |  4.90±0.02s | 19.51±0.03s | 15.15±0.08s | 2.46±0.02s | 2.61±0.08s |    8.4±0.2s |
| Reflect      |  3.80±0.01s | 14.43±0.02s |  5.83±0.03s | 3.27±0.01s | 3.24±0.01s |  4.99±0.02s |
| CodeGen      |  1.13±0.01s | 11.89±0.03s |  3.14±0.02s | 0.48±0.01s | 0.56±0.01s |  1.93±0.02s |
| ndn-cxx      |    6.4±0.1s |   15.1±0.2s |   25.8±0.3s | 1.12±0.01s | 1.13±0.01s |  2.68±0.02s |
| C++ template |  1.06±0.01s |  9.86±0.08s |  3.83±0.02s | 0.39±0.01s | 0.47±0.01s |  1.33±0.01s |
| python-ndn   |  118.0±0.6s |  141.5±0.9s |  221.7±0.7s |  37.5±0.3s |  39.7±0.2s |   54.7±0.3s |
| (pypy3)      |   23.0±0.2s |   44.9±0.4s |   41.2±0.6s | 3.96±0.04s | 4.02±0.04s |  4.61±0.02s |

### Time delta

This table shows the computed incremental time per payload and name component.
Time in this table is per one packet, so 1s in the last table becomes 1μs.

The two empty cells in the Payload column are because the number is less than SEM and considered as "too small to measure".
Should be considered as 0.

|              | Enc Base (μs) | Payload (μs/kB) | NameComp (μs/1) | Dec Base (μs) | Payload (μs/kB) | NameComp (μs/1) |
|--------------|---------------|-----------------|-----------------|---------------|-----------------|-----------------|
| YaNFD        |     3.50±0.06 |       3.75±0.01 |     0.342±0.003 |     1.86±0.03 |       0.04±0.03 |     0.198±0.007 |
| Reflect      |     3.32±0.02 |     2.726±0.006 |     0.068±0.002 |     3.07±0.02 |          ±0.004 |     0.057±0.001 |
| CodeGen      |     0.65±0.02 |     2.759±0.009 |     0.067±0.001 |     0.33±0.02 |     0.021±0.004 |     0.048±0.001 |
| ndn-cxx      |       4.3±0.2 |       2.23±0.06 |       0.64±0.02 |     0.96±0.02 |          ±0.004 |     0.052±0.001 |
| C++ template |     0.56±0.02 |       2.25±0.03 |     0.092±0.001 |     0.29±0.02 |     0.021±0.004 |     0.031±0.001 |
| python-ndn   |     107.0±0.7 |         6.0±0.3 |       3.46±0.03 |      35.7±0.3 |         0.6±0.1 |       0.57±0.02 |
| (pypy3)      |      20.6±0.3 |         5.6±0.2 |       0.61±0.03 |     3.89±0.05 |       0.02±0.02 |     0.022±0.001 |

### Lines of Code

This table counts the number of lines of code (LOC).
In this column, `def` column counts the lines to define specified data structure,
`alg` column indicates the lines for encoding and decoding functions,
`etc` column indicates all other LOC, especially property accessors.
Definitions of TLV type numbers, test cases are excluded.
For `Data`, lines of code of inner data structures are exlcuded, but the computation of signature value (the piece of code calling SHA256 signer) is included.
The `Library` column only counts the LOC of necessary data structures and functions for TLV encoding,
everything related to NDN packet specification is excluded.
The `Library` column is counted by [cloc](https://github.com/AlDanial/cloc), others by hand.

Also, note that `ndn-cxx`, `python-ndn` and `YaNFD` are published libraries,
so it is natural for them to have more LOC than others, which are more or less toy code.

|              | Library | MetaInfo (def) | MetaInfo (alg) | MetaInfo (etc) | Data (def) | Data (alg) | Data (etc) |
|--------------|---------|----------------|----------------|----------------|------------|------------|------------|
| YaNFD        |    1052 |              6 |             76 |             67 |         10 |        129 |         98 |
| Reflect      |     422 |              5 |              6 |              0 |          7 |         41 |          0 |
| CodeGen      |     666 |              8 |             12 |              0 |         12 |         30 |          0 |
| ndn-cxx      |    1890 |              7 |             59 |            125 |          9 |        128 |        427 |
| C++ template |     653 |              9 |              6 |              0 |         13 |         32 |          0 |
| python-ndn   |     909 |              4 |              0 |             12 |         16 |         53 |         11 |


### Memory Usage

|                       | Enc 1 Data | Enc 1 Wire | Enc 2 Data | Enc 2 Wire | Enc 3 Data | Enc 3 Wire |
|-----------------------|------------|------------|------------|------------|------------|------------|
| Bare Wire             |          - |        189 |          - |       4093 |          - |        189 |
| YaNFD (before sign)   |        575 |        213 |       4475 |       4117 |       2255 |        213 |
| YaNFD (after sign)    |       1719 |          - |      13425 |          - |       4569 |          - |
| Reflect               |        341 |        213 |       4241 |       4117 |       1061 |        213 |
| CodeGen               |        341 |        213 |       4241 |       4117 |       1061 |        213 |
| C++ template          |        581 |        229 |       4481 |       4133 |       2501 |        229 |
| ndn-cxx (before sign) |       1655 |      10372 |       5557 |      10372 |       5735 |      10372 |
| ndn-cxx (after sign)  |      12092 |          - |      12092 |          - |      14732 |          - |
| python-ndn            |         ?? |         ?? |         ?? |         ?? |         ?? |         ?? |

Note:
- In ndn-cxx and YaNFD, the total memory overheads of Data before and after encoding+signing are different,
  so there are two rows.
  For all other implementations, the data is before encoding+signing.
- In ndn-cxx, every wire is allocated 8800B (maximum) regardless of the real wire size.
  This is because ndn-cxx cannot handle variable-sized signature whose size is **unknown before encoding**.
- There is no quick way to measure memory overhead of Python.
  Personally I believe it would use more than any other implementation, including ndn-cxx.
- Go reflection and code generation have the same number because they share the definition of data structures.
  The algorithm does not affect memory overhead.
- The size of wire is the same for all Golang implementations, because the type `[]byte` is internally defined by Golang.

Approximation formula:
This is **not meant to be precise**.
Let *X* be the payload size, *Y* be the number of name components, *Z* be the length of wire as a binary string.
Then,
- Reflect & CodeGen
  - Data = 169 + X + 24Y
  - Wire = 24 + Z
- C++ template
  - Data = 289 + X + 64Y
  - Wire = 40 + Z
- YaNFD
  - Data (before signing) = 307 + X + 56Y
  - Wire = 24 + Z
  - Data (after signing) = 1134 + 3X + 95Y
- ndn-cxx
  - Data (before signing) = 1147 + X + 136Y
  - Wire = 1572 + 8800
  - Data (after signing) = 1456 + (1572 + 8800) + 88Y
