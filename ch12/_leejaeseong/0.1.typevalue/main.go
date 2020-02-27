package main

import (
	"fmt"
	"reflect"
)

func main() {
	//t := reflect.TypeOf(3)
	//fmt.Println(t.String())
	//fmt.Println(t)
	//
	//var w io.Writer = os.Stdout
	//fmt.Println(reflect.TypeOf(w))
	//
	//fmt.Printf("%T\n", 3)
	//fmt.Printf("%T\n", os.Stdout)

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	v = reflect.ValueOf("string")
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
}
