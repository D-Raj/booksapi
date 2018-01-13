package types

// Book type is used across board | Represents datatype for a single book
type Book struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Author    string `json:"author,omitempty"`
	Genre     string `json:"genre,omitempty"`
	Publisher string `json:"publisher,omitempty"`
}

// User : user login details is decoded into this struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JwtToken is holds token string that goes back to the client
type JwtToken struct {
	Token string `json:"token"`
}

// Exception is used to hold error messages in authentication
type Exception struct {
	Message string `json:"message"`
}
