package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dharnnie/booksapi/app/db"
	"github.com/dharnnie/booksapi/app/types"
	"github.com/gorilla/mux"
)

// Array of book
var books []types.Book

// dummy function that simulates getting/initializing books from database
func getBooks() {
	books = db.Books()
}

// AllBooks function that returns all books
func AllBooks(w http.ResponseWriter, r *http.Request) {
	books = nil
	//get all books from magical database
	getBooks()
	// encode and return books
	json.NewEncoder(w).Encode(books)
	// clear the array so we don't overpopulate the response on multiple calls
	// remove the line below, make the call multiple times and see the difference
	books = nil
}

// GetBook gets a single book by Id
func GetBook(w http.ResponseWriter, r *http.Request) {
	books = nil
	getBooks()
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&types.Book{})
	books = nil
}

// AddBook adds a single book to the list. This does not persist as we do not have a db
func AddBook(w http.ResponseWriter, r *http.Request) {
	books = nil
	getBooks()
	params := mux.Vars(r)
	var book types.Book
	_ = json.NewDecoder(r.Body).Decode(&book) //Decode body into a book struct and Ignore errors
	book.ID = params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
	books = nil // empty the book variable
}

// DeleteBook deletes a single book. This action is not persistent. It returns a list of books
// excluding deleted book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	books = nil
	getBooks()
	params := mux.Vars(r)
	println(params["id"])
	for i, item := range books {
		if item.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
		json.NewEncoder(w).Encode(books)
	}
	books = nil // empty the book variable
}
