package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/tinmanjk/tgpl/ch05-functions/06-anonymousFunctions/links"
)

func main() {

	fmt.Println("Anonymous functions")
	// the value of a function literal expression
	// helps us defined functions at the point of use
	mapped := strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
	fmt.Println(mapped) // IBM.:111

	fmt.Println("\nClosures")
	fmt.Println(squares()()) // "1"
	fmt.Println(squares()()) // "1" // separate function value

	fmt.Println()
	f := squares()
	// ** function values CAN HAVE state **
	// that's why they are not comparable and are reference types
	// called closures
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4" // has access to the state of x
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"

	fmt.Println("\nTopological Sorting")
	// Directed Acyclic Graph
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

	fmt.Println("\nExtract with absolute link for crawling")
	url := `https://pkg.go.dev/golang.org/x/net/html?tab=versions`
	linkss, err := links.Extract(url)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, l := range linkss {
		fmt.Println(l)
	}

	fmt.Println("\nBread-First Search Crawler")
	workListQueue := []string{`https://pkg.go.dev/golang.org/x/net/html?tab=versions`}
	breadthFirst(workListQueue, printURLandExtractLinks, 2, 3)

	// 5.6.1 Caveat: Capturing Iteration Variables
	fmt.Println("5.6.1 Caveat: Capturing Iteration Variables")
	var closureWrongIterationCapture []func()
	var closureCorectIterationCapture []func()

	stringsSlice := []string{"one", "two", "three"}
	for _, s := range stringsSlice {
		closureWrongIterationCapture = append(closureWrongIterationCapture,
			func() {
				fmt.Println(s)
			})
		s := s // can be the same name -> shadowing but different memory location to capture
		closureCorectIterationCapture = append(closureCorectIterationCapture,
			func() {
				fmt.Println(s)
			})
	}

	for i := 0; i < 3; i++ {
		closureWrongIterationCapture[i]()  // always "three" -> last iteration value
		closureCorectIterationCapture[i]() // always -> one two three
		fmt.Println()
	}
	// issue to be aware of when using go/defer

}

// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int {
		// has access to the state of x
		x++
		return x * x
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string) // need for recursion -> otherwise undeclared name
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item]) //DFS
				order = append(order, item)
			}
		}
	}

	// why do we need to have them sorted alphabetically ?
	// deterministic result output?
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	visitAll(keys)
	return order
}

// "crawl" in book
func printURLandExtractLinks(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// breadthFirst calls getChildNodes (f in book) for each item in the worklistQueue (queue mine)
// Any items returned by getChildNodes are added to the worklistQueue
// getChildNodes (f in book) is called at most once for each item.
func breadthFirst(worklistQueue []string, getChildNodes func(item string) []string,
	depth int, maxChildNodes int) {
	seen := make(map[string]bool)
	currentDepth := 0
	for len(worklistQueue) > 0 && currentDepth < depth {
		fmt.Printf("\nDepth: %d\n", currentDepth)
		currentDepthItems := worklistQueue
		worklistQueue = nil // next level to be traversed
		for _, item := range currentDepthItems {
			if !seen[item] {
				seen[item] = true
				// maxChildNodes -> can have duplicates -> so != maxBreadth
				worklistQueue = append(worklistQueue, getChildNodes(item)[:maxChildNodes]...)
			}
		}
		currentDepth++
	}
}
