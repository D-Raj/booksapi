package handlers

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Name      string `json:"name,omitempty"`
	Author    string `json:"author,omitempty"`
	Genre     string `json:"genre,omitempty"`
	Publisher string `json:"publisher,omitempty"`
}

var books []Book

func AllBooks(w http.ResponseWriter, r *http.Request) {
	gobook := Book{Name: "GoBook", Author: "Caleb Doxsey", Genre: "Programming", Publisher: "Not specified"}
	elon := Book{Name: "Elon Musk", Author: "Ashley Vance", Genre: "Biography", Publisher: "Not specified"}
	strengthInGod := Book{Name: "Finding strength in the word of God", Author: "Kenneth Hagin", Genre: "Spiritual", Publisher: "Not specified"}
	java := Book{Name: "Java for Dummies", Author: "Barry Burd", Genre: "Programming", Publisher: "John Wiley and Sons"}
	books = append(books, gobook, elon, strengthInGod, java)
	json.NewEncoder(w).Encode(books)
}
