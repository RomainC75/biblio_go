package responses

import (
	"github.com/google/uuid"
)

type Book struct {
	ID      uuid.UUID
	Title   string
	ISBN    string
	Authors []Author
}

type Author struct {
	ID   uuid.UUID
	Name string
}

// func ToBook(book bookModel.Book) Book {
// 	var authors []Author
// 	for _, author := range book.Authors {
// 		newAuthor := Author{
// 			ID:   author.ID,
// 			Name: author.Name,
// 		}
// 		authors = append(authors, newAuthor)
// 	}
// 	return Book{
// 		ID:      book.ID,
// 		Title:   book.Title,
// 		ISBN:    book.ISBN,
// 		Authors: authors,
// 	}
// }

