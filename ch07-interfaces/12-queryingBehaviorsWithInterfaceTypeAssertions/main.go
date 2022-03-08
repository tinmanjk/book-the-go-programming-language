package main

import (
	"fmt"
	"io"
)

func main() {
	fmt.Println("Querying Behaviors With Interface Type Assertions")
	// since go 1.12 we have the **io.StringWriter** interface as part of the io public API
	// Commit hash:	112f28defcbd8f48de83f4502093ac97149b4da6 2018-10-03

}

//lint:ignore U1000 unused
func writeHeader(w io.Writer, contentType string) error {
	// issue is the conversion from string to []byte
	// -> memory allocation
	if _, err := w.Write([]byte("Content-Type: ")); err != nil {
		return err
	}
	if _, err := w.Write([]byte(contentType)); err != nil {
		return err
	}

	return nil
}

// writeString writes s to w.
// If w has a WriteString method, it is invoked instead of w.Write.
//lint:ignore U1000 ...
func writeString(w io.Writer, s string) (n int, err error) {
	// io.WriteString() // -> the same logic

	// !!! -> interface declaration INSIDE function
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	// type assertion to check if a method exists -> then use it
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) // avoid a copy
	}
	return w.Write([]byte(s)) // allocate temporary copy
}

//lint:ignore U1000 ...
func writeHeaderNew(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-Type: "); err != nil {
		return err
	}
	if _, err := writeString(w, contentType); err != nil {
		return err
	}
	return nil
}

// from fmt package at 1.5 -> dealing with empty interface and querying it for behaviors
//lint:ignore U1000 ...
func formatOneValue(x interface{}) string {
	if err, ok := x.(error); ok {
		return err.Error()
	}
	if str, ok := x.(fmt.Stringer); ok {
		return str.String()
	}
	// ...all other types... -> reflection see chapter 12
	return ""
}
