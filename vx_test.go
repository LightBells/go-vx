package vx

import (
	"math"
	"math/rand"
	"testing"
)

func TestMalloc(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := AlignedAlloc(size)
			defer Free(x)

			if len(x) != align(size) {
				t.Errorf("AlignedAlloc should return float slice size of %d, but size is %d", align(size), len(x))
			}
		}(size)
	}
}

func TestAdd(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := AlignedAlloc(size)
			y := AlignedAlloc(size)
			z := AlignedAlloc(size)
			defer Free(x)
			defer Free(y)
			defer Free(z)

			truth := make([]float32, size)
			for i := 0; i < size; i++ {
				x[i] = float32(i)
				y[i] = float32(i + 1)
				truth[i] = x[i] + y[i]
			}

			Add(size, x, y, z)

			for i := 0; i < size; i++ {
				if truth[i] != z[i] {
					t.Errorf("Add should return %f in %d, but %f", truth[i], i, z[i])
				}
			}
		}(size)
	}
}

func TestSub(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := AlignedAlloc(size)
			y := AlignedAlloc(size)
			z := AlignedAlloc(size)
			defer Free(x)
			defer Free(y)
			defer Free(z)

			truth := make([]float32, size)
			for i := 0; i < size; i++ {
				x[i] = float32(i)
				y[i] = float32(i + 1)
				truth[i] = x[i] - y[i]
			}

			Sub(size, x, y, z)

			for i := 0; i < size; i++ {
				if truth[i] != z[i] {
					t.Errorf("Mul should return %f in %d, but %f", truth[i], i, z[i])
				}
			}
		}(size)
	}
}

func TestMul(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := AlignedAlloc(size)
			y := AlignedAlloc(size)
			z := AlignedAlloc(size)
			defer Free(x)
			defer Free(y)
			defer Free(z)

			truth := make([]float32, size)
			for i := 0; i < size; i++ {
				x[i] = float32(i)
				y[i] = float32(i + 1)
				truth[i] = x[i] * y[i]
			}

			Mul(size, x, y, z)

			for i := 0; i < size; i++ {
				if truth[i] != z[i] {
					t.Errorf("Mul should return %f in %d, but %f", truth[i], i, z[i])
				}
			}
		}(size)
	}
}

func TestDiv(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := AlignedAlloc(size)
			y := AlignedAlloc(size)
			z := AlignedAlloc(size)
			defer Free(x)
			defer Free(y)
			defer Free(z)

			truth := make([]float32, size)
			for i := 0; i < size; i++ {
				x[i] = float32(i)
				y[i] = float32(i + 1)
				truth[i] = x[i] / y[i]
			}

			Div(size, x, y, z)

			for i := 0; i < size; i++ {
				if truth[i] != z[i] {
					t.Errorf("Div should return %f in %d, but %f", truth[i], i, z[i])
				}
			}
		}(size)
	}
}

func TestDot(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := AlignedAlloc(size)
			y := AlignedAlloc(size)
			defer Free(x)
			defer Free(y)

			var truth float32
			for i := 0; i < size; i++ {
				x[i] = float32(i)
				y[i] = float32(i + 1)
				truth += x[i] * y[i]
			}

			result := Dot(size, x, y)
			if truth != result {
				t.Errorf("Dot should return %f, but %f", truth, result)
			}
		}(size)
	}
}

func TestNormalize(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := AlignedAlloc(size)
			z := AlignedAlloc(size)
			defer Free(x)
			defer Free(z)

			truth := make([]float32, size)
			sum := float32(0.0)
			for i := 0; i < size; i++ {
				x[i] = float32(i)
				truth[i] = x[i]
				sum += truth[i] * truth[i]
			}
			invNorm := 1 / float32(math.Sqrt(float64(sum)))
			for i := 0; i < size; i++ {
				truth[i] = truth[i] * invNorm
			}

			Normalize(size, x, z)

			for i := 0; i < size; i++ {
				if truth[i] != z[i] {
					t.Errorf("Normalize should return %f in %d, but %f", truth[i], i, z[i])
				}
			}
		}(size)
	}
}

func BenchmarkDotVx(b *testing.B) {
	num := 16384
	size := 512

	vx := AlignedAlloc(size)
	for j := 0; j < size; j++ {
		vx[j] = rand.Float32()
	}

	vys := make([][]float32, num)
	for i := range vys {
		vys[i] = AlignedAlloc(size)
		for j := 0; j < size; j++ {
			vys[i][j] = rand.Float32()
		}
	}

	similarities := make([]float32, num)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j, vy := range vys {
			similarities[j] = Dot(size, vx, vy)
		}
	}

	Free(vx)
	for i := range vys {
		Free(vys[i])
	}
}

func BenchmarkDotNative(b *testing.B) {
	num := 16384
	size := 512

	vx := make([]float32, size)
	for j := 0; j < size; j++ {
		vx[j] = rand.Float32()
	}

	vys := make([][]float32, num)
	for i := range vys {
		vys[i] = make([]float32, size)
		for j := 0; j < size; j++ {
			vys[i][j] = rand.Float32()
		}
	}

	similarities := make([]float32, num)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j, vy := range vys {
			similarity := float32(0)
			for k := 0; k < size; k++ {
				similarity += vx[k] * vy[k]
			}
			similarities[j] = similarity
		}
	}
}
