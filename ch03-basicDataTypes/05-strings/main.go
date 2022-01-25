package main

import "fmt"

func main() {

	fmt.Println("Strings")

	s := "hello, world"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
	// c := s[len(s)]          // panic: index out of range

	// substring operation
	// s[i:j] [i,j) -> j-i bytes result
	fmt.Println(s[0:5]) // "hello"
	// default values if omitted -> i=0, j=len(s)
	fmt.Println(s[:5]) // "hello"
	fmt.Println(s[7:]) // "world"
	fmt.Println(s[:])  // "hello, world"

	// the + operator -> new string
	fmt.Println("goodbye" + s[5:]) // "goodbye, world"

	// comparison -> byte by byte - lexicographic ordering

	// Immutability
	// assembly of string concatenation with + operator -> runtime function
	// new byte array
	{
		s := "left foot"
		// lea rbx, ptr [rip+0x19e0e] -> "left foot" address
		// mov qword ptr [rsp+0xb8], rbx -> local variable for beginning of byte array
		// mov qword ptr [rsp+0xc0], 0x9 -> local variable for length of byte array - 9
		t := s
		s += ", right foot"
		// lea rdi, ptr [rip+0x1a5b0] -> ", right foot" byte array adress
		// mov esi, 0xc -> len
		// call $runtime.concatstring2
		// results are returned in rax - address of NEW byte array
		// and rbx -> length of new byte array

		fmt.Println(s)
		fmt.Println(t)
	}

	// not allowed to modify strings in place
	// s[0] = 'L' // compile error: cannot assign to s[0]

	// substrings SHARE the same data because it's immutable
	// copies of a string is cheap -> just the pointer to the beginning and len are copied

	// 3.5.1 String Literals
	fmt.Println("\n3.5.1 String Literals: ")
	stringLiterals()

	// 3.5.2 String Literals
	fmt.Println("\n3.5.2 Unicode: ")
	unicode()

}

func stringLiterals() {
	// escape sequences:

	// \a “alert” or bell
	// \b backspace
	// \f form feed
	// \n newline
	// \r carriage return
	// \t tab
	// \v vertical tab
	// \' single quote (only in the rune literal '\'' )
	// \" double quote (only within "..." literals)
	// \\ backslash

	// \xhh -> hexadecimal escape - EXACTLY TWO!!! hex -> 2*4 = one BYTE
	s := "\"Hell\x6F\"" // Hello
	fmt.Println(s)

	// RAW String Literal backquotes `...` -> no escape sequences
	rawStringLiteral := `"Hello \ no escaping \x6F
	hello from new line`
	fmt.Println(rawStringLiteral)

	// carriage returns are deleted and processed as newline only -> just LF, no CR
	newLineExperiment := `new
line`
	fmt.Printf("%X\n", newLineExperiment[3]) // here the new line is simply 'A' line feed -> LF no CR
}

func unicode() {
}
