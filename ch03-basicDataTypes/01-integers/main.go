package main

import "fmt"

func main() {
	// % only on integers -> sign is always the same as the dividend
	fmt.Println(-5 % 3)  // -2
	fmt.Println(-5 % -3) // -2
	// / division depends on the value
	fmt.Println(5 / 4)     // 1 -> truncates for integers
	fmt.Println(5.0 / 4.0) // 1.25

	// overflow -> high-order bits are discarded
	var u uint8 = 255   // FF + 1 = 100 -> 0
	fmt.Println(u, u+1) // "255 0"

	var i int8 = 127    // 7F + 1  = 80 -> -1
	fmt.Println(i, i+1) // "127 -128

	// bit-wise operations
	// x << n | x may be signed or unsigned, n should be unsigned
	var x uint8 = 1<<1 | 1<<5 // 2 | 32
	var y uint8 = 1<<1 | 1<<2 // 2 | 4

	// %b is the verb for binary
	// 08 modified %b -> the adverb for formatting -> pad result with 0s for exactly 8
	fmt.Printf("%08b\n", x) // "00100010", the set{1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set{1, 2}

	fmt.Printf("%08b\n", x&y) // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y) // "00100110", the union {1, 2, 5}"

	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference{5} (excluding intersection with y)
	// novo
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set{2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set{0, 4}

	fmt.Println("\nRight shift and signs")

	var unsignedInt uint8 = 32
	fmt.Printf("%v -> %08b\n", unsignedInt, unsignedInt) // "00100000"
	unsignedInt >>= 1
	fmt.Printf("%v -> %08b\n", unsignedInt, unsignedInt) // "00010000"

	// right shift and signed numbers -> bit sign is copied (0 or 1)

	// printing int8 in %08b prints the sign and the positive representation, i.e. -1 = -000 0001, not 1111 1111
	// **so conversion to uint** is needed to see tha actual bytes
	// this conversion does NOT alter the bits - reuses the value
	{
		// https://stackoverflow.com/questions/37582550/twos-complement-and-fmt-printf
		var negativeInt int8 = -1           // FF
		convertedUint := uint8(negativeInt) // just reuses the bits -> no conversion -> see disassembly
		fmt.Println(convertedUint)          // 255
	}
	var negativeInt int8 = -32
	fmt.Printf("%v -> %08b\n", negativeInt, uint8(negativeInt)) // "11100000"
	negativeInt >>= 1
	fmt.Printf("%v -> %08b\n", negativeInt, uint8(negativeInt)) // "11110000"

	fmt.Println("\nUse of unsigned numbers as indeces of array to prevent overflow")
	medals := []string{"gold", "silver", "bronze"}
	// len returns int -> if it did uint then i would be i and 0-1 would not be -1 but max int
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}

	fmt.Println("\nConversions")
	var apples int32 = 1
	var oranges int16 = 2
	// var compote int = apples + oranges // compile error
	var compote = int(apples) + int(oranges) // common type

	fmt.Println(apples, oranges, compote)

	// out of range
	f := 1e100        // a float64
	implInt := int(f) // result is implementation-dependent 32 or 64bit trunc
	// either always minvalue though for some reason...not FFFF but 8000
	fmt.Printf("%f -> %d -> %x\n", f, implInt, implInt)

	{
		// Octal Hex Etc
		fmt.Println("\nOctal vs Hex")
		octalNumber := 0666
		// [1] adverb -> use the first operand -> apparently 1based
		// # adverb -> for octal and hex -> prefix
		fmt.Printf("%d %[1]o %#[1]o\n", octalNumber) // "438 666 0666"
		x := int64(0xdeadbeef)
		fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
		// Output:
		// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
	}

	{
		fmt.Println("\nRunes")
		ascii := 'a'
		unicode := '国'
		newline := '\n'
		fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a' "
		fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
		fmt.Printf("%d %[1]c %[1]q\n", newline) // "10 ..newline gets printed.. '\n'"
	}

}
