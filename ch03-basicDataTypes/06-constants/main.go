package main

import (
	"fmt"
	"time"
)

// untyped package level constants
const (
	e  = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
)

func main() {

	// pi = 3.15 // compiler error
	fmt.Println(e, pi)
	// constant in type
	var p [IPv4Len]byte
	fmt.Println(p)

	// const declaration inside a function
	const noDelay time.Duration = 0 // explicit typed declaration
	// time.Duration is named type with underlying int64
	const timeout = 5 * time.Minute // inferred but still typed
	// -> %T verb for the type of a const/variable
	fmt.Printf("%T %[1]v\n", noDelay)     //	"time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     //	"time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) //	"time.Duration 1m0s"

	const (
		a = 1
		b // b = 1
		c = 2
		d // d = 2
	)

	fmt.Println(a, b, c, d) // "1 1 2 2"

}

const IPv4Len = 4
