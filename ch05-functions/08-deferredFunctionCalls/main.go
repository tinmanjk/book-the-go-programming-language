package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("Deferred Function Calls")

	fmt.Println("\nRelease resources use")
	url := `http://gopl.io`
	if err := printTitle(url); err != nil {
		fmt.Fprintf(os.Stderr, "url: %v\n", err) // should be no error here
	}
	url = `https://golang.org/doc/gopher/frontpage.png`
	if err := printTitle(url); err != nil {
		fmt.Fprintf(os.Stderr, "url: %v\n", err)
		// title: https://golang.org/doc/gopher/frontpage.png has type image/png, not text/html
	}

	fmt.Println("\nMultiple defers - LIFO - stack")
	func() {
		// **reason** imo seems to be that releasing of resources order matters
		for i := 0; i < 5; i++ {
			defer fmt.Printf("%d ", i) // we are registering 5 defers capturing i at the time of creation
		}
	}() // 4 3 2 1 0
	fmt.Println()

	fmt.Println("\nEvaluation of arguments vs execution") // arguments immediatley, function body at end
	func() {
		fmt.Println(i)                 // 0
		defer fmt.Println(increment()) // increment() is argument expression -> will be evaluated HERE
		fmt.Println(i)                 // 1 not 0
		// 1 here as well
	}()

	fmt.Println("\nTiming slow functions - defer with functions that return closures")
	bigSlowOperation()

	fmt.Println("\nDefer with closure to access and modify return variables")
	_ = double(4)
	// Output:
	// "double(4) = 8"
	fmt.Println(triple(4)) // "12"

	fmt.Println("\nDefer statement in a loop for releasing resources - separate function")
	func() error {
		filenames := []string{}
		for _, filename := range filenames {
			// if we put defers in the loop body
			// they will ALL happen at once at the end of the function/loop
			// better to refactor file open/close logic in separate function
			if err := doFile(filename); err != nil {
				return err
			}
		}
		return nil
	}()

	fmt.Println("\nSpecial case of NOT using defer with File Closing")
	func() (filename string, n int64, err error) {
		f, err := os.Create("local")
		if err != nil {
			return "", 0, err
		}

		// why not ? TODO clear
		useDefer := false
		defer func(useMe bool) {
			if closeErr := f.Close(); err == nil { // only if other errors are nil should we take the close err
				err = closeErr
			}
		}(useDefer)

		n, err = io.Copy(f, f) // original resp.Body for second f
		// NFS File System -> errors are postponed until file is closed

		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil { // only if copyError is nil
			err = closeErr
		}
		return "local", n, err // err can be nil, or the CopyError or the Close Error
	}()

}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...process f...
	return nil
}

// pattern for printing arguments and results of a function
func double(x int) (result int) {
	// nothing special about the return variable
	// it's the closure that captures the variables BUT does NOT evaluate them
	// or maybe **evaluation = capture of address** of outside variable scope
	defer func() {
		// the process of capturing enclosing variables seems akin to getting the arguments
		// it is NOT part of the body of the function to be executed
		// **implicit function value evaluation -> capturing enclosing variables**
		capturedResult := &result
		capturedX := &x
		// body
		fmt.Printf("double(%d) = %d\n", *capturedX, *capturedResult)

		// original version
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	return x + x
}

// change of results
func triple(x int) (result int) {
	defer func() { result += x }()
	// this is the same as
	// result = double(x)
	// return -> which calls the defer BEFORE ACTUALLY returning
	// so return -> call the chain of deferred functions -> THEN return
	return double(x)
}

var i int

func increment() int {
	i++
	return i
}

func printTitle(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // immediately after the resource -> resp is acquired

	// Check Content-Type is HTML (e.g., "text/html; charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func bigSlowOperation() {
	// arguments are evaluated immediately
	// here the actual function to be deferred is the one returned by trace
	// in this sense trace is evaluated/executed first which returns the closure
	// **on-entry** -> the func generator, **on-exit** the generated fun
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(2 * time.Second) // simulate slow operation by sleeping
}

// nice use of closure
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	// here we have start initialized to the expression time.Now() AND HAVE NOT CHANGED IT
	// maybe **closure capture** of the address is the evaluation, not the VALUE
	return func() { log.Printf("exit  %s (%s)", msg, time.Since(start)) }
}
