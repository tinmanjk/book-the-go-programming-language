package main

import (
	"fmt"

	tc "github.com/tinmanjk/tgpl/ch02-programStructure/06-packagesAndFiles/tempconv"
)

func main() {
	// because of the String method it prints "°C" suffix on the constant of the type Celsius
	fmt.Printf("Brrrr! %v\n", tc.AbsoluteZeroC) // "Brrrr! -273.15°C"
	fmt.Println(tc.CToF(tc.BoilingC))           // "212°F"

	fmt.Println("2.6.1 Imports")
	// 1. Import paths -> "github.com/tinmanjk/tgpl/ch02-programStructure/06-packagesAndFiles/tempconv"
	// the language specification doesn't define where the import path strings come from
	// it's the go tooling's job to interpret them

	// go tool ->  an import path denotes a directory containing one
	// or more Go source files that together make up the package.

	// 2. Package name -> tempconv (convention last segment of import path)

	// goimports tool automatically adds imports -> perhaps integrated well with gopls and others into vscode

	fmt.Println("----")

}
