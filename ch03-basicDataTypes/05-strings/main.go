package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

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

	// 3.5.3 Unicode-UTF8
	fmt.Println("\n3.5.3 UTF-8: ")
	utf8Examples()

	// 3.5.4 Strings and Byte Slices
	fmt.Println("\n3.5.4 Strings and Byte Slices")
	stringsByteSlices()
}

// 3.5.1
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

// 3.5.3
func utf8Examples() {

	// Variable encoding 1 to 4 bytes: xxxxxx are left for values

	// 1 byte: 7 bits available
	// 0xxxxxx runes 0–127 (ASCII)

	// 2 bytes: 11 bits available -> 2048 possible values
	// 110xxxxx 10xxxxxx -> [128–2047] -> (values <128 unused)

	// 3 bytes: 16 bits available for 65535 possible values
	// 1110xxxx 10xxxxx 10xxxxxx -> [2048–65535] -> (values <2048 unused)

	// 4 bytes: 21 bits available
	// 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx  65536–0x10ffff (other values unused)

	sameString := [4]string{}
	sameString[0] = "世界"                       // 2 3byte chars
	sameString[1] = "\xe4\xb8\x96\xe7\x95\x8c" // 6 hex byte escapes
	sameString[2] = "\u4e16\u754c"             // 2 16bit unicode escapes
	sameString[3] = "\U00004e16\U0000754c"     // 2 32bit Unicode escapes
	fmt.Println(sameString)

	// e4b896 -> 24 bits -> 3 byte ENCODED
	for i := 0; i < 3; i++ {
		fmt.Printf("%08b ", sameString[1][i])
		// 11100100 10111000 10010110
		// e4 b8 96
	}

	fmt.Println()
	// RUNE FORM -> No Encodings - Fixed size 4bytes
	sameRune := [3]rune{'世', '\u4e16', '\U00004e16'}
	fmt.Printf("%c - %d - %x\n", sameRune, sameRune, sameRune) // [世 世 世] - [19990 19990 19990] - [4e16 4e16 4e16]
	fmt.Printf("%032b", sameRune[0])                           // 0000000000000000 0100 1110 0001 0110 -> 4 e 1 6
	// 0100 1110 0001 0110 -> 16 bit 4e16 without encodings
	// original encoded -> e4 b8 96
	// [1110 -> 0100] e4 -> just the 4
	// [10 -> 1110 00] b8 -> e (and 00 carry for next)
	// [10 -> 01 0110] 96 -> 16 (with carry from previous)
	fmt.Println()

	runeOneHex := '\xFF' // valid hex escape -> ONLY for one byte though -> can be higher than ASCII!!!
	fmt.Printf("%c -%c \n", runeOneHex, '\u00ff')
	// more than one does not work
	// runeOneHex = '\xFF\xFF'// won't compile

	// iterating a string for code points -> it **decodes** the UTF-8 encodings
	helloEastAsia := "Hello, 世界"

	fmt.Println("Implicit UTF-8 Decoding for-range loop")
	for i, r := range helloEastAsia {
		fmt.Printf("%d\t%q\t%#x\n", i, r, r)
	}

	fmt.Println("No Decoding for loop") //
	for i := 0; i < len(helloEastAsia); i++ {
		if helloEastAsia[i] < 0x80 {
			fmt.Printf("%d\t%q\t%#x\n", i, helloEastAsia[i], helloEastAsia[i])
		} else {
			fmt.Printf("%d\t'x'\t%#x\n", i, helloEastAsia[i])
		}

	}
	fmt.Println()
	fmt.Println(len(helloEastAsia))                    // 13
	fmt.Println(utf8.RuneCountInString(helloEastAsia)) // 9

	// explicit decoding -> same as for range loop
	fmt.Println("\nExplicit Decoding for loop")
	for i := 0; i < len(helloEastAsia); {
		r, size := utf8.DecodeRuneInString(helloEastAsia[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	// invalid utf8- replacement character
	// 110xxxxx 10xxxxxx -> [128–2047] -> (values <128 unused)
	// 1100 0000 1000 0000-> 0 which should be invalid for 2 byte UTF-8
	// c 0 8 0
	fmt.Println("\nInvalid Decoding - replacement character �")
	invalidDecoding := "\xc0\x80"
	for i := 0; i < len(invalidDecoding); {
		r, size := utf8.DecodeRuneInString(invalidDecoding[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	// � -> /uFFFD character replacement character
	// �

	fmt.Println("\nRunes more convenient - uniform size [4bytes] integers")

	// Rune conversion

	s := "プログラム"
	// % x (with space) inserts space between a hex bytes
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0" -> encoded
	r := []rune(s)         // **conversion decodes**
	fmt.Printf("%x\n", r)  // "[30d7 30ed 30b0 30e9 30e0]" -> decoded Unicode code points

	fmt.Println(string(r)) // "プログラム" -> string(r) -> **conversion encodes**

	// converting from int to string
	//-> int is understood or int or non-decoded code point to be decoded
	// **1.15 warning in go vet** -> flag to disable "stringintconv"
	fmt.Println(string(65)) // "A", not "65"
	// new correct way
	fmt.Println(string(rune(65)))
	fmt.Println(string(0x4eac))  // "京"
	fmt.Println(string(1234567)) // "�" -> invalid rune
}

// No need to Decode to get to the code point number representation for these operations
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) &&
		s[:len(prefix)] == prefix // byte comparison
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) &&
		s[len(s)-len(suffix):] == suffix // byte comparison
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

// 3.5.4 Strings and Byte Slices
func stringsByteSlices() {
	fmt.Println(basename("a/b/c.go")) // "c"
	fmt.Println(basename("c.d.go"))   // "c.d"
	fmt.Println(basename("abc"))      // "abc"

	fmt.Println(basenameWithLastIndex("a/b/c.go")) // "c"
	fmt.Println(basenameWithLastIndex("c.d.go"))   // "c.d"
	fmt.Println(basenameWithLastIndex("abc"))      // "abc"

	fmt.Println(comma("12000000"))

	fmt.Println("\nImmutability of string in convertion to/from byte slice")
	s := "abc"
	b := []byte(s) // copies values of string as not to overwrite them
	b[0] = 'd'
	s2 := string(b) // also copies to ensure immutability
	fmt.Println(s)  // abc
	fmt.Println(s2) // dbc

	// bytes.Buffer type that serves as a string builder
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"

}

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c,a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basenameWithLastIndex(s string) string {

	directoryLastIndex := strings.LastIndex(s, "/")

	// no need for if due to -1+1=0 beginning of string
	// ... smart code
	s = s[directoryLastIndex+1:]

	if dotIndex := strings.LastIndex(s, "."); dotIndex != -1 {
		s = s[:dotIndex]
	}
	return s
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		// writes to writer -> which is the Buffer
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
