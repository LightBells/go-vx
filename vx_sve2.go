//go:build arm64 && sve2
// +build arm64,sve2

package vx

/*
#cgo CFLAGS: -std=c11 -O3 -march=armv9-a
#cgo LDFLAGS: -lm

#include <arm_sve.h>

void vx_add(const size_t size, const float *x, const float *y, float *z) {
    const size_t vl = svcntw();
    svbool_t pg;

    for (size_t i = 0; i < size; i += vl) {
        pg = svwhilelt_b32(i, size);
        svfloat32_t vx = svld1(pg, &x[i]);
        svfloat32_t vy = svld1(pg, &y[i]);
        svfloat32_t vz = svadd_m(pg, vx, vy);
        svst1(pg, &z[i], vz);
    }
}

void vx_sub(const size_t size, const float *x, const float *y, float *z) {
    const size_t vl = svcntw();
    svbool_t pg;

    for (size_t i = 0; i < size; i += vl) {
        pg = svwhilelt_b32(i, size);
        svfloat32_t vx = svld1(pg, &x[i]);
        svfloat32_t vy = svld1(pg, &y[i]);
        svfloat32_t vz = svsub_m(pg, vx, vy);
        svst1(pg, &z[i], vz);
    }
}

void vx_mul(const size_t size, const float *x, const float *y, float *z) {
    const size_t vl = svcntw();
    svbool_t pg;

    for (size_t i = 0; i < size; i += vl) {
        pg = svwhilelt_b32(i, size);
        svfloat32_t vx = svld1(pg, &x[i]);
        svfloat32_t vy = svld1(pg, &y[i]);
        svfloat32_t vz = svmul_m(pg, vx, vy);
        svst1(pg, &z[i], vz);
    }
}

void vx_div(const size_t size, const float *x, const float *y, float *z) {
    const size_t vl = svcntw();
    svbool_t pg;

    for (size_t i = 0; i < size; i += vl) {
        pg = svwhilelt_b32(i, size);
        svfloat32_t vx = svld1(pg, &x[i]);
        svfloat32_t vy = svld1(pg, &y[i]);
        svfloat32_t vz = svdiv_m(pg, vx, vy);
        svst1(pg, &z[i], vz);
    }
}

float vx_dot(const size_t size, const float *x, const float *y) {
    const size_t vl = svcntw();
    svbool_t pg;
    svfloat32_t vsum = svdup_f32(0.0f);

    for (size_t i = 0; i < size; i += vl) {
        pg = svwhilelt_b32(i, size);
        svfloat32_t vx = svld1(pg, &x[i]);
        svfloat32_t vy = svld1(pg, &y[i]);
        vsum = svmla_m(pg, vsum, vx, vy);
    }

    return svaddv(svptrue_b32(), vsum);
}
*/
import "C"
import (
	"math"
	"sync"
)

var (
	vl int
	once sync.Once
)

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

func initVectorLength() {
	vl := int(C.svcntw())
	if vl == 0 {
		panic("failed to get vector length")
	}
}

func vectorLength() int {
	once.Do(initVectorLength)
	return vl
}
