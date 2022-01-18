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

	const freezingF = 32.0                                  // function level constant
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

// function declaration
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
