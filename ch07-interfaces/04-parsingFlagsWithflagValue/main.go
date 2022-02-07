package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/tinmanjk/tgpl/ch02-programStructure/05-typeDeclarations/tempconv"
)

// from package flag

// Value is the interface to the value stored in a flag.
type Value interface {
	String() string   // also satisfies fmt.Stringer
	Set(string) error // the inverse of String -> parses the string to the value
}

// -- time.Duration Value -> for interface satisfaction
// a thin wrapper over the methods of time package
// type durationValue time.Duration -> name type from other package for methods

// -> below conversion from time.Duration to durationValue with a pointer
// func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
// *p = val
// return (*durationValue)(p)
// !!! -> pointer type conversion possible -> extra () because of ambiguity
// https://go.dev/ref/spec#Conversions
// }

// -> reusing time.ParseDuration for setting the value - same underlying type though
// func (d *durationValue) Set(s string) error {
// 	v, err := time.ParseDuration(s)
// 	if err != nil {
// 		err = errParse
// 	}
// 	*d = durationValue(v)
// 	return err
// }

// func (d *durationValue) Get() interface{} { return time.Duration(*d) }

// -> reuses the String method defined for time.Duration by converting back to the type
// func (d *durationValue) String() string { return (*time.Duration)(d).String() }

func main() {
	// internal conversion to durationValue for extension methods to satisfy time.Value interface
	// can be in main, not necessarily a package-level variable as in example code
	var period = flag.Duration("period", 1*time.Second, "sleep period")
	var temp = CelsiusFlag("temp", 20.0, "the temperature")

	// after all flags are defined and before are used
	flag.Parse()

	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
	fmt.Print(temp)
}

// embedding vs naming an underlying type ?
// naming does not allow for getting the defined methods
// embedding allows for method promotion -> such as String
// compare this with **durationValue** from flag package
type celsiusFlag struct {
	tempconv.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed -> won't match the switch
	// essentialy parsing s according to the format with out variables
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
// called from func (f *FlagSet) parseOne() (bool, error)
func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius // &(f.Celsius) -> . has precedence over &
}
