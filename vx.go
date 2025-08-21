package vx

/*
#cgo CFLAGS: -std=c11 -O3
#cgo LDFLAGS: -lm

#include <stdlib.h>

(void *) c_memcopy(void *dest, const void *src, size_t n) {
	return memcopy(dest, src, n)
}
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func Memcopy(dst, src []float32, n int) {
	return C.c_memcopy(dst, src, n * C.sizeof_float)
}

func AlignedAlloc2D(rows, cols int) [][]float32 {
	alignedCols := align(cols)

	totalSize := rows * alignedCols

	alignmentBytes := C.size_t(C.sizeof_float * vectorLength())
	totalBytes := C.size_t(C.sizeof_float * totalSize)
	ptr := C.aligned_alloc(alignmentBytes, totalBytes)
	if ptr == nil {
		panic("aligned_alloc failed to allocate memory")
	}

	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ptr)),
		Len:  totalSize,
		Cap:  totalSize,
	}
	data := *(*[]float32)(unsafe.Pointer(&hdr))

	if cols != alignedCols {
		for r := 0; r < rows; r++ {
			rowStart := r * alignedCols
			for c := cols; c < alignedCols; c++ {
				data[rowStart+c] = 0.0
			}
		}
	}

	matrix := make([][]float32, rows)

	for i := 0; i < rows; i++ {
		start := i * alignedCols
		end := start + cols
		capacityEnd := start + alignedCols

		matrix[i] = data[start:end:capacityEnd]
	}

	return matrix
}

func Free2D(matrix [][]float32) {
	if matrix == nil || len(matrix) == 0 {
		return
	}
	ptr := unsafe.Pointer(&matrix[0][0])
	C.free(ptr)
}

func AlignedAlloc(size int) []float32 {
	size_ := size
	size = align(size)
	ptr := C.aligned_alloc((C.size_t)(C.sizeof_float * vectorLength()), (C.size_t)(C.sizeof_float*size))
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ptr)),
		Len:  size,
		Cap:  size,
	}
	goSlice := *(*[]float32)(unsafe.Pointer(&hdr))
	if size_ != size {
		for i := size_; i < size; i++ {
			goSlice[i] = 0.0
		}
	}
	return goSlice
}

func Free(v []float32) {
	C.free(unsafe.Pointer(&v[0]))
}

func align(size int) int {
	v := vectorLength()

	rem := size % v
	if rem == 0 {
		return size
	}
	return size + (v - rem)
}
