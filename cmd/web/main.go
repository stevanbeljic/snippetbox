package main

//up to Page 46

/**
A webapp for viewing, editing, and creating notes/snippets.
Made as a follow-along tutorial to Alex Edwards' Let's Go book
*/

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}