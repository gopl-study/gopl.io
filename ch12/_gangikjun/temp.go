package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(3)
	x := v.Interface()
	i := x.(int)
	fmt.Printf("%d\n", i)
}
