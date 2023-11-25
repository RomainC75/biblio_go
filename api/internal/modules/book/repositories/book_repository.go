package repositories

import (
	"errors"
	"fmt"

	ApisHandler "gitub.com/RomainC75/biblio/internal/modules/apis/handler"
	bookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
	"gitub.com/RomainC75/biblio/pkg/database"
	"gitub.com/RomainC75/biblio/pkg/utils"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func New() *BookRepository {
	return &BookRepository{
		DB: database.Connection(),
	}
}

func (BookRepository *BookRepository) CreateBook(bookInfos ApisHandler.SearchInApisResponse) (bookModel.Book, error) {
	var newBook bookModel.Book

	editor := BookRepository.FirstOrCreateEditor(bookInfos.Editor)
	bookInfos.Book.EditorRef = editor.ID
	fmt.Println("====> EDITOR IN DB : ")
	utils.PrettyDisplay(editor)

	// if err := BookRepository.DB.Create(&bookInfos.Book).Scan(&newBook).Error; err!=nil{
	// 	return bookModel.Book{}, err
	// }
	if err := BookRepository.DB.FirstOrCreate(newBook, &bookInfos.Book).Error; err!=nil{
		return bookModel.Book{}, err
	}

	links := BookRepository.FirstOrCreateLinks(bookInfos.Links, newBook.ID)
	newBook.Links = links
	languages := BookRepository.FirstOrCreateLanguages(newBook, bookInfos.Language)
	newBook.Languages = languages
	genres := BookRepository.FirstOrCreateGenres(newBook, bookInfos.Genres)
	newBook.Genres = genres

	fmt.Println("==>")
	// BookRepository.DB.Model(&book).Association("Authors")
	// BookRepository.DB.Save(&book)
	return newBook, nil
}

// func (BookRepository *BookRepository) CreateBook(book bookModel.Book) bookModel.Book {
	// var newBook bookModel.Book
// 	BookRepository.DB.Create(&book).Scan(&newBook)
// 	// BookRepository.DB.Model(&book).Association("Authors")
// 	// BookRepository.DB.Save(&book)
// 	return newBook
// }

func (BookRepository *BookRepository) FirstOrCreateEditor(editorName string) bookModel.Editor {
	var editor bookModel.Editor
	BookRepository.DB.FirstOrCreate(&editor, &bookModel.Editor{Name: editorName})
	return editor
}

func (BookRepository *BookRepository) FirstOrCreateLinks(newLinks []string, bookId uint ) []bookModel.Link {
	var links []bookModel.Link
	for _, newLink := range newLinks {
		var link bookModel.Link
		BookRepository.DB.FirstOrCreate(&link, &bookModel.Link{Url: newLink, BookRef: bookId})
		links = append(links, link)
	}
	return links
}

func (BookRepository *BookRepository) FirstOrCreateLanguages(book bookModel.Book, newLanguages []string) []bookModel.Language {
	var languages []bookModel.Language
	for _, newLanguage := range newLanguages {
		var link bookModel.Language
		BookRepository.DB.FirstOrCreate(&link, &bookModel.Language{Name: newLanguage})
		languages = append(languages, link)
	}
	// err
	BookRepository.DB.Model(&book).Association("Languages").Append(languages)
	return languages
}

func (BookRepository *BookRepository) FirstOrCreateGenres(book bookModel.Book, newGenres []string) []bookModel.Genre {
	var genres []bookModel.Genre
	for _, newGenre := range newGenres {
		var genre bookModel.Genre
		BookRepository.DB.FirstOrCreate(&genre, &bookModel.Genre{Name: newGenre})
		genres = append(genres, genre)
	}
	BookRepository.DB.Model(&book).Association("Genres").Append(genres)
	return genres
}

func (BookRepository *BookRepository) FirstOrCreateAuthors(book bookModel.Book, newAuthors []string) []bookModel.Author {
	var authors []bookModel.Author
	for _, newAuthor := range newAuthors {
		var author bookModel.Author

		BookRepository.DB.FirstOrCreate(&author, &bookModel.Author{Name: newAuthor})
		authors = append(authors, author)
	}
	BookRepository.DB.Model(&book).Association("Authors").Append(authors)
	return authors
}

func (BookRepository *BookRepository) FindByISBN(isbn string) (bookModel.Book, error) {
	var foundBook bookModel.Book
	result := BookRepository.DB.First(&foundBook, "ISBN = ?", isbn)
	if result.RowsAffected == 0 {
		return foundBook, errors.New("not found")
	}
	return foundBook, nil
}

func (BookRepository *BookRepository) FindById(id string) (bookModel.Book, error) {
	var foundBook bookModel.Book
	result := BookRepository.DB.First(&foundBook, "id = ?", id)
	if result.RowsAffected == 0 {
		return foundBook, errors.New("book not found")
	}
	return foundBook, nil
}

func (BookRepository *BookRepository) FindByUserID(userId string) []bookModel.Book {
	var foundBooks []bookModel.Book
	result := BookRepository.DB.Where("user_refer = ?", userId).Find(&foundBooks)
	if result.Error != nil {
        fmt.Printf("Error: %s", result.Error)
    }
	// utils.PrettyDisplay(foundBooks)
	return foundBooks
}

func (BookRepository *BookRepository) DeleteBookById(bookId string) (bookModel.Book, error){
	var deletedBook bookModel.Book

	if err := BookRepository.DB.First(&deletedBook, "id = ?", bookId).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return bookModel.Book{}, errors.New("Book not found")
        }
    }

	if err := BookRepository.DB.Delete(&deletedBook).Error; err != nil {
        return bookModel.Book{}, errors.New("problem on delete")
    }
	return deletedBook, nil
}

func (BookRepository *BookRepository) UpdateBookById(book bookModel.Book) (bookModel.Book, error){
	var newBook bookModel.Book

	if err := BookRepository.DB.First(&newBook, "id = ?", book.ID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return bookModel.Book{}, errors.New("Book not found")
        }
    }

	if err := BookRepository.DB.Model(&newBook).Updates(book).Error; err != nil {
        return bookModel.Book{}, errors.New("problem on update")
    }
	return newBook, nil
}



