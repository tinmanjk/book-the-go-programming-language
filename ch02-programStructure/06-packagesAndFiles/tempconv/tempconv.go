// Package tempconv performs Celsius and Fahrenheit conversions.
// for go doc -> can be placed in a separate go.doc file -> see chapter 10
// only ONE file should have the doc
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func init() {
	fmt.Println("Initializing tempconv")
}
