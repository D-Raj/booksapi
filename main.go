package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dharnnie/booksapi/handlers"
	"github.com/gorilla/mux"
)

func main() {
	server()
}

func server() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	myMux := mux.NewRouter()
	// This is the first route
	myMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		println("Hello World!")
	})
	myMux.HandleFunc("/books", handlers.AllBooks).Methods("GET")

	http.Handle("/", myMux)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}
