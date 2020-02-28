package main

import (
	"strconv"
)

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case stringer: return x.String()
	case string: return x
	case int: return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}

type noStringerType struct {
	Id int
	Name string
}

type StringerType struct {
	Id int
	Name string
}

func (s StringerType) String() string {
	return strconv.Itoa(s.Id) + " / " + s.Name
}

func main() {
	println(Sprint("string")) // string
	println(Sprint(123456)) // 123456
	println(Sprint(true)) // true
	println(Sprint(false)) // false
	println(Sprint(StringerType{
		Id:   135,
		Name: "Stringer Type",
	})) // 135 / Stringer Type

	println(Sprint(123.0)) // ???
	println(Sprint(int64(1234))) // ???
	println(Sprint(int32(12))) // ???
	println(Sprint(noStringerType{
		Id:   321,
		Name: "No Stringer Type",
	})) // ???
}
