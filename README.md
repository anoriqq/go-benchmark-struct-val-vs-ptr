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
❯ go test -bench . -benchmem
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
