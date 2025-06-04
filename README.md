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
GOARCH=amd64 GOAMD64=v2 go build ./vx

# for amd64 (avx512)
GOARCH=amd64 GOAMD64=v3 go build -tags=avx512 ./vx

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

Run `go test -bench Benchmark ./... -run="^Benchmark"`

### AMD64, AVX2 (Tested on AWS EC2 c5.large instance)
```
goos: linux
goarch: amd64
pkg: github.com/gumigumi4f/go-vx
BenchmarkDotVx-2             364           3201061 ns/op
BenchmarkDotNative-2         100          10833781 ns/op
PASS
ok      github.com/gumigumi4f/go-vx     3.953s
```

### AMD64, AVX512 (Tested on AWS EC2 c5.large instance)
WIP

### ARM64, NEON (Tested on GCP c4a-standard-1 instance)
```
goos: linux
goarch: arm64
pkg: github.com/gumigumi4f/go-vx
BenchmarkDotVx     	     498	   2379579 ns/op
BenchmarkDotNative 	     219	   5443635 ns/op
PASS
ok  	github.com/gumigumi4f/go-vx	4.111s
```

### ARM64, SVE2 (Tested on GCP c4a-standard-1 instance)
```
goos: linux
goarch: arm64
pkg: github.com/gumigumi4f/go-vx
BenchmarkDotVx     	     352	   3375619 ns/op
BenchmarkDotNative 	     219	   5440813 ns/op
PASS
ok  	github.com/gumigumi4f/go-vx	4.219s
```

On GCP’s c4a instances, SVE2 can only use a 128-bit vector length—the same as NEON—so the added overhead makes it run slower than NEON.

## License

[MIT](https://github.com/gumigumi4f/go-vx/blob/master/LICENSE)
