package methods

import (
	"fmt"
	"reflect"
	"strings"
)

// Print 는 값 x의 메소드들을 출력한다.
func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s \n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
		strings.TrimPrefix(methType.String(), "func"))
	}
}