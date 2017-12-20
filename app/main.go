package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dharnnie/booksapi/app/auth"
	"github.com/dharnnie/booksapi/app/handlers"
	"github.com/gorilla/mux"
)

func main() {
	server()
}

func server() {
	// check for environment our service is running on
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	myMux := mux.NewRouter()
	// This is the first route, creates and returns token to client
	myMux.HandleFunc("/", auth.CreateToken).Methods("POST")
	// route that gets some of the books I have read
	myMux.HandleFunc("/books", handlers.AllBooks).Methods("GET")
	myMux.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	myMux.HandleFunc("/books/{id}", handlers.AddBook).Methods("POST")
	myMux.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// bind my multiplexer (from gorilla to app)
	http.Handle("/", myMux)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}

// {
// 	"id":"5",
// 	"name": "Enchantment  the Art of changing  hearts",
// 	"author": "Guy Kawasaki",
// 	"publisher": "Not specified"
// }
