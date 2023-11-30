package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	config "github.com/AliiAhmadi/ShareCode/config"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	config := new(config.Config)

	flag.IntVar(&config.Port, "port", 4000, "Listen port")
	flag.StringVar(&config.Address, "addr", "127.0.0.1", "HTTP network address")
	flag.StringVar(&config.DSN, "dsn", "sharecode:pass@/sharecode?parseTime=true", "Mysql database dsn for connection")

	app := &application{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}
	db, err := openDB(config.DSN)
	if err != nil {
		app.errorLog.Fatal(err)
	}
	defer db.Close()

	flag.Parse()

	server := &http.Server{
		Addr:     config.Get(),
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("Starting server on %s", config.Get())
	err = server.ListenAndServe()
	if err != nil {
		app.errorLog.Fatal(err)
	}
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
