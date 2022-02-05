package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

func main() {
	fmt.Println("Composing Types By Struct Embedding")
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	// **method promotion** of embedded Point methods to ColoredPoint
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
	// Point is **NOT a base class** for ColoredPoint
	// p.Distance(q) // compile error: cannot use q (ColoredPoint) as Point

	fmt.Println("\nEmbedded Pointer Type - Same Access to methods")
	{
		p := ColoredPointWithPointer{&Point{1, 1}, red}
		q := ColoredPointWithPointer{&Point{5, 4}, blue}
		// access to Distance method of the value type
		fmt.Println(p.Distance(*q.Point)) // "5" -> seems that *(q.Point) dereferencing the selected field TODO?
		q.Point = p.Point                 // p and q now share the same Point
		p.ScaleBy(2)
		fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
	}

	fmt.Println("\nMore than one embedded field")

	fmt.Println("\nUnnamed struct types - still access to methods via embedded fields")

}

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point // embedded type
	Color color.RGBA
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Distance(q Point) float64 {
	// has no access to the fields of a ColorPoint type
	// that it might be embedded into
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// what the compiler does for accessing **directly** the methods of embedded types ??? not tested
// func (cp ColoredPoint) Distance(q Point) float64 {
// 	return cp.Point.Distance(q)
// }
// func (cp *ColoredPoint) ScaleBy(factor float64) {
// 	cp.Point.ScaleBy(factor)
// }

type ColoredPointWithPointer struct {
	*Point
	Color color.RGBA
}

// more than one anonymous field - embed
// selector resolving such as p.ScaleBy by compiler
// 1. Directly declared methods
// 2. Promoted by embedded fields - once promoted
// 3. Promoted by embedded fields of embedded fields - twice promoted
// 4. ...
// ambiguous if methods of the same rank
type ColoredPointTwoEmbedded struct {
	Point
	color.RGBA
}

// unnamed struct types -> STILL HAVE METHODS
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock() // instead of mu.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
