package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	config "github.com/AliiAhmadi/ShareCode/config"
)

func main() {
	config := new(config.Config)

	flag.IntVar(&config.Port, "port", 4000, "Listen port")
	flag.StringVar(&config.Address, "addr", "127.0.0.1", "HTTP network address")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	flag.Parse()
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := &http.Server{
		Addr:     config.Get(),
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", config.Get())
	err := server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
