package main

import "fmt"

// package level declaration
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9

	// %g -> general -> gives you the shorter of
	// %e (scientific) or %f (fixed point)
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
