This is a minimalist Web service built with the Go Programming Language.

#WORK IN PROGRESS

How to run: 
Navigate to your GOPATH folder
go get "github.com/dharnnie/booksapi"
cd booksapi
go get ./... (to install dependencies)
go run main.go

localhost:5000 (Prints hello world in console)
localhost:5000/books (Returns an Array if books)
localhost:5000/books/{id} (Returns a single book by ID) ID could be from 1 - 4 [at this time]