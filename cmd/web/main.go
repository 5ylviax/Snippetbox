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

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home) // broweser sends a request to the server, it maches `/` to the home handler the home function runs and sends back the response
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate) // route only responds to HTTP request which use the POST method

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting serer %s", *addr)
	err := srv.ListenAndServe()
	// Port 0 - 1023 are restricted and typically can only be used by services which have root privileges

	errorLog.Fatal(err)
}
