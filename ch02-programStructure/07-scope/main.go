package main

import (
	"fmt"
	"log"
	"os"
)

//lint:ignore U1000 ...
func f() {}

var g = "g"

func main() {
	// brackets block
	{
		// shadowing package level f
		f := "f"
		fmt.Println(f) // "f"; local var f shadowspackage-level func f
		fmt.Println(g) // "g"; package-level var
		// fmt.Println(h) // compile error: undefined: h
	}

	{
		x := "hello!"
		for i := 0; i < len(x); i++ {
			x := x[i]
			if x != '!' {
				x := x + 'A' - 'a'
				fmt.Printf("%c", x) // "HELLO" (oneletter per iteration)
			}
		}
	}

	{
		x := "hello"
		for _, x := range x { // implicit block x
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter  per iteration)
		}
	}

	{
		f := func() int { return 42 }
		g := func(input int) int { return 42 }

		// implicit scope in if declaration
		if x := f(); x == 0 {
			fmt.Println(x)
		} else if y := g(x); x == y { // if implicit seen in else and elseif scopes!
			fmt.Println(x, y)
		} else {
			fmt.Println(x, y)
		}
		// fmt.Println(x, y) // compile error: x and y are	   not visible here
	}

}

//lint:ignore U1000 ...
var cwd string

func init() {
	// shadowing with short variable declaration
	// use var err error instead
	var err error
	// short variable declaration only assigns if the variable
	// is first declared in the **SAME** lexical block -> here the function
	// cwd is package-level, so it is declared anew! and err is assigned to!!!
	// -> subtle shadowing
	cwd, err := os.Getwd() // assignment to err, cwd is stil declared!
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}

	log.Printf("Working directory = %s", cwd)
}
