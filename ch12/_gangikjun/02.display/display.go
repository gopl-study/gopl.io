package display

import (
	"fmt"
	"reflect"
)

// Display Display
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display()

