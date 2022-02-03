package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

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

	fmt.Println("\n5.4.1 Error-Handling Strategies")

	// 1. Propagation (simple and with context prefixing)
	fmt.Println("\n1. Propagation")
	_, err = propagateErrorFromHttpGet("")
	fmt.Println(err) // Get "": unsupported protocol scheme ""
	_, err = contextPrefixErrorFromHttpGet("")
	fmt.Println(err) // context prefix: Get "": unsupported protocol scheme ""

	// 2. **Retry** - limited time/number of trials before completely giving up
	fmt.Println("\n2. Retry")
	if err := WaitForServer("", 6); err != nil {
		fmt.Println(err)
	}

	// 3. Safely **ignore** the error (fifth in text)
	fmt.Println("\n3. Ignore")
	// will create directory in OS Temp [rewritten by me anti-pattern normally]
	if dir, err := ioutil.TempDir("", "ignore"); err == nil {
		// ...use temp dir...
		os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically
	}

	// 4. **Log** the error & continue (fourth in text)
	fmt.Println("\n4. Log and Continue")
	if err := ping(); err != nil {
		log.Printf("ping failed: %v; networking	disabled", err) // **appends a new line** even without \n
		// alternatively
		// fmt.Fprintf(os.Stderr, "ping failed: %v;networking disabled\n", err)
	}

	// 5. Progress impossible -> log error and stop the program (third in text)
	// reserved for main package, library should propagate
	fmt.Println("\n5. Log and Terminate the Program")
	if err := WaitForServer("", 1); err != nil {
		// fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		// os.Exit(1)

		// alternative
		// log.SetFlags(0b11)      // default is 0b11 -> to not display date/time 0b00
		log.SetFlags(0b10)      // time only
		log.SetPrefix("wait: ") // this can be combined with the date/time or replace it
		log.Fatalf("Site is down: %v\n", err)
		// wait: 07:36:12 Site is down: server  failed to respond after 1s
		// exit status 1
	}

	fmt.Println("\n5.4.2 End of File (EOF)")
	// apparently we can have named errors
	// var EOF = errors.New("EOF") ->
	// how do we compare interface types?
	// https://golangbyexample.com/comparing-error-go/
	// **changes in error handling** as of 1.13
	// https://stackoverflow.com/questions/39121172/how-to-compare-go-errors
	// errors as constants -> interesting
	// https://dave.cheney.net/2016/04/07/constant-errors
	// to be continued in **7.11**
	// if err == io.EOF {
	// 	break // finished reading
	// }
	// if err != nil {
	// 	return fmt.Errorf("read failed: %v", err)
	// }

}

func propagateErrorFromHttpGet(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func contextPrefixErrorFromHttpGet(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		// fmt.Errorf -> uses fmt.Sprintf to format and return a new error VALUE
		// **prefix context** to underlying error - that the called f doesn't have
		// -> building clear **causal chain** root problem - overal failure
		// e.g.
		// genesis: crashed: no parachute: G-switch failed: bad relay orientation
		// due to chaining - no new-line no capital letters
		return nil, fmt.Errorf("context prefix: %v", err)
	}
	return resp, nil
}

func ping() error {
	return fmt.Errorf("no network")
}

// WaitForServer attempts to contact the server of a URL.
// It tries for 15 seconds using exponential back-off. // 1-2-4-8 second intervals
// It reports an error if all attempts fail.
func WaitForServer(url string, seconds int) error {
	timeout := time.Second * time.Duration(seconds)
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ { // tries used for exponential
		_, err := http.Head(url)
		if err != nil {
			log.Printf("server not responding (%s); retrying...", err)
			// blocking
			time.Sleep(time.Second << uint(tries)) // exponential back-off * 2
			continue
		}
		return nil // success
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
