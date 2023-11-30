package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	config "github.com/AliiAhmadi/ShareCode/config"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	config := new(config.Config)

	flag.IntVar(&config.Port, "port", 4000, "Listen port")
	flag.StringVar(&config.Address, "addr", "127.0.0.1", "HTTP network address")

	app := &application{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}

	flag.Parse()

	server := &http.Server{
		Addr:     config.Get(),
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("Starting server on %s", config.Get())
	err := server.ListenAndServe()
	if err != nil {
		app.errorLog.Fatal(err)
	}
}
