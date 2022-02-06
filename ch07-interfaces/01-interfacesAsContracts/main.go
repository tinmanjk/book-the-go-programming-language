package main

import (
	"fmt"
)

func main() {
	fmt.Println("Interfaces as Contracts")
	// example about Printf/Sprintf/Fprintf -> **no longer the case** in stdlib
	// internal implementation no longer uses bytes.Buffer
	// but simpler implementation 'type buffer []byte'

	fmt.Println("\nio.Writer single method interface -> Write(p []byte) (n int, err error)")
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	// io.Writer interface
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

	fmt.Println("\nfmt.Stringer interface -> String() string -> see 7.10")
}

type ByteCounter int

// io.Writer -> Write(p []byte) (n int, err error)
func (bc *ByteCounter) Write(b []byte) (n int, err error) {
	*bc += ByteCounter(len(b))
	return len(b), nil
}
