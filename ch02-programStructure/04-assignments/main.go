package main

import (
	"fmt"
	"os"
)

func main() {
	// 2.4.1. Tuple Assignment
	fmt.Println(gcd(3, 4))

	// discard operator if unwanted
	_, err := os.Open("foo.txt") // function call returns two values
	fmt.Println(err)

	// three operators that return two values, one a bool
	// v, ok = m[key] // map lookup -> _,ok = m[key]
	// v, ok = x.(T)  // type assertion
	// v, ok = <-ch   // channel receive
}

// Greatest Common Divisor
func gcd(x, y int) int {
	for y != 0 {
		// x = y, y = x%y
		// right is evaluated before assignment to left
		// so no need to save expressions in temp variables
		x, y = y, x%y
	}
	return x
}

//lint:ignore U1000 ...
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		// easier to write, but maybe harder to read
		x, y = y, x+y
	}
	return x
}
