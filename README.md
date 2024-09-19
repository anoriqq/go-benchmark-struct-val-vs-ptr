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

size = 10 ^ 8
```
$ go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/anoriqq/go-benchmark-struct-val-vs-ptr
cpu: 13th Gen Intel(R) Core(TM) i7-1365U
BenchmarkVal_ShortLife-12               1000000000               0.3730 ns/op          0 B/op          0 allocs/op
BenchmarkPtr_ShortLife-12               1000000000               0.3981 ns/op          0 B/op          0 allocs/op
BenchmarkVal_LongLife-12                139280827               12.24 ns/op           86 B/op          0 allocs/op
BenchmarkPtr_LongLife-12                26889417                44.81 ns/op           64 B/op          1 allocs/op
BenchmarkVal_LongLife_NoGrow-12         407656084                3.247 ns/op          16 B/op          0 allocs/op
BenchmarkPtr_LongLife_NoGrow-12         80364201                14.97 ns/op           24 B/op          1 allocs/op
PASS
ok      github.com/anoriqq/go-benchmark-struct-val-vs-ptr       9.295s
```

size = 10 ^ 32
```
❯ go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/anoriqq/go-benchmark-struct-val-vs-ptr
cpu: 13th Gen Intel(R) Core(TM) i7-1365U
BenchmarkVal_ShortLife-12               1000000000               0.3757 ns/op          0 B/op          0 allocs/op
BenchmarkPtr_ShortLife-12               1000000000               0.3618 ns/op          0 B/op          0 allocs/op
BenchmarkVal_LongLife-12                 5001573               234.7 ns/op          1790 B/op          0 allocs/op
BenchmarkPtr_LongLife-12                19648651                59.02 ns/op          394 B/op          1 allocs/op
BenchmarkVal_LongLife_NoGrow-12         24927114               127.5 ns/op           336 B/op          0 allocs/op
BenchmarkPtr_LongLife_NoGrow-12         15164799                67.95 ns/op          360 B/op          1 allocs/op
PASS
ok      github.com/anoriqq/go-benchmark-struct-val-vs-ptr       9.186s
```
