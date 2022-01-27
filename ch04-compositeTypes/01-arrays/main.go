package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var a [3]int             // array of 3 integers -> zero initialized
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// array literal
	var q [3]int = [3]int{1, 2, 3}
	fmt.Println(q)
	var r [3]int = [3]int{1, 2} // 3rd is also initiazlied to 0
	fmt.Println(r[2])           // "0"

	e := [...]int{1, 2, 3} // array ellipsis -> number of literal is the number of elements
	fmt.Printf("%T\n", e)  // "[3]int"

	// Size is part of the type
	{
		q := [3]int{1, 2, 3}
		// q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
		fmt.Println(q)
	}

	// literal syntax with indeces
	symbol := [...]string{EUR: "€", USD: "$", GBP: "£", RMB: "¥"} // order doesn't matter
	fmt.Println(RMB, symbol[RMB])                                 // "3 ¥"

	{
		r := [...]int{99: -1}
		fmt.Println(len(r), r[99]) //100, -1
	}

	// **array comparison**
	fmt.Println("\nArray Comparison")
	{
		a := [2]int{1, 2}
		b := [...]int{1, 2}
		c := [2]int{1, 3}
		fmt.Println(a == b, a == c, b == c) // "true false false"
		d := [3]int{1, 2}
		fmt.Println(d)
		// fmt.Println(a == d) // compile error: cannotcompare [2]int == [3]int
	}

	c1 := sha256.Sum256([]byte("x")) // [32]byte
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("% x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1) // %x with byte array seems to print without [] the values of the the bytes in hex
	// Output:
	//2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881 -> 64 length - 32byte * 2hex per byte
	//4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

	// for the call compiler uses the xmm0 floatingpoint register to actually COPY AND PASTE from one place of the stack
	// to another convenient!
	// main.go:69      0x4aa5ba        0f10842400020000                movups xmm0, xmmword ptr [rsp+0x200] -> the c1 in register
	// main.go:69      0x4aa5c2        0f11442410                      movups xmmword ptr [rsp+0x10], xmm0 -> from register to stack
	// main.go:69      0x4aa5c7        e8f4010000                      call $main.ones
	ones(c1)                       // pass by value -> all of the valus of the array on the stack/register
	fmt.Println("After one: ", c1) // stays the same

	// with explicit pass by refernce we just copy the address of the array into rax (1.17+register based)
	// and call main.zero
	// main.go:77      0x4aa6ab        488d8424f0010000                lea rax, ptr [rsp+0x1f0]
	// main.go:77      0x4aa6b3        e868020000                      call $main.zero
	zero(&c1)                       // explicitly by reference
	fmt.Println("After zero: ", c1) // 0
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func ones(arr [32]byte) {
	for i := range arr {
		arr[i] = 1
	}

	fmt.Println("Inside one: ", arr)
}

// explicitly pass by pointer - reference
func zero(ptr *[32]byte) {
	*ptr = [32]byte{} // use literal initialiation -> sets all to 0
}
