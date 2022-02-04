package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
)

func main() {

	// Mistakes found by the runtime -> index out of bound e.g.
	drawCard := func() string {
		return "Ace of Spades"
	}

	suit := func(card string) string {
		return "Spades" // change to "NonSpades" for panic
	}

	switch s := suit(drawCard()); s {
	case "Spades": // ...
	case "Hearts": // ...
	case "Diamonds": // ...
	case "Clubs": // ...
	default:
		panic(fmt.Sprintf("invalid suit %q", s)) //s Joker?
	}

	fmt.Println("\nUnnecessary panics")
	nilPanic := func(f *fieldAccess) {
		if f == nil {
			panic("x is nil") // unnecessary!
		}
		fmt.Println(f.f1) // will panic here
		// panic: runtime error: invalid memory address or nil pointer dereference
	}
	nilPanic(&fieldAccess{}) // change to nil for panic

	// incorrect input, misconfiguration, or failing I/O -> ERRORS
	// rest is reserved for panic

	fmt.Println("\nMust*** wrapper pattern that converts errors into panics")
	_, _ = regexp.Compile("example") // regexp and error
	// not checking errors if a call **CANNOT FAIL**
	// here reason is that regular expressions are LITERALS known at compile time
	// Must*** wrapper function -> like template.Must
	_ = regexp.MustCompile("example") // returns just the regexp
	// useful for **package-level variable initialization**

	fmt.Println("\nPanic and FIFO order of defer (reverse than normal stack)")

	defer printStack() // manually output the stack -> maybe for diagnostic purposes
	// deferred function calls are made **BEFORE ANY STACK UNWINDING** -> preserving the state

	f(3, true) // modify second argument to true to cause panic
	// defer functions across the stack -> from topmost function first called
	// then panic message and stack trace to std.Err + os.Exit(2)

	// 	f(3)
	// 	f(2)
	// 	f(1)
	// 	defer 0
	// 	defer 1
	// 	defer 2
	// 	defer 3
	// panic: runtime error: integer divide by zero // -> panic message
	// goroutine 1 [running]: // -> stack trace
	// main.f(0x0, 0x1)
	// main.f(0x1, 0x1)
	// main.f(0x2, 0x1)
	// main.f(0x3, 0x1)
	// main.main()
}

//lint:ignore U1000 ...
var httpSchemeRE = regexp.MustCompile(`^https?:`)

// "http:" or "https:"

type fieldAccess struct {
	f1 int
}

func f(x int, causePanic bool) {
	if !causePanic {
		return
	}
	defer fmt.Printf("defer %d\n", x) // moved before panic (text version)
	fmt.Printf("f(%d)\n", x+0/x)      // panics if x == 0 -> triggers defer before exiting
	f(x-1, causePanic)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
