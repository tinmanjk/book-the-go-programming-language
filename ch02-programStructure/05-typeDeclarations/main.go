package main

import (
	"fmt"

	"github.com/tinmanjk/tgpl/ch02-programStructure/05-typeDeclarations/tempconv"
)

func main() {
	// 2.5 Type Declarations
	fmt.Println("2.5 Type Declarations")
	// !!! assembly if values int for clearer instructions
	var cel tempconv.Celsius = 100 // mov qword ptr [rsp+0x20], 0x64

	// conversion operation T(x) is NOT a function call
	// it does NOT change the value of representation -> see assembly
	fahr := tempconv.Fahrenheit(cel) // mov qword ptr [rsp+0x18], 0x64
	// Conversion allowed if same underlying type OR unnamed pointers
	// point to variables of the same underlying type
	fmt.Println(fahr) // 100
	// !!! Conversion between numeric types/ strings MAY CHANGE
	// the representation -> see next chapters
	// CONVERSION NEVER FAILS AT RUNTIME!

	// actual change of value with custom logic
	fahr = tempconv.CToF(cel)
	fmt.Println(fahr) // 212

	var c tempconv.Celsius
	var f tempconv.Fahrenheit
	// Arithmetic operations
	fmt.Printf("%g\n", c-tempconv.Celsius(f)) // "0" subtraction

	// Comparisons
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	// fmt.Println(c == f)          // compile error:	type mismatch
	fmt.Println(c == tempconv.Celsius(f)) // "true"!

	// Method calls for a named type
	c = tempconv.FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need tocall String explicitly !!!
	fmt.Printf("%s\n", c)   // "100°C" -> no need here too
	fmt.Println(c)          // "100°C" -> no need because of the String method apaprently
	fmt.Printf("%g\n", c)   // "100"; does not callString
	fmt.Println(float64(c)) // "100"; does not call String

	fmt.Println("------")

}
