package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Function Declarations")

	// Func Declaration form
	// func name(parameter-list) (result-list) {
	// 	body
	// }

	fmt.Println(hypot(3, 4)) // "5"

	// here the Type is the function signature
	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"

}

// func SignatureOnly(x float64) float64
// missing function body compilation error
// https://stackoverflow.com/questions/44763440/source-missing-function-body-how-is-it-compiling
// it has to be defined elsewhere in assembly

// like func archLog(x float64) float64
// and:
// func Log(x float64) float64
// TEXT ·archLog(SB),NOSPLIT,$0
// in log_amd64.s assembly file in the same package

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// four ways to declare a function with same signature
// func(int,int) int
func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }
