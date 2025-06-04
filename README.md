# go-vx

SIMD (Single Instruction Multiple Data) extension for golang.
Provide AVX, AVX512 binding for amd64, and provide NEON, SVE2 binding for arm64.

Forked from [https://github.com/monochromegane/go-avx](https://github.com/monochromegane/go-avx)

## Golang code example

```go
package main

import (
	"fmt"

	"github.com/gumigumi4f/go-vx"
)

func main() {
	dim := 8
	x := vx.AlignedAlloc(dim)
	y := vx.AlignedAlloc(dim)
	z := vx.AlignedAlloc(dim)
	defer vx.Free(x)
	defer vx.Free(y)
	defer vx.Free(z)

	for i := 0; i < dim; i++ {
		x[i] = float32(i)
		y[i] = float32(i + 1)
	}

	vx.Add(dim, x, y, z)

	fmt.Printf("%v\n", z) // [1 3 5 7 9 11 13 15]
}
```

## How to build

```sh
# for amd64 (avx2)
GOARCH=amd64 go build ./vx

# for amd64 (avx512)
GOARCH=amd64 go build -tags=avx512 ./vx

# for armv8 (NEON)
GOARCH=arm64 go build ./vx

# for armv9 (SVE2)
GOARCH=arm64 go build -tags=sve2 ./vx
```

## Features

- Add
- Sub
- Mul
- Div
- Dot

See also `vx_test.go`.

## Benchmark

Run `go test -bench Benchmark ./... -run="^Benchmark -benchmem`

### AMD64, AVX2 (Tested on GCP c4d-standard-2 instance)
```
goos: linux
goarch: amd64
pkg: github.com/gumigumi4f/go-vx
cpu: AMD EPYC 9B45
BenchmarkDotVx-2       	     817	   1408282 ns/op	     481 B/op	      20 allocs/op
BenchmarkDotNative-2   	     303	   3935093 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/gumigumi4f/go-vx	3.562s
```

### AMD64, AVX512 (Tested on GCP c4d-standard-2 instance)
```
goos: linux
goarch: amd64
pkg: github.com/gumigumi4f/go-vx
cpu: AMD EPYC 9B45
BenchmarkDotVx-2       	    1046	   1018718 ns/op	     375 B/op	      15 allocs/op
BenchmarkDotNative-2   	     304	   3947970 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/gumigumi4f/go-vx	3.446s
```

### ARM64, NEON (Tested on GCP c4a-standard-1 instance)
```
goos: linux
goarch: arm64
pkg: github.com/gumigumi4f/go-vx
BenchmarkDotVx     	     495	   2384785 ns/op	     794 B/op	      33 allocs/op
BenchmarkDotNative 	     220	   5423197 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/gumigumi4f/go-vx	4.100s
```

### ARM64, SVE2 (Tested on GCP c4a-standard-1 instance)
```
goos: linux
goarch: arm64
pkg: github.com/gumigumi4f/go-vx
BenchmarkDotVx     	     475	   2493596 ns/op	     827 B/op	      34 allocs/op
BenchmarkDotNative 	     219	   5430850 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/gumigumi4f/go-vx	4.116s
```

On GCP’s c4a instances, SVE2 can only use a 128-bit vector length—the same as NEON—so the added overhead makes it run slower than NEON.

## License

[MIT](https://github.com/gumigumi4f/go-vx/blob/master/LICENSE)
