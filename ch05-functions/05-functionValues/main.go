package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("Function values")

	f := square
	fmt.Println(f(3)) // "9"
	f = negative
	fmt.Println(f(3))     // "-3"
	fmt.Printf("%T\n", f) // "func(int) int"
	// f = product           // compile error: can't assign func(int, int) int to func(int) int

	// zero value - nil, only comparable with nil
	fmt.Println("\nZero value")
	{
		var f func(int) int
		fmt.Println(f == nil) // true
		// f(3) // panic: call of nil function
	}

	fmt.Println("\nParameterize over behavior")
	// strings.Map -> applies a function over each CHARACTER
	fmt.Println(strings.Map(add1, "HAL-9000")) //"IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    //"Benjy"

	fmt.Println("\nUsing functions as arguments to apply different actions")
	url := `https://pkg.go.dev/golang.org/x/net/html?tab=versions`
	outline(url)

}

func square(n int) int   { return n * n }
func negative(n int) int { return -n }

//lint:ignore U1000 ....
func product(m, n int) int { return m * n }

func add1(r rune) rune { return r + 1 }

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, printOpeningElementTag, printClosingElementTag)
	return nil
}

// separation of tree traversal from actions to be applied on nodes
// much more flexible
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

// to track the state of indendation and trere traversal
var depth int

func printOpeningElementTag(n *html.Node) {
	if n.Type == html.ElementNode {
		// * adverb -> number of spaces to pad = depth*2
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func printClosingElementTag(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
