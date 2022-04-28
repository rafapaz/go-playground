package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32
type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %f\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%f\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price_str := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(price_str, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "wrong price: %s\n", price_str)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "Added %s = %f\n", item, price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	price_str := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(price_str, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "wrong price: %s\n", price_str)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "Updated %s = %f\n", item, price)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "Deleted %s\n", item)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/add", http.HandlerFunc(db.create))
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.Handle("/delete", http.HandlerFunc(db.delete))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
