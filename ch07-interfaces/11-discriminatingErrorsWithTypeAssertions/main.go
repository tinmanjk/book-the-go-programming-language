package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"syscall"
)

func main() {
	fmt.Println("Discriminating Errors With Type Assertions")

	_, err := os.Open("/no/such/file")
	fmt.Println(err)                            // "open /no/such/file: No such file or directory" e.Op+e.Path+e.err.Error()
	fmt.Printf("%#v\n", err)                    // Output: &fs.PathError{Op:"open", Path:"/no/such/file", Err:0x2}
	fmt.Println(os.IsNotExist(err))             // "true" -> !!!! must be checked HERE only before propagating the error ... the type information will be lost otherwise
	fmt.Println(errors.Is(err, fs.ErrNotExist)) // "true" -> new way
}

// !! stdlib code has changed to reflect the new errors.Is mechanic for wrapping errors (comments below, "old" code from book)
var ErrNotExist = errors.New("file does not exist") // same logic actually used in fs package that os depends upon

// IsNotExist returns a boolean indicating whether the error is known to
// report that a file or directory does not exist. It is satisfied by
// ErrNotExist as well as some syscall errors.
//
// This function predates errors.Is. It only supports errors returned by
// the os package. New code should use errors.Is(err, fs.ErrNotExist). !!!!!
func IsNotExist(err error) bool {
	if pe, ok := err.(*PathError); ok {
		err = pe.Err // extracting the field error for the interface comparison below
	}
	return err == syscall.ENOENT || err == ErrNotExist
}

// type PathError = fs.PathError -> type aliasing it is the SAME EXACT TYPE, not a NAMED TYPE

// PathError records an error and the operation and file path that caused it.
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
