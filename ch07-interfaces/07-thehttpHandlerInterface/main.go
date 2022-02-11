package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("The HTTP Handler Interface")

	db := database{"shoes": 50, "socks": 5}

	// func ListenAndServe(address string, h Handler) error
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d) // f -> decimal point, .2 precision
}

// net/http

// type Handler interface {
// 	ServeHTTP(w ResponseWriter, r *Request)
// }

type database map[string]dollars

// logic for EVERY http request
func (db database) ServeHTTP(w http.ResponseWriter, // type response struct
	req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
