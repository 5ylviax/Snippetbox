package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

// Responsibility of main() is limited to:
// i. Parsing the runtime configuration settings for the application
// ii. Establishing the dependencies for the handlers
// iii. Running the HTTP server
func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting serer %s", *addr)
	err := srv.ListenAndServe()
	// Port 0 - 1023 are restricted and typically can only be used by services which have root privileges

	errorLog.Fatal(err)
}
