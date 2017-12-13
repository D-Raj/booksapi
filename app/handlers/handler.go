package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dharnnie/booksapi/types"
	"github.com/gorilla/mux"
)

// Array of books
var books []types.Book

// dummy function that simulates getting/initializing books from database
func getBooks() {
	gobook := types.Book{ID: "1", Name: "GoBook", Author: "Caleb Doxsey", Genre: "Programming", Publisher: "Not specified"}
	elon := types.Book{ID: "2", Name: "Elon Musk | Tesla | SpaceX and the quest for a fantastic future", Author: "Ashley Vance", Genre: "Biography", Publisher: "Not specified"}
	strengthInGod := types.Book{ID: "3", Name: "Finding strength in the word of God", Author: "Kenneth Hagin", Genre: "Spiritual", Publisher: "Not specified"}
	java := types.Book{ID: "4", Name: "Java for Dummies", Author: "Barry Burd", Genre: "Programming", Publisher: "John Wiley and Sons"}
	// dump all the books in a single array
	books = append(books, gobook, elon, strengthInGod, java)
}

// AllBooks function that returns all books
func AllBooks(w http.ResponseWriter, r *http.Request) {
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
	getBooks()
	params := mux.Vars(r)
	var book types.Book
	_ = json.NewDecoder(r.Body).Decode(&book) //Decode body into a book struct and Ignore errors
	book.ID = params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
	books = nil // empty the book variable
}

// DeleteBook deletes a single book. This action is not persistent. It returns a list of books excluding deleted book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
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
