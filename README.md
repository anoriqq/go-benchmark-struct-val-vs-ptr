# go-benchmark-struct-val-vs-ptr

myStruct mem size: 8 * 1 = 8 bytes
```
$ go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/anoriqq/go-benchmark-struct-val-vs-ptr
cpu: 13th Gen Intel(R) Core(TM) i7-1365U
BenchmarkVal_ShortLife-12               1000000000               0.3622 ns/op          0 B/op          0 allocs/op
BenchmarkPtr_ShortLife-12               1000000000               0.3706 ns/op          0 B/op          0 allocs/op
BenchmarkVal_LongLife-12                227742208                6.495 ns/op          49 B/op          0 allocs/op
BenchmarkPtr_LongLife-12                31672814                41.99 ns/op           48 B/op          1 allocs/op
BenchmarkVal_LongLife_NoGrow-12         913061098                2.450 ns/op           8 B/op          0 allocs/op
BenchmarkPtr_LongLife_NoGrow-12         89300623                11.80 ns/op           16 B/op          1 allocs/op
PASS
ok      github.com/anoriqq/go-benchmark-struct-val-vs-ptr       8.737s
```

myStruct mem size: 8 * 2 = 16 bytes
```
$ go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/anoriqq/go-benchmark-struct-val-vs-ptr
cpu: 13th Gen Intel(R) Core(TM) i7-1365U
BenchmarkVal_ShortLife-12               1000000000               0.3998 ns/op          0 B/op          0 allocs/op
BenchmarkPtr_ShortLife-12               1000000000               0.3682 ns/op          0 B/op          0 allocs/op
BenchmarkVal_LongLife-12                96190101                12.89 ns/op           80 B/op          0 allocs/op
BenchmarkPtr_LongLife-12                22077771                49.12 ns/op           62 B/op          1 allocs/op
BenchmarkVal_LongLife_NoGrow-12         411377925                3.431 ns/op          16 B/op          0 allocs/op
BenchmarkPtr_LongLife_NoGrow-12         74971660                16.20 ns/op           24 B/op          1 allocs/op
PASS
ok      github.com/anoriqq/go-benchmark-struct-val-vs-ptr       7.086s
```

myStruct mem size: 8 * 4 = 32 bytes
```
$ go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/anoriqq/go-benchmark-struct-val-vs-ptr
cpu: 13th Gen Intel(R) Core(TM) i7-1365U
BenchmarkVal_ShortLife-12               1000000000               0.3813 ns/op          0 B/op          0 allocs/op
BenchmarkPtr_ShortLife-12               1000000000               0.3823 ns/op          0 B/op          0 allocs/op
BenchmarkVal_LongLife-12                42141897                24.29 ns/op          170 B/op          0 allocs/op
BenchmarkPtr_LongLife-12                21428671                56.16 ns/op           80 B/op          1 allocs/op
BenchmarkVal_LongLife_NoGrow-12         200596273                6.546 ns/op          32 B/op          0 allocs/op
BenchmarkPtr_LongLife_NoGrow-12         63249404                17.81 ns/op           40 B/op          1 allocs/op
PASS
ok      github.com/anoriqq/go-benchmark-struct-val-vs-ptr       7.173s
```

myStruct mem size: 8 * 8 = 64 bytes
```
$ go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/anoriqq/go-benchmark-struct-val-vs-ptr
cpu: 13th Gen Intel(R) Core(TM) i7-1365U
BenchmarkVal_ShortLife-12               1000000000               0.3632 ns/op          0 B/op          0 allocs/op
BenchmarkPtr_ShortLife-12               1000000000               0.3619 ns/op          0 B/op          0 allocs/op
BenchmarkVal_LongLife-12                25365835                42.92 ns/op          353 B/op          0 allocs/op
BenchmarkPtr_LongLife-12                27244147                38.38 ns/op          111 B/op          1 allocs/op
BenchmarkVal_LongLife_NoGrow-12         100000000               13.74 ns/op           64 B/op          0 allocs/op
BenchmarkPtr_LongLife_NoGrow-12         41581140                25.39 ns/op           72 B/op          1 allocs/op
PASS
ok      github.com/anoriqq/go-benchmark-struct-val-vs-ptr       6.450s
```

myStruct mem size: 8 * 16 = 128 bytes
```
$ go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/anoriqq/go-benchmark-struct-val-vs-ptr
cpu: 13th Gen Intel(R) Core(TM) i7-1365U
BenchmarkVal_ShortLife-12               1000000000               0.3945 ns/op          0 B/op          0 allocs/op
BenchmarkPtr_ShortLife-12               1000000000               0.3989 ns/op          0 B/op          0 allocs/op
BenchmarkVal_LongLife-12                13036196               108.5 ns/op           773 B/op          0 allocs/op
BenchmarkPtr_LongLife-12                24967234                43.53 ns/op          169 B/op          1 allocs/op
BenchmarkVal_LongLife_NoGrow-12         63160264                68.60 ns/op          128 B/op          0 allocs/op
BenchmarkPtr_LongLife_NoGrow-12         22647225                44.41 ns/op          136 B/op          1 allocs/op
PASS
ok      github.com/anoriqq/go-benchmark-struct-val-vs-ptr       10.134s
```

myStruct mem size: 8 * 32 = 256 bytes
```
$ go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/anoriqq/go-benchmark-struct-val-vs-ptr
cpu: 13th Gen Intel(R) Core(TM) i7-1365U
BenchmarkVal_ShortLife-12               1000000000               0.3800 ns/op          0 B/op          0 allocs/op
BenchmarkPtr_ShortLife-12               1000000000               0.3819 ns/op          0 B/op          0 allocs/op
BenchmarkVal_LongLife-12                 7561839               157.3 ns/op          1333 B/op          0 allocs/op
BenchmarkPtr_LongLife-12                21921075                54.09 ns/op          303 B/op          1 allocs/op
BenchmarkVal_LongLife_NoGrow-12         32120821                64.79 ns/op          256 B/op          0 allocs/op
BenchmarkPtr_LongLife_NoGrow-12         24415581                50.87 ns/op          264 B/op          1 allocs/op
PASS
ok      github.com/anoriqq/go-benchmark-struct-val-vs-ptr       9.172s
```
