package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Mine: form 'x.(T)' - x MUST be interface value, T can be both interface/concrete type")
	// s := os.Stdout
	// var w io.Writer = s.(io.Writer) // compile time error: s (variable of type *os.File) is not an interface
	fmt.Println("1. From interface TO concrete type - no runtime methods - itab address comparison only")
	{
		//lint:ignore S1021 ...
		var w io.Writer // just nil value for interface value

		// main.go:14      0x4a4905        440f117c2448                    movups xmmword ptr [rsp+0x48], xmm15

		w = os.Stdout // *os.File -> still wraps in an **interface value** with type/data double-word pointers
		// https://stackoverflow.com/questions/1658294/whats-the-purpose-of-the-lea-instruction -> calculate runtime address
		// **ptr** operator is necessary with indirect operands for dereferencing - see Irvine 4.4.1 "Using PTR with Indirect Operands"
		// but what does ptr do with lea?!? see next for my topic in stackoverflow
		// https://stackoverflow.com/questions/71085570/what-does-the-ptr-operator-do-as-part-of-operand-to-lea
		// - lea rsi, ptr [rip+0x26b0f] -> ptr is REDUNDANT - we just calculate the address -> i.e. &operation

		// main.go:18      0x4a490b        488b1556500a00                  mov rdx, qword ptr [os.Stdout] // dereferencing -> * - still a pointer *os.File
		// main.go:18      0x4a4912        488d350f6b0200                  lea rsi, ptr [rip+0x26b0f] // taking the address -> &
		// main.go:18      0x4a4919        4889742458                      mov qword ptr [rsp+0x58], rsi // type descr (simple pointer)
		// main.go:18      0x4a491e        4889542460                      mov qword ptr [rsp+0x60], rdx // value (dereferenced pointer) -> qword ptr [os.Stdout]

		// !! Conversion from interface to concrete type (concrete to concrete -> conversion i.e. int32(x))
		// x.(T) -> panics if conversion not possible, otherwise gives you an instance of that concrete type
		f := w.(*os.File) // success: f == os.Stdout -> **EXTRACTS THE VALUE FROM THE INTERFACE** -> here the value is an address which is just copied
		// ** see below example with non-pointer extracting - **copying all values from the interface value into another variable**
		// main.go:27      0x4a49d6        488b442458                      mov rax, qword ptr [rsp+0x58] // type descr
		// main.go:27      0x4a49db        488b542460                      mov rdx, qword ptr [rsp+0x60] // value
		// main.go:27      0x4a49e0        4c8d05416a0200                  lea r8, ptr [rip+0x26a41] // rip has advanced CE bytes see next comment
		// !! 0x4a49e0 - 0x4a4912 = CE see instruction pointer [rip+0x26b0f] - [rip+0x26a41]  = CE...so the effective address is **THE SAME** -> *os.File
		// main.go:27      0x4a49e7        4c39c0                          cmp rax, r8 // compares the type descr addresses
		// main.go:27      0x4a49ea        7405                            jz 0x4a49f1 // if the same addresses / pointers - skip panic
		// main.go:27      0x4a49ec        e984000000                      jmp 0x4a4a75 -> panic
		// main.go:31      0x4a49f1        4889542428                      mov qword ptr [rsp+0x28], rdx -> WE EXTRACT the Interface Value - **SAME** pointer

		// below is jmp panic code is PANIC code
		// main.go:27      0x4a4a75        488d1d645d0100                  lea rbx, ptr [rip+0x15d64]
		// main.go:27      0x4a4a7c        488d0d5dc40000                  lea rcx, ptr [rip+0xc45d]
		// main.go:27      0x4a4a83        e8d855f6ff                      call $runtime.panicdottypeI
		fmt.Printf("%T\n", f) // *os.File
		// !! -> if type descriptor addresses are **NOT THE SAME** betwee interface dynamic type and concrete type
		// c := w.(*bytes.Buffer) //panic: interface conversion: io.Writer **is** *os.File, not *bytes.Buffer
	}

	fmt.Println("\nMine: Extracting from interface value to type-asserted **non-pointer** concrete value:\n- copy the interface values to the type-asserted new variable ")
	{
		fieldSt := fieldStruct{1, 2, 3, 4}
		// main.go:51      0x4a4aca        48c744242801000000              mov qword ptr [rsp+0x28], 0x1
		// main.go:51      0x4a4ad3        48c744243002000000              mov qword ptr [rsp+0x30], 0x2
		// main.go:51      0x4a4adc        48c744243803000000              mov qword ptr [rsp+0x38], 0x3
		// main.go:51      0x4a4ae5        48c744244004000000              mov qword ptr [rsp+0x40], 0x4
		var interfaceValue interface{} = fieldSt
		// main.go:52      0x4a4b15        48c784248800000001000000        mov qword ptr [rsp+0x88], 0x1 // copy to the stack instead of memaloc
		// main.go:52      0x4a4b21        48c784249000000002000000        mov qword ptr [rsp+0x90], 0x2
		// main.go:52      0x4a4b2d        48c784249800000003000000        mov qword ptr [rsp+0x98], 0x3
		// main.go:52      0x4a4b39        48c78424a000000004000000        mov qword ptr [rsp+0xa0], 0x4
		// main.go:52      0x4a4b45        488d15f4160100                  lea rdx, ptr [rip+0x116f4]
		// main.go:52      0x4a4b4c        48899424f0000000                mov qword ptr [rsp+0xf0], rdx -> type descriptor
		// main.go:52      0x4a4b54        488d942488000000                lea rdx, ptr [rsp+0x88]
		// main.go:52      0x4a4b5c        48899424f8000000                mov qword ptr [rsp+0xf8], rdx -> values - pointer to [rsp+0x88]
		convertBack := interfaceValue.(fieldStruct)
		// main.go:53      0x4a4b70        488b9424f8000000                mov rdx, qword ptr [rsp+0xf8] // pointer to [rsp+0x88] actual values
		// main.go:53      0x4a4b78        488b8424f0000000                mov rax, qword ptr [rsp+0xf0] // type descriptor
		// main.go:53      0x4a4b80        488d1db9160100                  lea rbx, ptr [rip+0x116b9]
		// main.go:53      0x4a4b87        4839d8                          cmp rax, rbx // type assertion
		// main.go:53      0x4a4b8a        7405                            jz 0x4a4b91
		// main.go:53      0x4a4b8c        e907010000                      jmp 0x4a4c98
		// main.go:53      0x4a4b91        488b0a                          mov rcx, qword ptr [rdx] // copying the values from [rsp+0x88]
		// main.go:53      0x4a4b94        488b7208                        mov rsi, qword ptr [rdx+0x8]
		// main.go:53      0x4a4b98        488b7a10                        mov rdi, qword ptr [rdx+0x10]
		// main.go:53      0x4a4b9c        488b5218                        mov rdx, qword ptr [rdx+0x18]
		// main.go:53      0x4a4ba0        48894c2468                      mov qword ptr [rsp+0x68], rcx
		// main.go:53      0x4a4ba5        4889742470                      mov qword ptr [rsp+0x70], rsi
		// main.go:53      0x4a4baa        48897c2478                      mov qword ptr [rsp+0x78], rdi
		// main.go:53      0x4a4baf        4889942480000000                mov qword ptr [rsp+0x80], rdx
		fmt.Println(convertBack)
	}

	fmt.Println("\n2. From interface TO interface type")
	{
		var w io.Writer
		// main.go:88      0x4a5a05        440f11bc2410010000              movups xmmword ptr [rsp+0x110], xmm15
		w = os.Stdout
		// main.go:89      0x4a5a0e        488b15935f0a00                  mov rdx, qword ptr [os.Stdout]
		// main.go:89      0x4a5a15        488d351c720200                  lea rsi, ptr [rip+0x2721c]
		// main.go:89      0x4a5a1c        4889b42410010000                mov qword ptr [rsp+0x110], rsi
		// main.go:89      0x4a5a24        4889942418010000                mov qword ptr [rsp+0x118], rdx
		// modified from book (originally ReadWrite and mysterious type ByteCounter)
		// checking if w's **dynamic type** satisfies io.ReadWriteCloser
		fmt.Println("\n2.1 Assigning from **smaller interface value** to **bigger interface variable** -> Type asserting: \n" +
			"- Inspecting the dynamic value of the interface value against the bigger interface -> $runtime.assertI2I")
		rwc := w.(io.ReadWriteCloser) // success: *os.File has both Read, Write and Close -> !! values for dynamic type/dynamic value are COPIED to new interfacea value
		// Assembly -> call to $runtime.assertI2I
		// main.go:92      0x4a5a2c        440f11bc2460010000              movups xmmword ptr [rsp+0x160], xmm15
		// main.go:92      0x4a5a35        488b942418010000                mov rdx, qword ptr [rsp+0x118]
		// main.go:92      0x4a5a3d        48899424b8000000                mov qword ptr [rsp+0xb8], rdx // copies value temporarily
		// main.go:92      0x4a5a45        488b9c2410010000                mov rbx, qword ptr [rsp+0x110]
		// main.go:92      0x4a5a4d        488d054ccf0000                  lea rax, ptr [rip+0xcf4c]
		// Arguments for $runtime.assertI2I  rax = io.ReadWriteCloser, rbx = dynamic type
		// under src/runtime/iface.go for debugging

		// func assertI2I(inter *interfacetype, tab *itab) *itab {
		// used to be func assertI2I(inter *interfacetype, i iface) (r iface) same as convI2I -> 17-Feb-21 10:14:21 AM
		// 	if tab == nil {
		// 		// explicit conversions require non-nil interface value.
		// 		panic(&TypeAssertionError{nil, nil, &inter.typ, ""})
		// 	}
		// 	if tab.inter == inter {
		// 		return tab
		// 	}
		// 	return getitab(inter, tab._type, false) -> !! real satisfaction stuff -> panics in getitab
		// }
		// main.go:92      0x4a5a54        e84749f6ff                      call $runtime.assertI2I // no longer simple compare, PANIC INSIDE
		// Returns from $runtime.assertI2I  -> rax = dynamic type -> same type BUT different location -> **in another interface table**
		// main.go:92      0x4a5a59        4889842460010000                mov qword ptr [rsp+0x160], rax // same dynamic type -> DIFFERENT LOCATION
		// main.go:92      0x4a5a61        488b9424b8000000                mov rdx, qword ptr [rsp+0xb8]
		// main.go:92      0x4a5a69        4889942468010000                mov qword ptr [rsp+0x168], rdx // dynamic value
		fmt.Printf("%T\n", rwc) // *os.File
		w = new(bytes.Buffer)   // *bytes.Buffer does have Write, BUT NO Close
		fmt.Printf("%T\n", w)   // *bytes.Buffer
		// rwc = w.(io.ReadWriteCloser) // panic: interface conversion: *bytes.Buffer is not io.ReadWriteCloser: missing method Close

		fmt.Println("\n2.2 Assigning from **bigger interface value** to **small interface variable** -> Implicit Converting - no x.(T) syntax: \n" +
			"- Inspecting the dynamic value of the interface value against the smaller interface")

		fmt.Println("2.2.1 Simple assignment -> w = rwc or w = io.Writer(rwc) ->  $runtime.convI2I")
		w = rwc // implicit conversion see below ->  $runtime.convI2I
		{       // Assembly -> call to $runtime.convI2I
			// checking the interface value against the io.Writer interface see below -> similar to the one above

			// main.go:134     0x4a5e58        488b9c2458010000                mov rbx, qword ptr [rsp+0x158]
			// main.go:134     0x4a5e60        488b8c2460010000                mov rcx, qword ptr [rsp+0x160]
			// main.go:134     0x4a5e68        488d05f1d20000                  lea rax, ptr [rip+0xd2f1]
			// main.go:134     0x4a5e6f        e82c45f6ff                      call $runtime.convI2I
			{
				// Arguments for $runtime.convI2I  rax = io.ReadWriteCloser, rbx = dynamic type, rcx = dynamic value
				// under src/runtime/iface.go for debugging -> ACTUALLY CHANGED the itab -> logical (see book that it stays the same)
				// convI2I returns the **new itab** to be used for the destination value
				// when converting a value with itab src to the dst interface.

				// func convI2I(inter *interfacetype, i iface) (r iface) { // iface = itab and idata -> ACTUALLY CHANGED - August 2021
				// 	tab := i.tab
				// 	if tab == nil { // no panic with nil
				// 		return
				// 	}
				// 	if tab.inter == inter {
				// 		r.tab = tab
				// 		r.data = i.data
				// 		return
				// 	}
				// 	r.tab = getitab(inter, tab._type, false)
				// 	r.data = i.data
				// 	return
				// }
			}
			// -> rax = tab, rbx = data
			// main.go:134     0x4a5e74        4889842448010000                mov qword ptr [rsp+0x148], rax
			// main.go:134     0x4a5e7c        48899c2450010000                mov qword ptr [rsp+0x150], rbx
		}
		w = io.Writer(rwc) // itdentical to above uses $runtime.convI2I
		{
			// Assembly -> identical to above
			// main.go:136     0x4a5e84        488b9c2458010000                mov rbx, qword ptr [rsp+0x158]
			// main.go:136     0x4a5e8c        488b8c2460010000                mov rcx, qword ptr [rsp+0x160]
			// main.go:136     0x4a5e94        488d05c5d20000                  lea rax, ptr [rip+0xd2c5]
			// main.go:136     0x4a5e9b        0f1f440000                      nop dword ptr [rax+rax*1], eax
			// main.go:136     0x4a5ea0        e8fb44f6ff                      call $runtime.convI2I
			// main.go:136     0x4a5ea5        4889842448010000                mov qword ptr [rsp+0x148], rax
			// main.go:136     0x4a5ead        48899c2450010000                mov qword ptr [rsp+0x150], rbx
		}

		fmt.Printf("%T\n", w) // *os.File

		fmt.Println("2.2.2 Type assertion -> w = rwc.(io.Writer) - $runtime.assertI2I")

		w = rwc.(io.Writer) // uses $runtime.assertI2I -> same as smaller to bigger
		{
			// Assembly
			// main.go:184     0x4a6067        440f11bc24f8010000              movups xmmword ptr [rsp+0x1f8], xmm15
			// main.go:184     0x4a6070        488b942480010000                mov rdx, qword ptr [rsp+0x180]
			// main.go:184     0x4a6078        48899424b8000000                mov qword ptr [rsp+0xb8], rdx
			// main.go:184     0x4a6080        488b9c2478010000                mov rbx, qword ptr [rsp+0x178]
			// main.go:184     0x4a6088        488d05d1d00000                  lea rax, ptr [rip+0xd0d1]
			// main.go:184     0x4a608f        e8ac43f6ff                      call $runtime.assertI2I
			// main.go:184     0x4a6094        48898424f8010000                mov qword ptr [rsp+0x1f8], rax
			// main.go:184     0x4a609c        488b9424b8000000                mov rdx, qword ptr [rsp+0xb8]
			// main.go:184     0x4a60a4        4889942400020000                mov qword ptr [rsp+0x200], rdx
			// main.go:184     0x4a60ac        4889842468010000                mov qword ptr [rsp+0x168], rax
			// main.go:184     0x4a60b4        4889942470010000                mov qword ptr [rsp+0x170], rdx
		}
		// !! compile time check when smaller is clearly not satisfying bigger
		// rwc = io.ReadWriteCloser(w) // compile time error if **conversion** from smaller to bigger
		// !!

		fmt.Printf("%T\n", w) // *os.File

		fmt.Println("2.2.3 Nil handling difference")
		var rw io.ReadWriter // nil interface value
		w = rw               // call $runtime.convI2I -> no panic if nil value - just returns nil
		fmt.Println(w)       // nil -> no panic
		// w = b.(io.Writer) // call $runtime.assertI2I  panic: interface conversion: interface is nil, not io.Writer
		fmt.Println("panic: interface conversion: interface is nil, not io.Writer")

	}

	fmt.Println("\n3. Testing val,ok :=")
	{
		fmt.Println("\n3.1 Concrete types - extraction")
		{
			var w io.Writer = os.Stdout
			f, ok := w.(*os.File) // success: ok, f == os.Stdout
			fmt.Println(f, ok)

			b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil (no panic as normally with concrete type)
			{
				// main.go:223     0x4a6585        48c78424b001000000000000        mov qword ptr [rsp+0x1b0], 0x0
				// main.go:223     0x4a6591        488b9424e0010000                mov rdx, qword ptr [rsp+0x1e0]
				// main.go:223     0x4a6599        488bb424e8010000                mov rsi, qword ptr [rsp+0x1e8]
				// main.go:223     0x4a65a1        488d3d107a0200                  lea rdi, ptr [rip+0x27a10]
				// main.go:223     0x4a65a8        4839fa                          cmp rdx, rdi // still comparison of type
				// main.go:223     0x4a65ab        7402                            jz 0x4a65af
				// main.go:223     0x4a65ad        eb07                            jmp 0x4a65b6
				// main.go:223     0x4a65af        b801000000                      mov eax, 0x1
				// main.go:223     0x4a65b4        eb06                            jmp 0x4a65bc
				// main.go:223     0x4a65b6        31c0                            xor eax, eax // !! zero value instead of panic!!
				// main.go:223     0x4a65b8        31f6                            xor esi, esi
				// main.go:223     0x4a65ba        eb00                            jmp 0x4a65bc
				// main.go:223     0x4a65bc        4889b424b0010000                mov qword ptr [rsp+0x1b0], rsi
				// main.go:223     0x4a65c4        8844242f                        mov byte ptr [rsp+0x2f], al
				// main.go:223     0x4a65c8        488b9424b0010000                mov rdx, qword ptr [rsp+0x1b0]
				// main.go:223     0x4a65d0        48899424c8000000                mov qword ptr [rsp+0xc8], rdx
				// main.go:223     0x4a65d8        0fb654242f                      movzx edx, byte ptr [rsp+0x2f]
				// main.go:223     0x4a65dd        8854242e                        mov byte ptr [rsp+0x2e], dl
			}
			fmt.Println(b, ok) // nil, false ( nil is zero-value here)

			// shadowing with if statement declaration - quick use
			if w, ok := w.(*os.File); ok {
				fmt.Println(w)
			}
		}

		fmt.Println("\n3.2 Interface types")
		{
			var w io.Writer = os.Stdout
			rwc, ok := w.(io.ReadWriteCloser) // uses $runtime.assertI2I2 (not asserI2I) which does not panic
			{
				// main.go:256     0x4a68d1        440f11bc24c8020000              movups xmmword ptr [rsp+0x2c8], xmm15
				// main.go:256     0x4a68da        488b9c24e8010000                mov rbx, qword ptr [rsp+0x1e8]
				// main.go:256     0x4a68e2        488b8c24f0010000                mov rcx, qword ptr [rsp+0x1f0]
				// main.go:256     0x4a68ea        488d05afd00000                  lea rax, ptr [rip+0xd0af]
				// main.go:256     0x4a68f1        e80a3cf6ff                      call $runtime.assertI2I2
				// main.go:256     0x4a68f6        48898424c8020000                mov qword ptr [rsp+0x2c8], rax
				// main.go:256     0x4a68fe        48899c24d0020000                mov qword ptr [rsp+0x2d0], rbx
				// main.go:256     0x4a6906        4885c0                          test rax, rax
				// main.go:256     0x4a6909        0f9544242f                      setnz byte ptr [rsp+0x2f]
				// main.go:256     0x4a690e        488b9424d0020000                mov rdx, qword ptr [rsp+0x2d0]
				// main.go:256     0x4a6916        488bb424c8020000                mov rsi, qword ptr [rsp+0x2c8]
				// main.go:256     0x4a691e        4889b42418020000                mov qword ptr [rsp+0x218], rsi
				// main.go:256     0x4a6926        4889942420020000                mov qword ptr [rsp+0x220], rdx
				// main.go:256     0x4a692e        0fb654242f                      movzx edx, byte ptr [rsp+0x2f]
				// main.go:256     0x4a6933        8854242d                        mov byte ptr [rsp+0x2d], dl
			}
			fmt.Println(rwc, ok)
		}

	}

}

type fieldStruct struct {
	a, b, c, d int64
}
