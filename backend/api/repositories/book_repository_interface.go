package repositories

import (
	Handlers "gitub.com/RomainC75/biblio/api/handlers"
	Models "gitub.com/RomainC75/biblio/data/models"
)

type BookRepositoryInterface interface {
	CreateBook(bookInfos Handlers.SearchInApisResponse) (Models.Book, error)
	
	FirstOrCreateEditor(editorName string) Models.Editor 
	FirstOrCreateLinks(newLinks []string, bookId uint) []Models.Link
	FirstOrCreateLanguages(book Models.Book, newLanguages []string) []Models.Language
	FirstOrCreateAuthors(book Models.Book, newAuthors []string) []Models.Author
	FirstOrCreateGenres(book Models.Book, newGenres []string) []Models.Genre
	
	FindBookByISBN(isbn string) (Models.Book, error)

	FindByISBN(isbn string) (Models.Book, error)
	FindById(id string) (Models.Book, error)
	FindByUserID(userId string) []Models.Book
	DeleteBookById(userId string) (Models.Book, error)
	UpdateBookById(book Models.Book) (Models.Book, error)

	UpdateAuthors(authors []Models.Author) ([]Models.Author, error)


}
