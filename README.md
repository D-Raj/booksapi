This is a minimalist Web service built with the Go Programming Language.

#WORK IN PROGRESS

How to run: 
Navigate to your GOPATH folder
go get "github.com/dharnnie/booksapi"
cd booksapi
go get ./... (to install dependencies)
go run main.go

localhost:5000 {GET} (Prints hello world in console)
localhost:5000/books {GET}(Returns an Array if books)
localhost:5000/books/{id} --- {GET}(Returns a single book by ID) ID could be from 1 - 4 [at this time]
localhost:5000/books/{id} --- {POST} (Adds specified request body to the list and returns an updated list)