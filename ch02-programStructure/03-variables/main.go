package main

import (
	"fmt"
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
}

// even if it is declared after main, it is initialized before main
// so it can be used there
var packageLevelVariable = 3
