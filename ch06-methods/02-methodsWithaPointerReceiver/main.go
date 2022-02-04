package main

import "fmt"

func main() {
	fmt.Println("Methods with a Pointer Receiver")
	p := Point{3, 3}
	p.ScaleByNoPointer(5)
	fmt.Println(p) // {3,3}

	pp := &p
	pp.ScaleBy(2)
	fmt.Println(p) // {6,6}

	(&p).ScaleBy(2)
	fmt.Println(p) // {12,12}

	p.ScaleBy(2)   // shorthand notation -> implicit &p by compiler
	fmt.Println(p) // {24,24}

	// Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal
	(*pp).ScaleByNoPointer(2) // pointer receiver with dereferencing works with non-pointer methods
	pp.ScaleByNoPointer(2)    // -> shorthand - where the comipler implicitly *pp

	// Summary: valid method call expression - 3 cases
	// 1. Receiver argument's type = Receiver paramter's type
	pp.ScaleBy(5)
	p.ScaleByNoPointer(3)
	// 2. Receiver argument's type T and receiver parameter *T -> implicit &T
	p.ScaleBy(2)
	// 3. Receiver argument's type *T and receiver parameter is T -> implicit *T
	pp.ScaleByNoPointer(3)

}

type Point struct {
	X, Y float64
}

// cannot be named the same as the pointer receiver method
func (p Point) ScaleByNoPointer(factor float64) {
	//lint:ignore SA4005 ineffective assignment -> it will not be assigned to
	p.X *= factor
	//lint:ignore SA4005 ineffective assignment -> it will not be assigned to
	p.Y *= factor
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
