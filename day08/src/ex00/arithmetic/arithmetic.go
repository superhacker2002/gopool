package arithmetic

import (
	"errors"
	"unsafe"
)

var (
	ErrEmptySlice    = errors.New("slice is empty")
	ErrNegativeIndex = errors.New("index is negative")
	ErrOutOfBounds   = errors.New("index is out of bounds")
)

func GetElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, ErrEmptySlice
	}
	if idx < 0 {
		return 0, ErrNegativeIndex
	}
	if idx >= len(arr) {
		return 0, ErrOutOfBounds
	}

	start := unsafe.Pointer(&arr[0])
	size := unsafe.Sizeof(int(0))
	return *(*int)(unsafe.Pointer(uintptr(start) + size*uintptr(idx))), nil
}
