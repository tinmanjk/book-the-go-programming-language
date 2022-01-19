package main

import (
	"github.com/tinmanjk/tgpl/ch02-programStructure/06-packagesAndFiles/exercises/02-01/tempconv"

	"fmt"
)

func main() {

	fmt.Println(tempconv.CelsiusZeroK)
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"
	fmt.Println(tempconv.CToK(0))                     // 273.15°K
	fmt.Println(tempconv.KToC(0))                     // -273.15°C
}
