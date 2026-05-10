package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home) // broweser sends a request to the server, it maches `/` to the home handler the home function runs and sends back the response
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate) // route only responds to HTTP request which use the POST method

	log.Println("Starting serer on :4000")
	err := http.ListenAndServe(":4000", mux) // host:port --> localhost::4000

	log.Fatal(err)
}
