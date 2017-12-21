This is a minimalist Web service built with the Go Programming Language.

#WORK IN PROGRESS

How to run: 
<hr>
Navigate to your GOPATH folder
<hr>
git clone https://github.com/dharnnie/booksapi.git
cd booksapi/app
<hr>
go get ./... (to install dependencies)
<hr>
go run main.go

You can easily test via postman
<hr>
localhost:5000 {GET} (Prints hello world in console)

Requires a request body to create a bearer token.
sample body: 
{
    "username":"poo",
    "password": "poo"
}
Meanwhile, I have created a token with that sample body and you can access other routes with it without having to request for your own bearer token

Bearer token : eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6IiIsInVzZXJuYW1lIjoiIn0.QzJ1vY-cvf3uHHadiE5IZLl40_kd-zgt1jQS6LOPdN77dtFa8gnSQe_3nNcvkJj-mzHGnbtINSd2q5nwBxU4GA
<hr>
localhost:5000/books {GET}(Returns an Array if books) 
- Requires the bearer token
<hr>
localhost:5000/books/{id} --- {GET}(Returns a single book by ID) ID could be from 1 - 8 [at this time]
- Requires the bearer token
<hr>
localhost:5000/books/{id} --- {POST} (Adds specified request body to the list and returns an updated list)
- Requires the bearer token

<hr>
localhost:5000/books/{id} --- {DELETE} (Deletes specified a single book with ID from the list)
- Requires the bearer token
