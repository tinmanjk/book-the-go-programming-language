package main

import (
	"fmt"
	"math"
	"time"
)

// untyped package level constants
const (
	e  = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
)

func main() {

	// pi = 3.15 // compiler error
	fmt.Println(e, pi)
	// constant in type
	var p [IPv4Len]byte
	fmt.Println(p)

	// const declaration inside a function
	const noDelay time.Duration = 0 // explicit typed declaration
	// time.Duration is named type with underlying int64
	const timeout = 5 * time.Minute // inferred but still typed
	// -> %T verb for the type of a const/variable
	fmt.Printf("%T %[1]v\n", noDelay)     //	"time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     //	"time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) //	"time.Duration 1m0s"

	const (
		a = 1
		b // b = 1
		c = 2
		d // d = 2
	)

	fmt.Println(a, b, c, d) // "1 1 2 2"

	// 3.6.1 The Constant Generator iota
	fmt.Println("\n3.6.1 The Constant Generator iota")
	iotaExamples()

	// 3.6.2 Untyped Constants
	fmt.Println("\n3.6.2 Untyped Constants")
	untypedConstants()
}

const IPv4Len = 4

type Weekday int

const (
	// constant generator iota -> succesive numbering starting 0
	Monday    Weekday = iota
	Tuesday           // Weekday = 1
	Wednesday         // Weekday = 2
	Thursday
	Friday
	Saturday
	Sunday
)

type Flags uint

// more complex exppression using iota
const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // 2 supports broadcast access capability
	FlagLoopback                       // 4 is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast фaccess capability
)

func iotaExamples() {
	// Flags example
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001  true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}

func IsUp(v Flags) bool {
	return v&FlagUp ==
		FlagUp
}
func TurnDown(v *Flags) { *v &^= FlagUp }
func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}
func IsCast(v Flags) bool {
	return v&
		(FlagBroadcast|FlagMulticast) != 0
}

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1   << 64)
	YiB // 1208925819614629174706176
)

func untypedConstants() {
	// too big to store in 64bit int -> however expression is possible
	fmt.Println(YiB / ZiB) // "1024"

	// one constant can be assigned to many types
	// no need for type conversions
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)

	// constant expressions (unnamed)
	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9) // "100"; (f -32) * 5 is a float64
	//lint:ignore SA4025 ...
	fmt.Println(5 / 9 * (f - 32))     // "0"; 5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100";5.0/9.0 is an untyped float

	// conversion from constants/constant expressions to type
	{
		const untypedComplex = 3 + 0i
		const untypedInt = 3
		const untypedFloat = 1e123
		const untypedRUne = 'a'

		// implicit conversion from untyped constants to a variable
		var f float64 = untypedComplex // untyped complex ->float64
		fmt.Println(f, untypedComplex)
		f = untypedInt // untyped integer ->float64
		fmt.Println(f, untypedInt)
		f = untypedFloat // untyped floating-point -> float64
		fmt.Println(f, untypedFloat)
		f = untypedRUne // untyped rune -> float64
		fmt.Println(f, untypedRUne)

		// above equivalent to this explicit conversion
		f = float64(3 + 0i)
		f = float64(2)
		f = float64(1e123)
		f = float64('a')

		// Compile time Conversion Errors
		const deadbeef = 0xdeadbeef // untyped int with value 3735928559
		fmt.Println(deadbeef)
		const a = uint32(deadbeef) // uint32 with value 3735928559
		fmt.Println(a)
		const b = float32(deadbeef) // float32 with value 3735928576 (rounded up)
		fmt.Printf("%f\n", b)
		const c = float64(deadbeef) // float64 with value 3735928559 (exact)
		fmt.Printf("%f\n", c)
		// const d = int32(deadbeef)   // compile error: constant overflows int32
		// const e = float64(1e309) // compile error: constant overflows float64
		// const f = uint(-1)       // compile error: 	constant underflows uint

	}

	// flavor of untyped constant determines type with implicit conversions
	{
		i := 0      // untyped integer; 		implicit int(0)
		r := '\000' // untyped rune; 	 	   	implicit rune('\000')
		f := 0.0    // untyped floating-point; 	implicit float64(0.0)
		c := 0i     // untyped complex; 		implicit complex128(0i)

		fmt.Println(i, r, f, c)
		// explicit declaration needed if implicit conversion to int is not wished
		// identical assembly generated
		{
			var i = int8(0) // mov byte ptr [rsp+0x2e], 0x0
			fmt.Println(i)
		}
		{
			var i int8 = 0 // mov byte ptr [rsp+0x2d], 0x0
			fmt.Println(i)
		}
	}

	// Dynamic type of a constant value
	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)

}
