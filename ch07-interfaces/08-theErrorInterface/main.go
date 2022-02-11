package main

import (
	"errors"
	"fmt"
	"io"
	"syscall"
)

func main() {
	fmt.Println("The Error Interface")
	// builtin.go ->

	// type error interface {
	// 	Error() string
	// }
	err := errors.New("error message")
	fmt.Printf("%T->%[1]s\n", err) // *errors.errorString

	//lint:ignore SA4000 ...
	fmt.Println(New("EOF") == New("EOF"))    // true
	fmt.Println(errors.New("EOF") == io.EOF) // false (var EOF = errors.New("EOF") in io package) -> unique errors

	// prefer fmt.Errorf -> CONCEPTUAL IMPLEMENTATION (NO LONGER TRUE)
	// 	func Errorf(format string, args ...interface{}) error {
	//  return errors.New(Sprintf(format, args...))
	// }
	nonWrappingErr := fmt.Errorf("chaining: %v", err)
	fmt.Printf("%T->%[1]s\n", nonWrappingErr) // *errors.errorString->chaining: error message

	// 1.13 mechanics for wrapping/unwrapping
	wrappingErr := fmt.Errorf("chaining: %w", err) // the %w verb -> fmt -> print.go -> func (p *pp) handleMethods(verb rune) (handled bool)
	fmt.Printf("%T->%[1]s\n", wrappingErr)         // *fmt.wrapError->error message (** different type wrapError)
	// type wrapError struct {
	// 	msg string
	// 	err error
	// }
	// func (e *wrapError) Unwrap() error {	return e.err }

	// Is API to check if an error is wrapped in another one
	fmt.Println(errors.Is(nonWrappingErr, err)) // false
	fmt.Println(errors.Is(wrappingErr, err))    // true ( mind the order, first is where to search for the second as a match)
	unwrappedErr := errors.Unwrap(wrappingErr)
	fmt.Println(unwrappedErr == err) // true -> same pointers

	fmt.Println("\nErrors as number representation -> syscall.Errno")
	var errNo error = syscall.Errno(2) // creating an interface value
	fmt.Println(errNo.Error())         // "The system cannot find the file specified."
	// Error strings for invented errors -> Error method does a lookup based on the numeric code for the string representation
	// var errors = [...]string{
	fmt.Println(errNo)                   // "The system cannot find the file specified."
	fmt.Printf("%T -> %d", errNo, errNo) // syscall.Errno -> 2
}

// Modified version with values instead of pointers

// with values instead of pointers -> equality of errors based on string equality
func New(text string) error {
	return errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct { // -> so it cannot be changed - encapsulation of the message
	s string
}

func (e errorString) Error() string {
	return e.s
}
