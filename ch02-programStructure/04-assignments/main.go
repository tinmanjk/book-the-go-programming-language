package main

import (
	"fmt"
	"os"
)

func main() {
	// 2.4.1. Tuple Assignment

	fmt.Println("2.4.1. Tuple Assignment")
	fmt.Println(gcd(3, 4))
	// discard operator if unwanted
	_, err := os.Open("foo.txt") // function call returns two values
	fmt.Println(err)

	// three operators that return two values, one a bool
	// v, ok = m[key] // map lookup -> _,ok = m[key]
	// v, ok = x.(T)  // type assertion
	// v, ok = <-ch   // channel receive

	fmt.Println("------")

	// 2.4.2. Assignability
	fmt.Println("2.4.2. Assignability")
	// implicit assignments
	// function call -> arguments are assigned to parameters
	// return statement -> return operands assigned to result variables
	// composite type -> literal expression
	medals := []string{"gold", "silver", "bronze"}
	// equal to medals[0] = "gold", etc.
	fmt.Println(medals)

	// simple rules: exact match of types, nil only to interface/pointer types
	// constants: -> 3.6. more flexible rules
	// more difficult rules to follow

	// comparability with == and != is related to assignability
	fmt.Println("------")

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
