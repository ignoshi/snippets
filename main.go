package main

import (
	"log"
	"net/http"
	"os"
	"time"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ignoshi/snippets/snippets"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/snippets", snippets.ListSnippets).Methods("GET")
	r.HandleFunc("/api/snippets", snippets.CreateSnippet).Methods("POST")
	r.HandleFunc("/api/snipepts/{id}", snippets.GetSnippet).Methods("GET")

	lr := gh.LoggingHandler(os.Stdout, r)

	srv := &http.Server{
		Handler:      lr,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started, Listening on :8000")
	log.Fatal(srv.ListenAndServe())
}
