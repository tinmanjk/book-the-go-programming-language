package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tinmanjk/tgpl/ch06-methods/06-encapsulation/counter"
)

func main() {
	fmt.Println("Encapsulation")
	// unit of encapsulation = package -> the mechanism of import/export
	// NOT the type as in other languages
	cThisPackage := Counter{n: 3}
	fmt.Println(cThisPackage.N()) // 3
	cThisPackage.n = 5
	fmt.Println(cThisPackage.N()) // 5

	// cOtherPackage := counter.Counter{n: 3} // compilaton error unknown field
	cOther := counter.Counter{}
	cOther.Increment()
	fmt.Println(cOther.N()) // 1

	fmt.Println("\nOmit Get prefix from getters, but still have Set for setters")
	fmt.Println(log.Flags()) // No GetFlags
	log.SetFlags(3)
	fmt.Println(log.Flags())

	fmt.Println("\nEncapsulation NOT always good - sometimes expose")
	const day = 24 * time.Hour // time.Duration reveals that it is "type Duration int64" -> number nanoseconds
	// allows for **arithmetic and comparison** operations inherent for int and allows for constants
	fmt.Println(day.Seconds()) // "86400"

	// another difference is a named type -> for a slice e.g. allows clients
	// to use the slice literal, for-range etc which are NOT allowed on the IntSet
	// so -> Path - sequence of points with no fields to be added, so it's okay
	// IntSet However can change its internal representation!

}

type IntSet []uint64 // would allow client to access the slice directly

type IntSetEncapsulated struct {
	//lint:ignore U1000 ...
	words []uint64
}

type Counter struct{ n int }

func (c *Counter) N() int     { return c.n }
func (c *Counter) Increment() { c.n++ }
func (c *Counter) Reset()     { c.n = 0 }
