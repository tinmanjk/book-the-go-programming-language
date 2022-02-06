package main

import (
	"fmt"
)

func main() {
	fmt.Println("Interface Types")

	// Good sources on io and interfaces
	// https://developpaper.com/on-the-confusion-of-golang-io-reading-and-writing/
	// https://dev.to/andyhaskell/what-is-the-io-firstTwentyBytes-in-go-48co

	// **Mine**
	var firstTwentyBytes firstTwentyBytesReadWriter = make([]byte, 25) // larger data, only first twenty bytes though
	for i := range firstTwentyBytes {
		firstTwentyBytes[i] = byte(i) + 1
	}
	fmt.Println(firstTwentyBytes) // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25]

	readOp := func(r ReadWriter, toBeReadInto []byte) {
		r.Read(toBeReadInto)
	}

	writeOp := func(r ReadWriter, toBeWritten []byte) {
		r.Write(toBeWritten)
	}

	input := []byte{} // 0 length, nil slice
	readOp(firstTwentyBytes, input)
	fmt.Println(input) // []

	input = make([]byte, 15)
	readOp(firstTwentyBytes, input)
	fmt.Println(input) // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]

	input = make([]byte, 30)
	readOp(firstTwentyBytes, input)
	fmt.Println(input) // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 0 0 0 0 0 0 0 0 0 0]

	input = make([]byte, 30)
	writeOp(firstTwentyBytes, input)
	fmt.Println(firstTwentyBytes) // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 21 22 23 24 25]
	// /** Mine
}

// **Single Method Interfaces** from package io

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// Combinations of existing interfaces.
// resembles struct embedding -> shorthand for writing all of its methods
type ReadWriter interface {
	// embedding an interface
	Reader // Read(p []byte) (n int, err error)
	// possibly a MIXTURE of the two styles
	Write(p []byte) (n int, err error) // Writer
}

// ** Mine
type firstTwentyBytesReadWriter []byte

// Read reads up to len(p) bytes into p.
// **does NOT use append** -> use of copy built-in see strings.Reader implementation
func (tb firstTwentyBytesReadWriter) Read(p []byte) (n int, err error) {
	n = copy(p, tb[:20])
	return n, nil
}

func (tb firstTwentyBytesReadWriter) Write(p []byte) (n int, err error) {
	n = copy(tb[:20], p)
	return n, nil
}

// ** /Mine
