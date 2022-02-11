package main

import (
	"fmt"
	"log"
	"net/http"
)

const useMux = true

var dbHandlesListOnly bool

func main() {
	fmt.Println("The HTTP Handler Interface")

	db := database{"shoes": 50, "socks": 5}

	// ServeMux -> aggregates different Handlers into one
	// simplifies routing (association urls and handlers)
	mux := http.NewServeMux() // request multiplexer -> http.Handler
	mux.Handle("/list", db)   // normal handler (adapted, db.list originally in book)

	// func type adapter
	// 	type HandlerFunc func(ResponseWriter, *Request)

	// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	// 	f(w, r)
	// }
	// **let's a function value satisfy an interface**
	mux.Handle("/price", http.HandlerFunc(db.price)) // **conversion** of method value
	// mux.HandleFunc("/price", db.price) -> CONVENIENCE METHOD

	// http.Handler satisfaction by serve mux
	//func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	// ...
	// 	h, _ := mux.Handler(r)
	// 	h.ServeHTTP(w, r)
	// }

	// func ListenAndServe(address string, h Handler) error
	if useMux {
		dbHandlesListOnly = true
		log.Fatal(http.ListenAndServe("localhost:8000", mux))
	} else {
		dbHandlesListOnly = false
		log.Fatal(http.ListenAndServe("localhost:8000", db))

	}
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

func (db database) ServeHTTP(w http.ResponseWriter, // type response struct
	req *http.Request) {

	if dbHandlesListOnly {
		db.list(w, req)
		return
	}

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

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter,
	req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
