package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	fmt.Println(reflect.ValueOf(2))
	fmt.Println(reflect.ValueOf(x))
	fmt.Println(reflect.ValueOf(&x))
	fmt.Println(reflect.ValueOf(&x).Elem())
}
