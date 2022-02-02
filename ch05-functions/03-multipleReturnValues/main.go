package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("Multiple Return Values")

	url := `https://pkg.go.dev/golang.org/x/net/html?tab=versions`
	links, err := findLinksLog(url)
	// links, _ := findLinks(url) // errors ignored
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
	}
	for _, link := range links {
		fmt.Println(link)
	}

	// multiple return values as arguments to another function
	// log.Println(findLinks(url))
	// links, err := findLinks(url)
	// log.Println(links, err)

}

// well chosen names for return values -> especially important if the same type
// func Size(rect image.Rectangle) (width, height int)
// func Split(path string) (dir, file string)
// func HourMinSec(t time.Time) (hour, minute, second int)

func findLinksLog(url string) ([]string, error) {
	log.Printf("findLinks %s", url)
	return findLinks(url) // direct return without first assigning to local variables
}

// multiple returns -> result and error
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		// unchanged error is propagated
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		// need to release resources - open files and network connections
		resp.Body.Close()
		// augmenting the error with context
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	// need to release resources - open files and network connections
	resp.Body.Close()
	if err != nil {
		// augmenting the error with context
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
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

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		// bare return
		// return words, images, err
		// not obvious that it is return 0,0,err
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		// bare return
		// return words, images, err
		return
	}
	words, images = countWordsAndImages(doc)
	// bare return
	// return words, images, nil // not clear
	return
}

func countWordsAndImages(n *html.Node) (int, int) {
	// bare does NOT work with UNNAMED result variables
	return 0, 0
}
