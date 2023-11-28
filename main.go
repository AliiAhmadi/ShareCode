package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	log.Println("Listen on localhost:4000")
	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
	}
}

func home(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Home page"))
}
