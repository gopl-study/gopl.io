package main

import (
	"./pak1"
	"./pak2"
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//return &Src{
	//	a: 312960,
	//	B: 654321,
	//	c: 3306,
	//}
	src := pak1.New()

	//return &Dst{
	//	a: 960312,
	//	B: 123456,
	//	c: 6033,
	//}
	dst := pak2.New()


	srcV := reflect.ValueOf(src).Elem()
	fmt.Print("src : ")
	fmt.Println(src)
	fmt.Println("============== dst ==============")

	dstV := reflect.ValueOf(dst).Elem()
	fmt.Println(dst)
	dstV.Field(1).SetInt(srcV.Field(1).Int())
	fmt.Println(dst)

	*(*int)(unsafe.Pointer(dstV.Field(0).UnsafeAddr())) = int(srcV.Field(0).Int())
	fmt.Println(dst)
	*(*int16)(unsafe.Pointer(dstV.Field(2).UnsafeAddr())) = int16(srcV.Field(2).Int())
	fmt.Println(dst)

	dstV.Field(2).SetInt(srcV.Field(2).Int()) // error
}