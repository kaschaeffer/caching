Demo of how false sharing can cause a multi-threaded program to be *slower* than a single-threaded program. Running the benchmark on my laptop gives,
```
caching$ go test github.com/kaschaeffer/caching -bench=. -cpu=1
testing: warning: no tests to run
BenchmarkRandArrayConcurrency1              2000           1037069 ns/op
BenchmarkRandArrayConcurrency2              1000           1477427 ns/op
BenchmarkRandArrayConcurrency8               300           4111978 ns/op
BenchmarkRandArrayConcurrency16              300           4987083 ns/op
BenchmarkRandArrayConcurrency32              300           5269392 ns/op
BenchmarkSmartRandArrayConcurrency1         2000           1055584 ns/op
BenchmarkSmartRandArrayConcurrency2         1000           1088241 ns/op
BenchmarkSmartRandArrayConcurrency8         2000           1089109 ns/op
BenchmarkSmartRandArrayConcurrency16        2000           1085773 ns/op
BenchmarkSmartRandArrayConcurrency32        1000           1068439 ns/op
PASS
ok      github.com/kaschaeffer/caching  18.795s
```
