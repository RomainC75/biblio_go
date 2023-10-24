package responses

import (
	"github.com/google/uuid"
	bookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
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

func ToBook(book bookModel.Book) Book {
	return Book{
		ID:    book.ID,
		Title: book.Title,
		ISBN:  book.ISNB,
		// Authors: book.Authors,
	}
}
