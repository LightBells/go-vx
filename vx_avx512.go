//go:build amd64 && avx512
// +build amd64,avx512

package vx

/*
#cgo CFLAGS: -std=c11 -O3 -march=skylake-avx512
#cgo LDFLAGS: -lm

#include <math.h>
#include <immintrin.h>

void vx_add(const size_t size, const float *x, const float *y, float *z) {
    __m512 *vx = (__m512 *)x;
    __m512 *vy = (__m512 *)y;
    __m512 *vz = (__m512 *)z;

    const size_t l = size / 16;
    for (size_t i = 0; i < l; ++i) {
        vz[i] = _mm512_add_ps(vx[i], vy[i]);
    }
}

void vx_sub(const size_t size, const float *x, const float *y, float *z) {
    __m512 *vx = (__m512 *)x;
    __m512 *vy = (__m512 *)y;
    __m512 *vz = (__m512 *)z;

    const size_t l = size / 16;
    for (size_t i = 0; i < l; ++i) {
        vz[i] = _mm512_sub_ps(vx[i], vy[i]);
    }
}

void vx_mul(const size_t size, const float *x, const float *y, float *z) {
    __m512 *vx = (__m512 *)x;
    __m512 *vy = (__m512 *)y;
    __m512 *vz = (__m512 *)z;

    const size_t l = size / 16;
    for (size_t i = 0; i < l; ++i) {
        vz[i] = _mm512_mul_ps(vx[i], vy[i]);
    }
}

void vx_div(const size_t size, const float *x, const float *y, float *z) {
    __m512 *vx = (__m512 *)x;
    __m512 *vy = (__m512 *)y;
    __m512 *vz = (__m512 *)z;

    const size_t l = size / 16;
    for (size_t i = 0; i < l; ++i) {
        vz[i] = _mm512_div_ps(vx[i], vy[i]);
    }
}

float vx_dot(const size_t size, const float *x, const float *y) {
    __m512 vsum = _mm512_setzero_ps();
    __m512 *vx = (__m512 *)x;
    __m512 *vy = (__m512 *)y;

    const size_t l = size / 16;
    for (size_t i = 0; i < l; ++i) {
        vsum = _mm512_fmadd_ps(vx[i], vy[i], vsum);
    }

    return _mm512_reduce_add_ps(vsum);
}

void vx_normalize(const size_t size, const float *x, float *z) {
    __m512 *vx = (__m512 *)x;
    __m512 *vz = (__m512 *)z;

    const size_t l = size / 16;

    float sum_sq = vx_dot(size, x, x);

    float inv_norm = 1.0f / sqrtf(sum_sq);

    __m512 inv_vec = _mm512_set1_ps(inv_norm);
    for (size_t i = 0; i < l; ++i) {
        vz[i] = _mm512_mul_ps(vx[i], inv_vec);
    }
}
*/
import "C"

func Add(size int, x, y, z []float32) {
	size = align(size)
	C.vx_add((C.size_t)(size), (*C.float)(&x[0]), (*C.float)(&y[0]), (*C.float)(&z[0]))
}

func Sub(size int, x, y, z []float32) {
	size = align(size)
	C.vx_sub((C.size_t)(size), (*C.float)(&x[0]), (*C.float)(&y[0]), (*C.float)(&z[0]))
}

func Mul(size int, x, y, z []float32) {
	size = align(size)
	C.vx_mul((C.size_t)(size), (*C.float)(&x[0]), (*C.float)(&y[0]), (*C.float)(&z[0]))
}

func Div(size int, x, y, z []float32) {
	size = align(size)
	C.vx_div((C.size_t)(size), (*C.float)(&x[0]), (*C.float)(&y[0]), (*C.float)(&z[0]))
}

func Dot(size int, x, y []float32) float32 {
	size = align(size)
	dot := C.vx_dot((C.size_t)(size), (*C.float)(&x[0]), (*C.float)(&y[0]))
	return float32(dot)
}

func Normalize(size int, x, z []float32) {
	size = align(size)
	C.vx_normalize((C.size_t)(size), (*C.float)(&x[0]), (*C.float)(&z[0]))
}

func vectorLength() int {
	return 16
}
