package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dharnnie/booksapi/app/types"
)

// TestAllBooks tests for a handler function
func TestAllBooks(t *testing.T) {

	//getBooks()
	expected := books // books from the original function | predefined struct
	// json.Unmarshal(books, &expected)
	req, err := http.NewRequest("GET", "localhost:5000/books", nil)
	if err != nil {
		t.Fatalf("Error connecting to server: %v", err)
	}
	// create a record from the http server
	rec := httptest.NewRecorder()
	AllBooks(rec, req)

	res := rec.Result()
	res.Body.Close() // Not compulsory(because this is a test), but a great practice
	if res.StatusCode != http.StatusOK {
		t.Errorf("Request status not OK: got %v", res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading ioutil: %v", err)
	}
	fmt.Println("Data is :", data)
	var b []types.Book       // b is the actual
	json.Unmarshal(data, &b) // be becomes actual here
	fmt.Println(b)
	actual := b
	if actual == nil {
		t.Errorf("Got %v, but expected %v", actual, expected)
	}
}
