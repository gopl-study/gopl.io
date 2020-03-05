package cycle

import (
	"reflect"
)

type visit struct {
	ptr uintptr
	typ reflect.Type
}

func checkCycle(v reflect.Value, seen map[visit]bool) (ret bool) {
	ret = false
	var getPtr func() uintptr = nil

	if v.CanAddr() {
		getPtr = v.UnsafeAddr
	} else if v.Kind() == reflect.Map {
		getPtr = v.Pointer
	}

	if getPtr != nil {
		vt := visit{
			ptr: getPtr(),
			typ: v.Type(),
		}

		if seen[vt] {
			return true
		}

		seen[vt] = true
		defer func() {
			if !ret {
				delete(seen, vt)
			}
		}()
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return checkCycle(v.Elem(), seen)
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if checkCycle(v.Index(i), seen) {
				return true
			}
		}
	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if checkCycle(v.Field(i), seen) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			if checkCycle(v.MapIndex(k), seen) {
				return true
			}
		}
	}


	return
}

func IsCycle(v interface{}) bool {
	seen := make(map[visit]bool)
	return checkCycle(reflect.ValueOf(v), seen)
}
