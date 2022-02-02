package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("Recursion")

	file, err := os.Open("htmlPackage.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	// nil seems to work as well as empty []string{} literal
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	fmt.Println("\nOutline")
	outline(nil, doc)

}

func visit(links []string, n *html.Node) []string {
	// Data = Tag Name or Text
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				// apparently append can grow a 0-size slice
				links = append(links, a.Val)
			}
		}
	}
	// only FirstChild/LastChild
	// to **traverse children** go to the child and traverse its siblings
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
		// when this returns the stack that might have grown in the call
		// will NOT be saved -> since we pass by value
		// the underlying array might have grown or a new one allocated
		// but the **len is the same as before the call**
	}
}
