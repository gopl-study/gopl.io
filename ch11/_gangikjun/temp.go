package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "string"
	for i := range str {
		fmt.Println(str[i], reflect.ValueOf(str[i]).Type())
	}
}
