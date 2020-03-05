package pak2

import "fmt"

type Dst struct {
	a int
	B int
	c int16
}

func New() *Dst {
	return &Dst{
		a: 960312,
		B: 123456,
		c: 6033,
	}
}

func (d* Dst) String() string {
	return fmt.Sprintf("a : %d, B : %d, c : %d", d.a, d.B, d.c)
}
