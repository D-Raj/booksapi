This is a minimalist Web service built with the Go Programming Language.

#WORK IN PROGRESS

How to run: 
<hr>
Navigate to your GOPATH folder
<hr>
go get "github.com/dharnnie/booksapi"
cd booksapi
<hr>
go get ./... (to install dependencies)
<hr>
go run main.go

<hr>
localhost:5000 {GET} (Prints hello world in console)
<hr>
localhost:5000/books {GET}(Returns an Array if books)
<hr>
localhost:5000/books/{id} --- {GET}(Returns a single book by ID) ID could be from 1 - 4 [at this time]
<hr>
localhost:5000/books/{id} --- {POST} (Adds specified request body to the list and returns an updated list)
<hr>
localhost:5000/books/{id} --- {DELETE} (Deletes specified a single book with ID from the list)