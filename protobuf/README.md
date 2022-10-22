# Protobuf benchmark

## Usage

1. Replace the `Message` definition in `proto/benchmark.proto` with yours
2. Initialize the `Message` variables in each `*/benchmark_test.go`
3. Run `make`

More details can be found in [Makefile](./Makefile).

## Benchmark results

Here are benchmark results of the example:

```go
goos: linux
goarch: amd64
pkg: gogoprotobuf
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkProto3Marshal-12       22648688                51.27 ns/op           24 B/op          1 allocs/op
BenchmarkProto3Unmarshal-12      9473576               122.6 ns/op           112 B/op          3 allocs/op
BenchmarkProto3Clone-12          1545654               777.0 ns/op           176 B/op         11 allocs/op
```

```go
goos: linux
goarch: amd64
pkg: golangprotobuf
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkProto3Marshal-12       10488739               112.0 ns/op            24 B/op          1 allocs/op
BenchmarkProto3Unmarshal-12      8455938               143.6 ns/op           112 B/op          3 allocs/op
BenchmarkProto3Clone-12          5705482               210.8 ns/op           104 B/op          2 allocs/op
```

```go
goos: linux
goarch: amd64
pkg: golangprotobuf/v2
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkProto3Marshal-12        8762001               136.0 ns/op            24 B/op          1 allocs/op
BenchmarkProto3Unmarshal-12      5313693               223.8 ns/op           112 B/op          3 allocs/op
BenchmarkProto3Clone-12          6948472               170.6 ns/op           104 B/op          2 allocs/op
```

```go
goos: linux
goarch: amd64
pkg: vtprotobuf
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkProto3Marshal-12               27638942                43.21 ns/op           24 B/op          1 allocs/op
BenchmarkProto3Unmarshal-12             19143121                60.41 ns/op           16 B/op          2 allocs/op
BenchmarkProto3Marshal_Pool-12          30258559                40.06 ns/op            0 B/op          0 allocs/op
```

```go
goos: linux
goarch: amd64
pkg: csproto
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkProto3Marshal-12       23661498                48.01 ns/op           24 B/op          1 allocs/op
BenchmarkProto3Unmarshal-12     10480916               113.4 ns/op           104 B/op          2 allocs/op
```

```go
goos: linux
goarch: amd64
pkg: fastpb
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkProto3Marshal-12               15350882                74.96 ns/op           24 B/op          1 allocs/op
BenchmarkProto3Unmarshal-12              8406904               143.3 ns/op           112 B/op          3 allocs/op
BenchmarkProto3Marshal_Pool-12          20519343                58.51 ns/op            0 B/op          0 allocs/op
```

```go
goos: linux
goarch: amd64
pkg: msgp
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkProto3Marshal-12               20310513                57.10 ns/op           96 B/op          1 allocs/op
BenchmarkProto3Unmarshal-12             14050321                83.89 ns/op           16 B/op          2 allocs/op
BenchmarkProto3Marshal_Pool-12          28929116                41.00 ns/op            0 B/op          0 allocs/op
```

```go
goos: linux
goarch: amd64
pkg: gencode
cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
BenchmarkProto3Marshal-12               33004026                34.45 ns/op           24 B/op          1 allocs/op
BenchmarkProto3Unmarshal-12             31065225                36.94 ns/op           16 B/op          2 allocs/op
BenchmarkProto3Marshal_Pool-12          39586322                30.09 ns/op            0 B/op          0 allocs/op
```

