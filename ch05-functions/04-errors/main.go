package main

import "fmt"

func main() {
	fmt.Println("Errors")
	// single "error" in form of bool to indicate success or failure
	cache := map[string]int{} // or a function call
	//lint:ignore SA4006 ...
	_, ok := cache["non-existent"]
	//lint:ignore SA9003 ...
	if !ok {
		// ...cache[key] does not exist...
	}

	// **built-in interface type** -> See Chapter 7
	// The error built-in interface type is the conventional interface for
	// representing an error condition, with the nil value representing no error.
	// type error interface {
	// 		Error() string
	// }
	var err error = fmt.Errorf("some error")
	fmt.Println(err.Error()) // "some error"
	fmt.Println(err)         // "some error"
	fmt.Printf("%v\n", err)  // "some error"

	// err can be nil = success, non-nil = failure
	// if there is non-nil error ->
	// IGNORE other returns of a function OR partial results
	// -> documentation of return values essential for disambiguation

}
