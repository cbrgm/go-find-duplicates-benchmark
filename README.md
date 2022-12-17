# go-find-duplicates-benchmark

This repository provides benchmarks to answer the question what's the most performant way to find duplicates present in a given slice of elements.

* `map[string]bool` ?
* `map[string]interface{}` ?
* `map[string]struct{}` ?

Let's find it out! :D

## So what's this about?

This code is a benchmark test for comparing the performance of three different versions of the `findDuplicates` function. The three versions are `findDuplicatesTypeToBool`, `findDuplicatesTypeToInterface`, and `findDuplicatesTypeToStruct`.

 The benchmark test functions in `main_test.go` each call the respective `findDuplicates` function with a slice of strings that has a large number of elements and some identical duplicates. The benchmark test measures the performance of each findDuplicates function by running it a large number of times.

To run the benchmark test, you can use the go test command and pass the -bench flag:

```bash
➜  go-items-present-benchmark go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/cbrgm/go-find-duplicates-benchmark
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz

BenchmarkFindDuplicatesTypeToBool-8                    5         222490558 ns/op
BenchmarkFindDuplicatesTypeToInterface-8               4         276779110 ns/op
BenchmarkFindDuplicatesTypeToStruct-8                  5         217530765 ns/op

PASS
ok      github.com/cbrgm/go-find-duplicates-benchmark   8.067s
```

The winner is:

```go
func findDuplicatesTypeToStruct(strings []string) []string {
	seenStrings := make(map[string]struct{})
	var duplicates []string
	for _, str := range strings {
		if _, ok := seenStrings[str]; ok {
			duplicates = append(duplicates, str)
		} else {
			seenStrings[str] = struct{}{}
		}
	}
	return duplicates
}
```

using `make(map[string]struct{})`.
