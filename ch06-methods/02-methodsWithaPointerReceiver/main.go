package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Methods with a Pointer Receiver")
	p := Point{3, 3}
	p.ScaleByNoPointer(5)
	fmt.Println(p) // {3,3}

	pp := &p
	pp.ScaleBy(2)
	fmt.Println(p) // {6,6}

	(&p).ScaleBy(2)
	fmt.Println(p) // {12,12}

	p.ScaleBy(2)   // shorthand notation -> implicit &p by compiler
	fmt.Println(p) // {24,24}

	// Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal
	(*pp).ScaleByNoPointer(2) // pointer receiver with dereferencing works with non-pointer methods
	pp.ScaleByNoPointer(2)    // -> shorthand - where the comipler implicitly *pp

	// Summary: valid method call expression - 3 cases
	// 1. Receiver argument's type = Receiver paramter's type
	pp.ScaleBy(5)
	p.ScaleByNoPointer(3)
	// 2. Receiver argument's type T and receiver parameter *T -> implicit &T
	p.ScaleBy(2)
	// 3. Receiver argument's type *T and receiver parameter is T -> implicit *T
	pp.ScaleByNoPointer(3)

	fmt.Println("\n6.2.1 Nil Is a Valid Receiver Value")
	ilist := IntList{Value: 3, Tail: nil}
	fmt.Println(ilist.Sum())

	// Values maps a string key to a list of values.
	// type Values map[string][]string -> from url package

	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1" (first value)
	fmt.Println(m["item"])     // "[1 2]" (direct map access)

	m = nil
	url.Values(nil).Get("item") // works -> will return ""
	fmt.Println(m.Get("item"))  // "" -> still works
	// m.Add("item", "3")         // panic: assignment to entry in nil map

}

type Point struct {
	X, Y float64
}

// cannot be named the same as the pointer receiver method
func (p Point) ScaleByNoPointer(factor float64) {
	//lint:ignore SA4005 ineffective assignment -> it will not be assigned to
	p.X *= factor
	//lint:ignore SA4005 ineffective assignment -> it will not be assigned to
	p.Y *= factor
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum() // would call the method with nil receiver
}
