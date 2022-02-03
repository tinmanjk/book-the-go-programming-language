package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Vadiatic Functions")
	fmt.Println(sum())           // "0"
	fmt.Println(sum(3))          // "3"
	fmt.Println(sum(1, 2, 3, 4)) // "10"
	// implicitly we allocate an an array
	// saves the arguments in the array
	// makes a slice with len and cap = len of array
	// passing the slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10" -> ... elipsis after value

	fmt.Println("\nDifferent Type - signature although effectively the same")
	fmt.Printf("%T\n", f) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name) // "Line 	12: undefined: count"
}

// ... elipsis for the last parameter of the function
func sum(vals ...int) int {
	total := 0
	// vals is a []int
	for _, val := range vals {
		total += val
	}
	return total
}

func f(...int) {}
func g([]int)  {}

// typical signature (with f naming suffix)
// interface{} -> any value
func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
