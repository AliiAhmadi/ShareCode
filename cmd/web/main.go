package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	address := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on " + *address)
	err := http.ListenAndServe(*address, mux)
	if err != nil {
		log.Fatal(err)
	}
}
