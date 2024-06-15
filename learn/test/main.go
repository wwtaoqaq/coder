package main

import (
	"fmt"
	"unsafe"
)

func main() {
	n := int64(10)
	pointer := unsafe.Pointer(&n)
	addr := uintptr(pointer)
	fmt.Printf("pointer:%v, addr:%v", pointer, addr)
}

func read(p *int) int {
	v := *p
	return v
}

func swap(p *int64) *int32 {
	v := (*int32)(unsafe.Pointer(p))
	return v
}
