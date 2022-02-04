package main

import (
	"fmt"
	"math"
	"time"

	"github.com/tinmanjk/tgpl/ch06-methods/01-methodDeclarations/otherPackage"
)

func main() {
	fmt.Println("Method Declarations")
	// need of methods:
	// "methods to express the properties and operations of each data
	// structure so that clients need not access the object’s representation
	// directly."

	const day = 24 * time.Hour // time.Duration type -> type Duration int64 named type for int64
	fmt.Println(day.Seconds()) // "86400"

	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	// receiver argument (p.) before the method name -> parallels declaration
	fmt.Println(p.Distance(q)) // "5", method call -> **selector**
	// methods and fields use the same namespace -> cannot have method X on Point

	// mine - maybe another way to print the type of the method? TODO
	fmt.Printf("%T\n", Distance)   // func(main.Point, main.Point) float64
	fmt.Printf("%T\n", p.Distance) // func(main.Point) float64

	fmt.Println("\nReceiver constraints - same package, no pointer or interface type")
	renamedTypeValue := renamedType{Name: "Value"}
	renamedTypeValue.extensionWithRenamedType()
	// see below package-level declarations for more examples

	fmt.Println("\nSame method name different receiver")
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"
}

// compilation error - invalid receiver - should be **FROM THIS PACKAGE**
// func (o otherPackage.ExportedType) extensionOnAnotherPackage() {

// }

// if we provide a named type for the imported type however -> no compilation error
type renamedType otherPackage.ExportedType

func (r renamedType) extensionWithRenamedType() {
	fmt.Println(r.Name)
}

// named type's underlying type should NOT be **a pointer** or **an interface**
//lint:ignore U1000 ....
type renamedTypePointer *otherPackage.ExportedType

// compilation error: invalid receiver - pointer or interface
// func (r renamedTypePointer) extensionWithRenamedTypePointer() {
// 	fmt.Println(r.Name)
// }

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
// p Point is the method "receiver" parameter ->
// legacy from old OOP -> calling a method = sending a message to an object
// no special names such as **this** or **self** -> convention first name of type name
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// A Path is a journey connecting the points with straight lines.
// named slice type <-> not a struct type
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}

	return sum
}
