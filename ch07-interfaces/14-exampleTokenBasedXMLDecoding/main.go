package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Example: Token based XML Decoding")
	// encoding/xml -> uses empty interface Token - no static type checking
	// -> see discussion here:
	// https://stackoverflow.com/questions/21553398/best-practice-for-unions-in-go
	// -> non-empty interfaces but e.g. ImplementsInterfaceMethod and then no impl hack
	// https://github.com/golang/go/issues/19412 -> add unions/sum types - rejected
	// however generics/constraints may alleviate the issue with compile time checks
	// https://stackoverflow.com/questions/69772512/go-generics-unions
	// https://github.com/golang/go/issues/45380 -> in the future maybe fully with type switches
	// -> workaround at the moment -> investigate generics further

	xmlFile, err := os.Open("example.htm")
	// http://www.w3.org/TR/2006/REC-xml11-20060816
	if err != nil {
		log.Fatal(err)
	}
	dec := xml.NewDecoder(xmlFile)
	var stack []string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		// here we have to read the comments from Token interface
		// which types it unionizes
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			args := []string{"div", "div", "h2"} // instead of os.Args[1:]
			if containsAll(stack, args) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
