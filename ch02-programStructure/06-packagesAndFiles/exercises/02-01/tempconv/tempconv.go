// Package tempconv performs Celsius and Fahrenheit conversions.
// for go doc -> can be placed in a separate go.doc file -> see chapter 10
// only ONE file should have the doc
package tempconv

import "fmt"

// Exercise 2.1: Add types, constants, and functions to tempconv for
// processing temperatures in the Kelvin scale, where zero Kelvin is
// −273.15°C and a difference of 1K has the same magnitude as 1°C.

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	CelsiusZeroK  Kelvin  = 273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
