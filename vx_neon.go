//go:build arm64 && !sve2
// +build arm64,!sve2

package vx

/*
#cgo CFLAGS: -std=c11 -O3 -march=armv8-a
#cgo LDFLAGS: -lm

#include <math.h>
#include <arm_neon.h>

void vx_add(const size_t size, const float *x, const float *y, float *z) {
    float32x4_t *vx = (float32x4_t *)x;
    float32x4_t *vy = (float32x4_t *)y;
    float32x4_t *vz = (float32x4_t *)z;

    const size_t l = size / 4;

    for (size_t i = 0; i < l; ++i) {
        vz[i] = vaddq_f32(vx[i], vy[i]);
    }
}

void vx_sub(const size_t size, const float *x, const float *y, float *z) {
    float32x4_t *vx = (float32x4_t *)x;
    float32x4_t *vy = (float32x4_t *)y;
    float32x4_t *vz = (float32x4_t *)z;

    const size_t l = size / 4;

    for (size_t i = 0; i < l; ++i) {
        vz[i] = vsubq_f32(vx[i], vy[i]);
    }
}

void vx_mul(const size_t size, const float *x, const float *y, float *z) {
    float32x4_t *vx = (float32x4_t *)x;
    float32x4_t *vy = (float32x4_t *)y;
    float32x4_t *vz = (float32x4_t *)z;

    const size_t l = size / 4;

    for (size_t i = 0; i < l; ++i) {
        vz[i] = vmulq_f32(vx[i], vy[i]);
    }
}

void vx_div(const size_t size, const float *x, const float *y, float *z) {
    float32x4_t *vx = (float32x4_t *)x;
    float32x4_t *vy = (float32x4_t *)y;
    float32x4_t *vz = (float32x4_t *)z;

    const size_t l = size / 4;

    for (size_t i = 0; i < l; ++i) {
        vz[i] = vdivq_f32(vx[i], vy[i]);
    }
}

float vx_dot(const size_t size, const float *x, const float *y) {
    float32x4_t vsum = vdupq_n_f32(0.0);
    float32x4_t *vx = (float32x4_t *)x;
    float32x4_t *vy = (float32x4_t *)y;

    const size_t l = size / 4;

    for (size_t i = 0; i < l; ++i) {
        vsum = vfmaq_f32(vsum, vx[i], vy[i]);
    }

    float32_t v[4];
    vst1q_f32(v, vsum);

    return v[0] + v[1] + v[2] + v[3];
}

void vx_normalize(const size_t size, const float *x, float *z) {
    float32x4_t *vx = (float32x4_t *)x;
    float32x4_t *vz = (float32x4_t *)z;

    const size_t l = size / 4;

    float sum_sq = vx_dot(size, x, x);
    if (sum_sq == 0.0f) {
        return;
    }

    float inv_norm = 1.0f / sqrtf(sum_sq);

    float32x4_t inv_vec = vdupq_n_f32(inv_norm);
    for (size_t i = 0; i < l; ++i) {
        vz[i] = vmulq_f32(vx[i], inv_vec);
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
	return 4
}
