package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("Example: Bit Vectory Type")
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(&x)         // "{1 9 42 144}" // implicit call String() from fmt package
	fmt.Println(x.String()) // "{1 9 42 144}" // compiler inserts implicit &x due to selection syntax
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	// -> **no implicit call from value to pointer** for String
	// alternatively make String a method for the value/not pointer -> breaking consistency
}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	// index of slice represents factor of 64,
	// the bit position in value represents teh remainder %64
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64) // word = index in slice
	return word < len(s.words) &&
		s.words[word]&(1<<bit) != 0 // simple bitmask - 0 if & fails at position bit
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	// if x = 131
	// word = 2
	// bit = 3
	// we extend the slice to length word + 1
	// so we can have effectively indeces 0 1 2
	// then the value at those indeces are zero initially
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// now we have [0 0 0] and we need to populate
	// the value there at the bit location -> i.e. 3
	s.words[word] |= 1 << bit
	// we have [0 0 8] or [0 0 0b1000] or 8
	// if there was already an element there the OR operation does not affect it
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword // simple OR-ing
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j) // index*64 + remainder
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
