package types

// Book type is used across board | Represents datatype for a single book
type Book struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Author    string `json:"author,omitempty"`
	Genre     string `json:"genre,omitempty"`
	Publisher string `json:"publisher,omitempty"`
}
