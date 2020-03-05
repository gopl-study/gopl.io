// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 357.

// Package unsafeptr demonstrates basic use of unsafe.Pointer.
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//!+main
	var x struct {
		a bool
		b int16
		c []int
	}


	// unsafe.Sizeof()
	// unsafe.Offsetof()
	// unsafe.Alignof()
	// equivalent to
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42

	fmt.Println(x.b) // "42"
	//!-main
}

/*
//!+wrong
	// NOTE: subtly incorrect!
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb := (*int16)(unsafe.Pointer(tmp))
	*pb = 42
//!-wrong
*/
