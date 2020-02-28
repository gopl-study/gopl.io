package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a struct {
		bool
		float64
		int16
	}
	var b struct {
		float64
		int16
		bool
	}
	var c struct {
		bool
		int16
		float64
	}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))

	var x struct {
		a bool
		b int16
		c []int
	}
	fmt.Println(unsafe.Sizeof(x))
}
