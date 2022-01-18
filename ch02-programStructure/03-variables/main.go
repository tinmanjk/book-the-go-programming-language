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
	shortVariableDeclaration()

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
