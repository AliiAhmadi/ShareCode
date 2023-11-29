package main

import (
	"flag"
	"log"
	"net/http"

	config "github.com/AliiAhmadi/ShareCode/config"
)

func main() {
	config := new(config.Config)

	flag.IntVar(&config.Port, "port", 4000, "Listen port")
	flag.StringVar(&config.Address, "addr", "127.0.0.1", "HTTP network address")

	flag.Parse()
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on " + config.Get())
	err := http.ListenAndServe(config.Get(), mux)
	if err != nil {
		log.Fatal(err)
	}
}
