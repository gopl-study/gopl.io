package pak1

import "fmt"

type Src struct {
	a int
	B int
	c int16
}

func New() *Src {
	return &Src{
		a: 312960,
		B: 654321,
		c: 3306,
	}
}

func (s* Src) String() string {
	return fmt.Sprintf("a : %d, B : %d, c : %d", s.a, s.B, s.c)
}