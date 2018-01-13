package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/dharnnie/booksapi/app/types"
)

// I genereted this token before hand with the CreateToken function
var token = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6IiIsInVzZXJuYW1lIjoiIn0.QzJ1vY-cvf3uHHadiE5IZLl40_kd-zgt1jQS6LOPdN77dtFa8gnSQe_3nNcvkJj-mzHGnbtINSd2q5nwBxU4GA"

func TestCreateToken(t *testing.T) {
	params := url.Values{}
	params.Set("username", "poo")
	params.Set("password", "poo")
	reqPayload := strings.NewReader(params.Encode())
	expected := token
	req, err := http.NewRequest("POST", "localhost:5000", reqPayload)
	if err != nil {
		t.Fatalf("Error connecting to server: %f", err)
	}
	rec := httptest.NewRecorder()
	CreateToken(rec, req)

	res := rec.Result()
	res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Request status not OK: got %v", res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading ioutil: %v", err)
	}
	var d types.JwtToken
	json.Unmarshal(data, &d)
	actual := d
	fmt.Println(actual.Token)
	if actual.Token != expected {
		t.Errorf("Got %v, but expected %v", actual, expected)
	}

}
