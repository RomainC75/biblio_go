package repositories

import (
	ApisHandler "gitub.com/RomainC75/biblio/internal/modules/apis/handler"
	bookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
)

type BookRepositoryInterface interface {
	CreateBook(bookInfos ApisHandler.SearchInApisResponse) (bookModel.Book, error)
	
	FirstOrCreateEditor(editorName string) bookModel.Editor 
	FirstOrCreateLinks(newLinks []string, bookId uint) []bookModel.Link
	FirstOrCreateLanguages(book bookModel.Book, newLanguages []string) []bookModel.Language
	FirstOrCreateAuthors(book bookModel.Book, newAuthors []string) []bookModel.Author
	FirstOrCreateGenres(book bookModel.Book, newGenres []string) []bookModel.Genre
	
	FindBookByISBN(isbn string) (bookModel.Book, error)

	FindByISBN(isbn string) (bookModel.Book, error)
	FindById(id string) (bookModel.Book, error)
	FindByUserID(userId string) []bookModel.Book
	DeleteBookById(userId string) (bookModel.Book, error)
	UpdateBookById(book bookModel.Book) (bookModel.Book, error)

	UpdateAuthors(authors []bookModel.Author) ([]bookModel.Author, error)


}
