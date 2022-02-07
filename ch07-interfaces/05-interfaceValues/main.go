package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Interface Values")

	// https://research.swtch.com/interfaces - also

	// 128-bit for the local interface value starting **[rsp+0x38]**,
	// 128bit -> **two 64bit-words for pointers** for **type descriptor** and **value**
	var w io.Writer //zero value = nil value based on the dynamic type value (also nil)
	// main.go:14      0x4a3805        440f117c2438                    movups xmmword ptr [rsp+0x38], xmm15
	// implementaion detail ? - apparently xmm15 holds nil value...?

	fmt.Println("----Nil Checks") // Mine - investigating the nil check mechanism.
	// Type Descriptor is checked for nil -> NOT value
	{
		// function variable to disable compiler optimizations for nil check
		checkNil := func(writer io.Writer) {
			result := writer == nil
			// main.go:21      0x4a3b4e        4889442418      mov qword ptr [rsp+0x18], rax (type)
			// main.go:21      0x4a3b53        48895c2420      mov qword ptr [rsp+0x20], rbx (value)
			// main.go:22      0x4a3b5d        48837c241800    cmp qword ptr [rsp+0x18], 0x0 // compare type to 0, not value

			// -> ** the comparison is just the type descriptor **, not the value
			fmt.Println(result)
		}

		m := (*mockWriter)(nil)
		// main.go:35      0x4a38e5        48c744242000000000              mov qword ptr [rsp+0x20], 0x0
		fmt.Println(m == nil) // true

		w = m // type is NOT nil, value is nil from above
		// main.go:37      0x4a3965        488b542420                      mov rdx, qword ptr [rsp+0x20] -> 0x0 see above
		// main.go:37      0x4a396f        488d35a26f0200                  lea rsi, ptr [rip+0x26fa2] // type
		// main.go:37      0x4a3976        4889742478                      mov qword ptr [rsp+0x78], rsi
		// main.go:37      0x4a397b        4889942480000000                mov qword ptr [rsp+0x80], rdx -> 0x0 see above
		checkNil(w) // false -> checks JUST the type descriptor NOT the value
		w = nil
		checkNil(w) // true
		fmt.Println("----/Nil Checks")
	}

	// w.Write([]byte("hello")) // panic: nil pointer dereference

	// type  [0x38] = [rip+0x26f  df] -> type descriptor -> *os.File
	// value [0x40] = [os.Stdout]
	// w = io.Writer(os.Stdout) // explicit
	w = os.Stdout // implicit conversion
	// -> places **the POINTER in the value slot**, not a pointer to the pointer
	// main.go:55      0x4a3a52        488b154f6f0a00                  mov rdx, qword ptr [os.Stdout]
	// main.go:55      0x4a3a59        488d3518800200                  lea rsi, ptr [rip+0x28018]
	// main.go:55      0x4a3a60        4889b424e0000000                mov qword ptr [rsp+0xe0], rsi
	// main.go:55      0x4a3a68        48899424e8000000                mov qword ptr [rsp+0xe8], rdx

	// dynamic dispatch call
	hello := []byte("hello\n")
	// main.go:63      0x4a3aa5        48899c2410010000                mov qword ptr [rsp+0x110], rbx
	// main.go:63      0x4a3aad        48c784241801000006000000        mov qword ptr [rsp+0x118], 0x6
	// main.go:63      0x4a3ab9        48c784242001000006000000        mov qword ptr [rsp+0x120], 0x6
	w.Write(hello) // "hello" -> **effectively os.Stdout.Write([]byte("hello"))**
	// main.go:64      0x4a3ac5        488b9424e0000000                mov rdx, qword ptr [rsp+0xe0] -> type descriptor location
	// main.go:64      0x4a3acf        488b5218                        mov rdx, qword ptr [rdx+0x18] -> offset to the method possibly
	// main.go:64      0x4a3ad3        488b8424e8000000                mov rax, qword ptr [rsp+0xe8] -> value as the receiver of the method call (os.StdOut)
	// main.go:64      0x4a3adb        b906000000                      mov ecx, 0x6
	// main.go:64      0x4a3ae0        4889cf                          mov rdi, rcx
	// implicitly the offset of the byte array is rbx see above (possibly compiler optimization)
	// main.go:64      0x4a3ae3        ffd2                            call rdx

	// main.go:64      0x4a3ac5        488b9424e0000000                mov rdx, qword ptr [rsp+0xe0]
	// main.go:64      0x4a3acf        488b5218                        mov rdx, qword ptr [rdx+0x18]
	// main.go:64      0x4a3ae3        ffd2                            call rdx // -> !!!

	// type  [0x38] = [rip+0x26f  3d] -> type descriptor -> bytes.Buffer (pointer type?)
	// value [0x40] = rax (new allocation pointer to value)
	w = new(bytes.Buffer)
	// main.go:18      0x4a3823        488d0596fa0000                  lea rax, ptr [rip+0xfa96]
	// main.go:18      0x4a382a        e85193f6ff                      call $runtime.newobject -> returns in rax pointer
	// main.go:18      0x4a382f        4889442430                      mov qword ptr [rsp+0x30], rax -> temp variable for return of new?
	// main.go:18      0x4a3834        488d153d6f0200                  lea rdx, ptr [rip+0x26f3d]
	// main.go:18      0x4a383b        4889542438                      mov qword ptr [rsp+0x38], rdx
	// main.go:18      0x4a3840        4889442440                      mov qword ptr [rsp+0x40], rax

	w = nil // resets both components to nil
	// main.go:14      0x4a3805        440f117c2438                    movups xmmword ptr [rsp+0x38], xmm15

	// MINE -> value vs pointer for interface conversion
	fmt.Println("\nMine: Value vs pointer for interface conversion")
	valueMock := mockWriter{1, 2, 3, 4}

	// copies values of struct first to a temp variable
	// then calls $runtime.convT2Inoptr to return the interface value
	w = valueMock
	// main.go:78      0x4a3ace        48c744244001000000              mov qword ptr [rsp+0x40], 0x1
	// main.go:78      0x4a3ad7        48c744244802000000              mov qword ptr [rsp+0x48], 0x2
	// main.go:78      0x4a3ae0        48c744245003000000              mov qword ptr [rsp+0x50], 0x3
	// main.go:78      0x4a3ae9        48c744245804000000              mov qword ptr [rsp+0x58], 0x4

	// main.go:78      0x4a3af2        488d054f6f0200                  lea rax, ptr [rip+0x26f4f] -> type
	// main.go:78      0x4a3af9        488d5c2440                      lea rbx, ptr [rsp+0x40] -> value
	// main.go:78      0x4a3b00        e89b68f6ff                      call $runtime.convT2Inoptr -> ?

	// func convT2Inoptr(tab *itab, elem unsafe.Pointer) (i iface) {
	// 	t := tab._type
	// 	x := mallocgc(t.size, t, false) //Allocate an object of size bytes. - get pointer
	// 	memmove(x, elem, t.size) // actual copying from the temp variable passed
	// 	i.tab = tab // same
	// 	i.data = x // the allocated pointer above
	// 	return

	// }
	// main.go:78      0x4a3b05        48898424b0000000                mov qword ptr [rsp+0xb0], rax -> type
	// main.go:78      0x4a3b0d        48899c24b8000000                mov qword ptr [rsp+0xb8], rbx -> value
	fmt.Println(w)

	pointerMock := &mockWriter{1, 2, 3, 4}
	w = pointerMock // just reuses pointer - does not call convT2Inoptr
	w.(*mockWriter).a = 333
	fmt.Println(pointerMock) // &{333 2 3 4}

	// Mine

	fmt.Println("\nComparing interface values - COMPARABLE but panic if dynamic type NOT comparable")
	var x interface{} = [3]int{1, 2, 3}
	//lint:ignore SA4000 ...
	fmt.Println(x == x) // true 1. Type equal 2. Dynamic values are equal accordin to the type equal
	// can be used as keys of a map/operands of switch
	// HOWEVER
	x = []int{1, 2, 3}
	// fmt.Println(x == x) // panic: comparing uncomparable type []int

	fmt.Println("\nDynamic type reporting via fmt package and T verb - reflection")
	{
		var w io.Writer
		fmt.Printf("%T\n", w) // "<nil>"

		w = os.Stdout
		fmt.Printf("%T\n", w) // "*os.File"

		w = new(bytes.Buffer)
		fmt.Printf("%T\n", w) // "*bytes.Buffer"
	}

}

type mockWriter struct {
	a, b, c, d int64
}

func (m mockWriter) Write(p []byte) (n int, err error) {
	return
}
