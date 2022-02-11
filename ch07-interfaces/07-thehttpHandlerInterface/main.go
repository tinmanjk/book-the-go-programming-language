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
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
		// alternatively
		// msg := fmt.Sprintf("no such page: %s\n", req.URL)
		// http.Error(w, msg, http.StatusNotFound) // 404
	}
}
