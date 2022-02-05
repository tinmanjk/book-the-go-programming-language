package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Method Value")
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance        // method value - binding a method to a specific value
	fmt.Println(distanceFromP(q))      // "5"
	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979",

	scaleP := p.ScaleBy // method value
	scaleP(2)
	fmt.Println(p) // p becomes (2, 4)
	scaleP(3)
	fmt.Println(p) // then (6, 12)
	scaleP(10)
	fmt.Println(p) // then (60, 120)

	// passing method value into package's API
	r := new(Rocket)
	// time.AfterFunc(2*time.Second, func() { r.Launch() })
	time.AfterFunc(1*time.Second, r.Launch)
	time.Sleep(2 * time.Second) //-> need to block for sync looking code / not exiting main
	// "ROCKET LAUNCHED"

	fmt.Println("\nMethod Expression - T.f or (*T).f ")
	{
		p := Point{1, 2}
		q := Point{4, 6}
		distance := Point.Distance // method expression -> no access from Point to Scaleby
		// effectively **making a static method** to pass the receiver as ordinary argument
		fmt.Println(distance(p, q))       // "5"
		fmt.Printf("%T\n", distance)      // "func(Point, Point) float64" -> receiver part of the signature
		distanceValue := p.Distance       // method value
		fmt.Printf("%T\n", distanceValue) // "func(main.Point) float64" -> receiver not part of the signature

		scale := (*Point).ScaleBy // (*Point) has access to Distance too!
		scale(&p, 2)
		fmt.Println(p)            // "{2 4}"
		fmt.Printf("%T\n", scale) // "func(*Point, float64)"

		// see Path.TranslateBy below -> useful to have a collection of methods to use on many receivers
	}

}

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}

type Rocket struct { /* ... */
}

func (r *Rocket) Launch() {
	fmt.Println("ROCKET LAUNCHED")
}
