package main

import (
	"fmt"

	"github.com/tinmanjk/tgpl/ch02-programStructure/06-packagesAndFiles/tempconv"
)

func main() {
	// because of the String method it prints "°C" suffix on the constant of the type Celsius
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"
	fmt.Println(tempconv.CToF(tempconv.BoilingC))     // "212°F"
}
