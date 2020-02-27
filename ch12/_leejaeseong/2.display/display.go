package main // import "display"

import (
	"display/eval"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), 10)
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

//!+display
func display(path string, v reflect.Value, limitDepth int) {
	if limitDepth <= 0 {
		fmt.Printf("%s = %s\n", path, formatAtom(v))
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), limitDepth - 1)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), limitDepth - 1)
		}
	case reflect.Map:
		for i, key := range v.MapKeys() {
			display(fmt.Sprintf("%s.key[%d]", path, i), key, limitDepth - 1)
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key), limitDepth - 1)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), limitDepth - 1)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), limitDepth - 1)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func main() {
	e, _ := eval.Parse("sqrt(A / pi)")
	println("=============eval.Parse(\"sqrt(A / pi)\")=============")
	Display("e", e)

	type Movie struct {
		Title, Subtitle string
		Year int
		Color bool
		Actor map[string]string
		Oscars []string
		Sequel *string
	}

	strangelove := Movie{
		Title: "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year: 1964,
		Color: false,
		Actor: map[string]string{
			"Dr. Strangelove": "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley": "Peter Sellers",
			"Gen. Buck Turgidson": "George C. Scott",
			"Brig. Gen. Jack D. Ripper": "Sterling Hayden",
			`Maj. T.J. "King" Kong`: "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	println()
	println()
	println("=============type Movie struct {...=============")
	Display("strangelove", strangelove)

	println()
	println()
	println("=============os.Stderr=============")
	Display("os.Stderr", os.Stderr)

	println()
	println()
	println("=============reflect.ValueOf(os.Stderr)=============")
	Display("rV", reflect.ValueOf(os.Stderr))

	println()
	println()
	var i interface{} = 3
	Display("i", i)

	Display("&i", &i)

	type SomeType struct {
		Id int
		Name string
	}

	println()
	println()
	println("=============map[string]type SomeType struct {...=============")

	mapSomeType := make(map[SomeType]bool)
	mapSomeType[SomeType{1, "type1"}] = true
	mapSomeType[SomeType{2, "type2"}] = true
	Display("mapSomeType", mapSomeType)

	println()
	println()
	println("=============map[[2]string]type SomeType struct {...=============")
	mapSlice := make(map[[2]SomeType]bool)
	mapSlice[[2]SomeType {
		SomeType{11, "array1Type1"},
		SomeType{12, "array1Type2"},
	}] = true
	mapSlice[[2]SomeType {
		SomeType{21, "array2Type1"},
		SomeType{22, "array2Type2"},
	}] = true
	Display("mapSlice", mapSlice)

	println()
	println()
	println("=============type Cycle struct {...=============")
	type Cycle struct { Value int; Tail *Cycle }
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c)
}