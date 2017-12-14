package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAllBooks(t *testing.T) {
	expected := books
	req, err := http.NewRequest("GET", "localhost:5000/books", nil)
	if err != nil {
		t.Fatal("Error connecting to server: %v", err)
	}
	rec := httptest.NewRecorder()
	AllBooks(rec, req)

	res := rec.Result()
	res.Body.Close() // Not compulsory(because this is a test), but a great practice
	if res.StatusCode != http.StatusOK {
		t.Errorf("Request status not OK: got %v", res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	actual, _ := json.Marshal(data)
	if err != nil {
		t.Errorf("Error reading ioutil: %v", err)
	}

	if string(actual) != expected {
		t.Errorf("Expected %v: got: %v", expected, actual)
	}
}
