package main

import (
	"fmt"
	"image/color"
	"math"
)

func main() {
	// var cp coloredpoint.ColoredPoint
	// cp.X = 1
	// fmt.Println(cp.Point.X)
	// cp.Point.Y = 2
	// fmt.Println(cp.Y)

	// red := color.RGBA{255, 0, 0, 255}
	// blue := color.RGBA{0, 0, 255, 255}

	// p := ColoredPoint{Point{1, 1}, red}
	// q := ColoredPoint{Point{5, 4}, blue}
	// fmt.Println(p.Distance(q.Point))
	// p.ScaleBy(2)
	// q.ScaleBy(2)
	// fmt.Println(p.Distance(q.Point))

	// p := ColoredPoint{&Point{1, 1}, red}
	// q := ColoredPoint{&Point{5, 4}, blue}
	// fmt.Println(p.Distance(*q.Point)) //	"5"
	// q.Point = p.Point                 // p와 q는 같은 Point를 공유한다
	// p.ScaleBy(2)
	// fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"

	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance   // 메소드 값
	fmt.Println(distanceFromP(q)) // "5"
	var origin Point              // {0, 0}
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy
	scaleP(2)
	fmt.Println(p)
	scaleP(3)
	fmt.Println(p)
	scaleP(10)
	fmt.Println(p)
}

// Point Point
type Point struct{ X, Y float64 }

// ColoredPoint ColoredPoint
type ColoredPoint struct {
	*Point
	Color color.RGBA
}

// Distance distance
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy ScaleBy
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Path 는 선분에서 점을 잇는 경로다
type Path []Point

// Distance distance
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
