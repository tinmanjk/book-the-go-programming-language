package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Println("Complex Numbers")

	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	// (1+2i)*(3+4i)
	// 3 + 4i + 6i + 8i² (i² = -1)
	// 3+10i-8 = -5 + 10i
	fmt.Println(x * y)       // "(-5+10i)"
	fmt.Println(real(x * y)) // "-5"
	fmt.Println(imag(x * y)) // "10"

	// imaginary literal -> i with 0 real component
	fmt.Println(1i) // "(0+1i)"
	{
		x := 1 + 2i // complex 128
		y := 3 + 4i // complex 128
		fmt.Println(x, y)
	}

	// cmplx package
	fmt.Println(cmplx.Sqrt(-1)) // "(0+1i)" -> i = sqrt(-i)
}
