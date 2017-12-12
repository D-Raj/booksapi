package handlers

import (
	"encoding/json"
	"net/http"
)

// Structure of a single book
type Book struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Author    string `json:"author,omitempty"`
	Genre     string `json:"genre,omitempty"`
	Publisher string `json:"publisher,omitempty"`
}

// Array of books
var books []Book

// dummy function that simulates getting/initializing books from database
func getBooks() {
	gobook := Book{ID: 1, Name: "GoBook", Author: "Caleb Doxsey", Genre: "Programming", Publisher: "Not specified"}
	elon := Book{ID: 2, Name: "Elon Musk", Author: "Ashley Vance", Genre: "Biography", Publisher: "Not specified"}
	strengthInGod := Book{ID: 3, Name: "Finding strength in the word of God", Author: "Kenneth Hagin", Genre: "Spiritual", Publisher: "Not specified"}
	java := Book{ID: 4, Name: "Java for Dummies", Author: "Barry Burd", Genre: "Programming", Publisher: "John Wiley and Sons"}
	// dump all the books in a single array
	books = append(books, gobook, elon, strengthInGod, java)
}

// function that returns all books
func AllBooks(w http.ResponseWriter, r *http.Request) {
	//get all books from magical database
	getBooks()
	// encode and return books
	json.NewEncoder(w).Encode(books)
	// clear the array so we don't overpopulate the response on multiple calls
	// remove the line below, make the call multiple times and see the difference
	books = nil
}
