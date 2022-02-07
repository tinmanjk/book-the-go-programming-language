package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Println("Interface Satisfaction")

	var w io.Writer
	w = os.Stdout // OK: *os.File has Write method
	// **assignability to interface** determined by method set -> no compiler * and & magic here
	// *T contains T, but T does not contain *T
	// https://stackoverflow.com/questions/33587227/method-sets-pointer-vs-value-receiver
	// If you have an interface I, and some or all of the methods in I's method set are
	// provided by methods with a receiver of *T (with the remainder being provided by methods with a receiver of T),
	// then *T satisfies the interface I, but T doesn't.
	w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
	// w = time.Second       // compile error:time.Duration lacks Write method

	//lint:ignore S1021 ...
	var rwc io.ReadWriteCloser
	rwc = os.Stdout // OK: *os.File has Read, Write, Close methods
	// rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
	// interface-to-interface

	w = rwc // OK: io.ReadWriteCloser has Write method
	// rwc = w // compile error: io.Writer lacks Close method

	w.Write([]byte("\nHello to std-out\n")) // use of variable for compilation

	// var _ = IntSet{}.String() // compile error: String requires *IntSet receiver
	// Mine:
	fmt.Println(IntSet{})  // {} -> cannot find the String method on the value
	fmt.Println(&IntSet{}) // composite literal exception - addressable -> works
	// https://stackoverflow.com/questions/40793289/go-struct-literals-why-is-this-one-addressable
	// composite literal, taking its address will create an **anonymous variable under the hood**,
	// and the result of the expression will be the address of this anonymous variable.

	// Assignability to interfaces
	var s IntSet
	var _ = s.String()      // OK: s is a variable and &s has a String method
	var _ fmt.Stringer = &s // OK
	// var _ fmt.Stringer = s  // compile error: IntSet lacks String method

	fmt.Println("\nThe empty interface type -> interface{}")
	// https://research.swtch.com/interfaces -> for more information about internal representation
	// of interface values -> maybe look at Section 7.5 Interface values
	var any interface{}
	any = true
	fmt.Println(any)
	any = 12.34
	// any += 3 // invalid operation
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)
	any = map[string]int{"one": 1}
	fmt.Println(any)
	any = new(bytes.Buffer)
	fmt.Println(any)

	fmt.Println("\nDocumenting explicitly satisfaction of interface")
	// Compile time assertion - so we can catch here if there is no satisfaction
	// ? where should this be placed
	// *bytes.Buffer must satisfy io.Writer
	var _ io.Writer = (*bytes.Buffer)(nil)

	fmt.Println("\nAbstraction of Contrete types as \"Property set\" of methods from interface")

}

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

// magazines, books
type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP3", "WAV"
}

// TV Shows, Videos
type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP4", "WMV"
	Resolution() (x, y int)
}

// non obvious intersection of properties - audio and video
type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

type IntSet struct { /* ... */
}

func (*IntSet) String() string {
	return "Not implemented"
}
