package db

import "github.com/dharnnie/booksapi/app/types"

var booklist []types.Book

// Books gets and returns an array of books from database
// Here there is a mock database
func Books() []types.Book {
	gobook := types.Book{ID: "1", Name: "GoBook", Author: "Caleb Doxsey", Genre: "Programming", Publisher: "Not specified"}
	elon := types.Book{ID: "2", Name: "Elon Musk | Tesla | SpaceX and the quest for a fantastic future", Author: "Ashley Vance", Genre: "Biography", Publisher: "Not specified"}
	strengthInGod := types.Book{ID: "3", Name: "Finding strength in the word of God", Author: "Kenneth Hagin", Genre: "Spiritual", Publisher: "Not specified"}
	java := types.Book{ID: "4", Name: "Java for dummies", Author: "Barry Burd", Genre: "Programming", Publisher: "John Wiley and Sons"}
	aod := types.Book{ID: "5", Name: "Acres of Diamond", Author: "Russell .H Conwell", Genre: "Inspiration", Publisher: "Not Specified"}
	tagr := types.Book{ID: "6", Name: "Think and grow rich", Author: "Napoleon Hill", Genre: "Finance|Wealth|Inspiration", Publisher: "Not specified"}
	wwla := types.Book{ID: "7", Name: "We will lead Africa", Author: "Gilpin-Jackson|Sarah|Judith", Genre: "Programming", Publisher: "John Wiley and Sons"}
	// dump all the books in a single array
	proReact := types.Book{ID: "8", Name: "Pro React", Author: "Cassio De Soussa", Genre: "Programming", Publisher: "Apress"}
	booklist = append(booklist, gobook, elon, strengthInGod, java, aod, tagr, wwla, proReact)
	return booklist
}
