package vx

/*
#cgo CFLAGS: -std=c11 -O3
#cgo LDFLAGS: -lm

#include <stdlib.h>
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func AlignedAlloc2D(dim int, size int) [][]float32 {
	alloc_size := size*dim
	alloc_size_ := alloc_size
	size = align(alloc_size)
	ptr := C.aligned_alloc((C.size_t)(C.sizeof_float * vectorLength()), (C.size_t)(C.sizeof_float*alloc_size))
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ptr)),
		Len:  alloc_size,
		Cap:  alloc_size,
	}
	goSliceToFill := *(*[]float32)(unsafe.Pointer(&hdr))
	if alloc_size_ != alloc_size {
		for i := alloc_size_; i < size; i++ {
			goSliceToFill[i] = 0.0
		}
	}

	goSlice := *(*[][]float32)(unsafe.Pointer(&hdr))
	return goSlice
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
