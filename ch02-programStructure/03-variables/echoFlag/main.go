package main

import (
	"flag"
	"fmt"
	"strings"
)

// needs to be accessed via -n (-n is implicitly true)
// otherwise flags are accessed via -n=true, -s=--- or -s="---" or -s --
// -h, -help, invalid flag, invalid value would cause the message to be printed
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()                               // needs to be called before flags can be used
	fmt.Print(strings.Join(flag.Args(), *sep)) // flag.Args -> non-flag arguments
	if !*n {
		fmt.Println()
	}
}
