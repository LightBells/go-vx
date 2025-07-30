//go:build amd64 && !avx512
// +build amd64,!avx512

package vx

/*
#cgo CFLAGS: -std=c11 -O3 -march=skylake
#cgo LDFLAGS: -lm

#include <math.h>
#include <immintrin.h>

void vx_add(const size_t size, const float *x, const float *y, float *z) {
    __m256 *vx = (__m256 *)x;
    __m256 *vy = (__m256 *)y;
    __m256 *vz = (__m256 *)z;

    const size_t l = size / 8;
    for (size_t i = 0; i < l; ++i) {
        vz[i] = _mm256_add_ps(vx[i], vy[i]);
    }
}

void vx_sub(const size_t size, const float *x, const float *y, float *z) {
    __m256 *vx = (__m256 *)x;
    __m256 *vy = (__m256 *)y;
    __m256 *vz = (__m256 *)z;

    const size_t l = size / 8;
    for (size_t i = 0; i < l; ++i) {
        vz[i] = _mm256_sub_ps(vx[i], vy[i]);
    }
}

void vx_mul(const size_t size, const float *x, const float *y, float *z) {
    __m256 *vx = (__m256 *)x;
    __m256 *vy = (__m256 *)y;
    __m256 *vz = (__m256 *)z;

    const size_t l = size / 8;
    for (size_t i = 0; i < l; ++i) {
        vz[i] = _mm256_mul_ps(vx[i], vy[i]);
    }
}

void vx_div(const size_t size, const float *x, const float *y, float *z) {
    __m256 *vx = (__m256 *)x;
    __m256 *vy = (__m256 *)y;
    __m256 *vz = (__m256 *)z;

    const size_t l = size / 8;
    for (size_t i = 0; i < l; ++i) {
        vz[i] = _mm256_div_ps(vx[i], vy[i]);
    }
}

float vx_dot(const size_t size, const float *x, const float *y) {
    __m256 vsum = _mm256_setzero_ps();
    __m256 *vx = (__m256 *)x;
    __m256 *vy = (__m256 *)y;

    const size_t l = size / 8;
    for (size_t i = 0; i < l; ++i) {
        vsum = _mm256_fmadd_ps(vx[i], vy[i], vsum);
    }

    __attribute__((aligned(32))) float v[8];
    _mm256_store_ps(v, vsum);
    return v[0] + v[1] + v[2] + v[3] + v[4] + v[5] + v[6] + v[7];
}

void vx_normalize(const size_t size, const float *x, float *z) {
    __m256 *vx = (__m256 *)x;
    __m256 *vz = (__m256 *)z;

    const size_t l = size / 8;

    float sum_sq = vx_dot(size, x, x);

    float inv_norm = 1.0f / sqrtf(sum_sq);

    __m256 inv_vec = _mm256_set1_ps(inv_norm);
    for (size_t i = 0; i < l; ++i) {
        vz[i] = _mm256_mul_ps(vx[i], inv_vec);
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
	return 8
}
