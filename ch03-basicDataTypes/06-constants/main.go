package main

import (
	"fmt"
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
