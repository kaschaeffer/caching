Demo of how false sharing can cause a multi-threaded program to be *slower* than a single-threaded program. Running the benchmark on my laptop gives,
```
caching$ go test github.com/kaschaeffer/caching -bench=. -cpu=8
testing: warning: no tests to run
BenchmarkRandArrayConcurrency1-8                    2000           1031416 ns/op
BenchmarkRandArrayConcurrency2-8                    1000           1183386 ns/op
BenchmarkRandArrayConcurrency8-8                    1000           1888705 ns/op
BenchmarkRandArrayConcurrency16-8                   1000           2073087 ns/op
BenchmarkRandArrayConcurrency32-8                   1000           2099112 ns/op
BenchmarkSmartRandArrayConcurrency1-8               2000           1077855 ns/op
BenchmarkSmartRandArrayConcurrency2-8               2000            890880 ns/op
BenchmarkSmartRandArrayConcurrency8-8               2000            882588 ns/op
BenchmarkSmartRandArrayConcurrency16-8              2000            832951 ns/op
BenchmarkSmartRandArrayConcurrency32-8              2000            872729 ns/op
PASS
ok      github.com/kaschaeffer/caching  19.787s
```
