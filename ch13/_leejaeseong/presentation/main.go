package main

import (
	_ "../equal"
	"reflect"
	"runtime"
	"unsafe"
)

type MemoryMark struct {
	then *runtime.MemStats
	now  *runtime.MemStats
}

func MemoryLeaks(f func()) int {
	var then = new(runtime.MemStats)
	var now = new(runtime.MemStats)

	runtime.GC()
	runtime.ReadMemStats(then)

	f()

	runtime.GC()
	runtime.ReadMemStats(now)

	return int((now.Mallocs - then.Mallocs) - (now.Frees - then.Frees))
}

func MarkMemory() *MemoryMark {
	m := &MemoryMark{
		then: new(runtime.MemStats),
		now:  new(runtime.MemStats),
	}

	runtime.GC()
	runtime.ReadMemStats(m.then)

	return m
}

func (m *MemoryMark) Release() int {
	runtime.GC()
	runtime.ReadMemStats(m.now)

	return int((m.now.Mallocs - m.then.Mallocs) - (m.now.Frees - m.then.Frees))
}

func main() {

	l := MemoryLeaks(func() {
		var pArr [10000]uintptr
		for i := range pArr {
			pArr[i] = uintptr(unsafe.Pointer(new(int)))
		}
		for i := 1; i < 10000; i++ {
			m := MarkMemory()
			for j := range pArr {
				*(*int)(unsafe.Pointer(pArr[j])) = i
			}
			println(m.Release())
		}
	})

	println(l)

	type link struct {
		val string
		tail *link
	}

	a, b, c := &link{val: "a"}, &link{val: "b"}, &link{val: "c"}
	println(reflect.DeepEqual(a, a))
	println(reflect.DeepEqual(b, b))
	println(reflect.DeepEqual(c, c))
	println(reflect.DeepEqual(a, b))
	println(reflect.DeepEqual(a, c))
}
