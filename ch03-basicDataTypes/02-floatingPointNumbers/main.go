package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Floating point numbers")
	fmt.Println("Float 32 is not that big for integers")
	var f float32 = 16_777_216 // 1 << 24 = 2^24 -> 16 million max
	fmt.Println(f == f+1)      // "true"!

	// Exkurs:
	// cvttsd2si is the instruction to convert from float64 to int64 -> could be slow
	floatNumber := float64(2.9)
	intNumber := int64(floatNumber)
	fmt.Println(intNumber)

	for x := 0; x < 8; x++ {
		// f -> no exponent
		// 8 character -> 3 for decimal
		// rounding takes place
		fmt.Printf("x = %d ex = %8.2f\n", x, math.Exp(float64(x)))
	}

	fmt.Println("\nNaN and special IEEE 754 values")
	var z float64 = 0                  // added 0 for clarity, could be omitted
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	// NaN IEEE-754
	nan := math.NaN() // float64
	//lint:ignore SA4012 ...
	fmt.Println(nan == nan, nan < nan, nan > nan, nan != nan) //	"false false false true" - last one is negation of false
	fmt.Println(math.IsNaN(nan))                              // true -> use this and not the "=="

}
