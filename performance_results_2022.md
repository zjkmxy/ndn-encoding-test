Performance Data
================

Environments
------------

- CPU: AMD EPYC 7702P 64-Core Processor 2GHz
- Memory: 256 GiB
- Operating System: Ubuntu 20.04
- Compilers
  - gcc 11.1.0
  - go 1.18.8
  - Python 3.10.8
  - Pypy 3.9 v7.3.9
- Method: 5 times, take average. Encode does not include signing.
- Library Versions
  - ndn-cxx: [0.8.0-40-g87208f96](https://github.com/named-data/ndn-cxx/commit/87208f964af86ebad17cb2ab287f1f957a622679)
  - YaNFD: [v1.1.1](https://github.com/named-data/YaNFD/commit/1b359373e000d5550da5e465ce346a79128d7f24)
  - go-ndn: [v0.0.1](https://github.com/zjkmxy/go-ndn/commit/7f5d9528206552d1310e42ff971981a25a58bbf2)
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
|--------------|-------------|-------------|-------------|------------|------------|-------------|
| YaNFD        |    4.8±0.2s |   10.1±0.2s |   18.7±0.3s |   3.0±0.2s | 3.30±0.08s |   11.8±0.3s |
| go-ndn       |  1.24±0.04s |  1.24±0.03s |  3.68±0.04s | 1.52±0.04s |   1.7±0.2s |  2.70±0.04s |
| Reflect      |  3.24±0.09s |    4.5±0.2s |  4.78±0.08s | 3.58±0.06s | 3.83±0.05s |  4.87±0.06s |
| CodeGen      |  0.41±0.01s |  1.28±0.02s |  1.71±0.02s | 0.49±0.03s | 0.56±0.04s |  1.68±0.03s |
|--------------|-------------|-------------|-------------|------------|------------|-------------|
| ndn-cxx      |  4.71±0.02s |  5.02±0.02s |   24.7±0.2s | 0.98±0.01s | 0.99±0.01s |  2.20±0.02s |
| C++ template |  0.65±0.01s |  0.86±0.01s |  3.85±0.03s | 0.44±0.01s | 0.55±0.01s |  1.44±0.01s |
|--------------|-------------|-------------|-------------|------------|------------|-------------|
| python-ndn   |   62.4±0.3s |   63.8±0.3s |  149.7±0.4s |  29.7±0.3s |  31.8±0.3s |   44.1±0.2s |
| (pypy3)      | 11.93±0.09s |   14.2±0.2s |   30.3±0.7s | 3.97±0.02s | 4.11±0.05s |  4.45±0.03s |

### Time delta

This table shows the computed incremental time per payload and name component.
Time in this table is per one packet, so 1s in the last table becomes 1μs.

The two empty cells in the Payload column are because the number is less than SEM and considered as "too small to measure".
Should be considered as 0.

|              | Enc Base (μs) | Payload (μs/kB) | NameComp (μs/1) | Dec Base (μs) | Payload (μs/kB) | NameComp (μs/1) |
|--------------|---------------|-----------------|-----------------|---------------|-----------------|-----------------|
| YaNFD        |       3.2±0.2 |       1.35±0.06 |       0.46±0.02 |       2.1±0.2 |       0.08±0.04 |       0.29±0.02 |
| go-ndn       |     1.00±0.04 |           ±0.02 |       0.08±0.01 |     1.39±0.04 |       0.05±0.03 |       0.04±0.01 |
| Reflect      |     3.02±0.09 |       0.33±0.04 |       0.05±0.01 |     3.43±0.06 |       0.07±0.02 |       0.04±0.01 |
| CodeGen      |     0.24±0.01 |       0.22±0.01 |       0.04±0.01 |     0.36±0.03 |           ±0.02 |       0.04±0.01 |
|--------------|---------------|-----------------|-----------------|---------------|-----------------|-----------------|
| ndn-cxx      |     2.70±0.03 |       0.08±0.01 |       0.67±0.01 |     0.85±0.01 |           ±0.01 |       0.04±0.01 |
| C++ template |     0.32±0.01 |       0.05±0.01 |       0.11±0.01 |     0.33±0.01 |       0.03±0.01 |       0.03±0.01 |
|--------------|---------------|-----------------|-----------------|---------------|-----------------|-----------------|
| python-ndn   |      53.6±0.3 |       0.36±0.09 |       2.91±0.02 |      28.1±0.3 |       0.54±0.09 |       0.48±0.02 |
| (pypy3)      |      20.6±0.3 |         5.6±0.2 |       0.61±0.03 |     3.91±0.02 |       0.04±0.02 |       0.02±0.01 |
