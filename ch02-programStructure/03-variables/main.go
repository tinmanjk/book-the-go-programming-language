package main

import (
	"fmt"
	"image/gif"
	"math/rand"
)

func main() {
	// var name type = expression

	var s string   // zero value -> always initialized
	fmt.Println(s) // ""

	// multiple variable declaration
	var i, j, k int // int, int, int
	fmt.Println(i, j, k)
	// multiple variable declaration and initiation
	var b, f, st = true, 2.3, "four" // bool, float64, string
	fmt.Println(b, f, st)

	fmt.Println(packageLevelVariable)

	// 2.3.1
	fmt.Println("2.3.1 Short Variable Declarations")
	shortVariableDeclaration()
	fmt.Println("------")

	// 2.3.2
	fmt.Println("2.3.2 Pointers")
	pointers()
	fmt.Println("------")

	// 2.3.3
	fmt.Println("2.3.3 The new Function")
	newFunc()
	fmt.Println("------")

}

// even if it is declared after main, it is initialized before main
// so it can be used there
var packageLevelVariable = 3

func shortVariableDeclaration() {

	// name := expression

	anim := gif.GIF{LoopCount: 30}
	freq := rand.Float64() * 3.0
	t := 0.0

	fmt.Println(anim, freq, t)

	i := 100                  // preferred an int
	var boiling float64 = 100 // a float64 -> due to expression literal having different default type
	fmt.Println(i, boiling)

	j := 90 // declaration
	// assignment
	i, j = j, i // swap values of i and j
	fmt.Println(i, j)

	k := 5
	// declaration + assignment if multiple variables used
	k, m := i, j // m is declared and assigned, k is just assigned a new value

	// k, m := i, j // no new variables error
	fmt.Println(k, m)
}

func pointers() {
	x := 1
	p := &x         // address-of operator -> *int
	fmt.Println(*p) // "1" // dereferencing operator *
	*p = 2          // equivalent to x = 2 // "dereferencing assignment?"
	fmt.Println(x)  // "2"

	{
		// pointers are comparable
		var x, y int
		var z *int = nil
		fmt.Println(&x == &y, &x == z, z == nil) // "false false true"
		z = &x
		fmt.Println(&x == z) // true
	}

	fmt.Println(packagePointer) //pointer variable
	//lint:ignore SA4000 <necessary reason after code>
	fmt.Println(f() == f()) // "false" -> distinct value -> different pointers effectively

	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
}

var packagePointer = f()

func f() *int {
	v := 1
	return &v // out of scope: escapes to the heap
}

func incr(p *int) int {
	*p++ // increments what p points to; does not   change p
	return *p
}

func newFunc() {
	p := new(int)   // p, of type *int, points to **unnamed** int variable
	fmt.Println(*p) // "0" -> memory at address IS zero-initialized
	*p = 2          // sets the unnamed int to 2
	fmt.Println(*p) // "2"

	{
		// new address
		p := new(int)
		q := new(int)
		fmt.Println(p == q) // "false"
	}
}

//lint:ignore U1000 ...
func newInt() *int {
	return new(int)
}

//lint:ignore U1000 ...
func newIntWithoutNew() *int {
	var dummy int
	return &dummy
}

//lint:ignore U1000 ...
var global *int

//lint:ignore U1000 ...
func fAssignsToGlobal() {
	//lint:ignore S1021 ...
	var x int
	x = 1
	global = &x // x needs to escape to the heap where the global can still reference it
}

//lint:ignore U1000 ...
func g() {
	y := new(int)
	*y = 1 // no need for y when g is finished -> it shouldn't be allocated on the heap
}
