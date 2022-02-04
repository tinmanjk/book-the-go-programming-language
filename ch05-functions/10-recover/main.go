package main

import (
	"fmt"
)

func main() {
	fmt.Println("Recover")
	// clean up the mess before quitting
	notInDefer := recover() // no effect just returns nil
	fmt.Println(notInDefer)

	fmt.Println("\nTurn panics into errors by recovering in defers")
	fmt.Println(recoverFromPanic())
	fmt.Println("Continue execution") // no panic state here

	// as a general rule -> do not recover from another package's panic
	// e.g. net/http server recovers from panics in the callbacks - convenient
	// but it continual service can have issues

	fmt.Println("\nConverting specific panics into errors") // somewhat of anti-pattern
	err := specialPanicValueToCatch(true)
	fmt.Println("Will not print if unexpected panic: ", err)
}

func specialPanicValueToCatch(causeExpectedPanic bool) (err error) {
	// type definition inside a function ??? - TODO: investigate
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
		// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("expected panic")
		default:
			panic(p) // unexpected panic; carry on	panicking
			// still adds [recovered] for the first panic to note the **rethrow**
			// panic: unexpected [recovered]
			// panic: unexpected
		}
	}()

	if causeExpectedPanic {
		panic(bailout{})
	} else {
		panic("unexpected")
	}
}

func recoverFromPanic() (err error) {
	defer func() {
		if p := recover(); p != nil {
			// optionally runtime.Stack to return stacktrace
			err = fmt.Errorf("panic value: %v", p)
			// cannot return here
		}
	}()

	// this can be a third party function call
	// which **COULD** lead to corrupted state there
	// so you should not really recover because you **cannot reason about safety**
	causePanic()
	return nil
}

func causePanic() {
	panic("bad stuff happening")
}
