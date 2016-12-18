Demo of how false sharing can cause a multi-threaded program to be *slower* than a single-threaded program. Running the benchmark on my laptop gives,
```
golang$ go test github.com/kaschaeffer/caching -bench=. -cpu=8
testing: warning: no tests to run
BenchmarkRandArray1Core-8                   1000           1043888 ns/op
BenchmarkRandArray2Core-8                   1000           1170034 ns/op
BenchmarkRandArray8Core-8                   1000           1875210 ns/op
BenchmarkRandArray16Core-8                  1000           2004669 ns/op
BenchmarkRandArray32Core-8                  1000           2094006 ns/op
BenchmarkSmartRandArray1Core-8              2000           1074827 ns/op
BenchmarkSmartRandArray2Core-8              2000            898244 ns/op
BenchmarkSmartRandArray8Core-8              2000            830723 ns/op
BenchmarkSmartRandArray16Core-8             2000            849677 ns/op
BenchmarkSmartRandArray32Core-8             2000            871643 ns/op
PASS
ok      github.com/kaschaeffer/caching  18.580s
```
