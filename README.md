# go-vx

SIMD (Single Instruction Multiple Data) extension for golang.
Provide AVX (Advanced Vector Extensions) binding for amd64 and NEON binding for arm64.

Forked from [https://github.com/monochromegane/go-avx](https://github.com/monochromegane/go-avx)

## Golang code example

Set `CGO_CFLAGS_ALLOW=-mfma` to build binary for amd64.

```go
package main

import (
	"fmt"

	"github.com/gumigumi4f/go-vx"
)

func main() {
	dim := 8
	x := vx.Malloc(dim)
	y := vx.Malloc(dim)
	z := vx.Malloc(dim)
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

## Features

- Add
- Sub
- Mul
- Div
- Dot

See also `vx_test.go`.

## Benchmark

Run `go test -bench Benchmark ./... -run="^Benchmark"`

### AMD64 (Tested on AWS EC2 c5.large instance)
```
goos: linux
goarch: amd64
pkg: github.com/gumigumi4f/go-vx
BenchmarkDotVx-2             364           3201061 ns/op
BenchmarkDotNative-2         100          10833781 ns/op
PASS
ok      github.com/gumigumi4f/go-vx     3.953s
```

### ARM64 (Tested on AWS EC2 c6g.medium instance)
```
goos: linux
goarch: arm64
pkg: github.com/gumigumi4f/go-vx
BenchmarkDotVx               274           4275091 ns/op
BenchmarkDotNative           100          10599838 ns/op
PASS
ok      github.com/gumigumi4f/go-vx     4.154s
```

## License

[MIT](https://github.com/gumigumi4f/go-vx/blob/master/LICENSE)
