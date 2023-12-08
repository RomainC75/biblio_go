package repositories

import (
	"errors"
	"fmt"

	Handlers "gitub.com/RomainC75/biblio/api/handlers"
	"gitub.com/RomainC75/biblio/data/database"
	Models "gitub.com/RomainC75/biblio/data/models"
	"gitub.com/RomainC75/biblio/utils"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepo() *BookRepository {
	return &BookRepository{
		DB: database.Connection(),
	}
}

func (BookRepository *BookRepository) CreateBook(bookInfos Handlers.SearchInApisResponse) (Models.Book, error) {
	

	editor := BookRepository.FirstOrCreateEditor(bookInfos.Editor)
	bookInfos.Book.EditorRef = editor.ID
	fmt.Println("====> EDITOR IN DB : ")
	utils.PrettyDisplay(editor)

	// if err := BookRepository.DB.Create(&bookInfos.Book).Scan(&newBook).Error; err!=nil{
	// 	return Models.Book{}, err
	// }
	if err := BookRepository.DB.Create(&bookInfos.Book).Error; err!=nil{
		return Models.Book{}, err
	}

	links := BookRepository.FirstOrCreateLinks(bookInfos.Links, bookInfos.Book.ID)
	bookInfos.Book.Links = links
	languages := BookRepository.FirstOrCreateLanguages(bookInfos.Book, bookInfos.Language)
	bookInfos.Book.Languages = languages
	genres := BookRepository.FirstOrCreateGenres(bookInfos.Book, bookInfos.Genres)
	bookInfos.Book.Genres = genres
	authors := BookRepository.FirstOrCreateAuthors(bookInfos.Book, bookInfos.Authors)
	bookInfos.Book.Authors = authors

	fmt.Println("==>")
	// BookRepository.DB.Model(&book).Association("Authors")
	// BookRepository.DB.Save(&book)
	return bookInfos.Book, nil
}

func (BookRepository *BookRepository) FindBookByISBN(isbn string) (Models.Book, error) {
	var foundBook Models.Book
	if err := BookRepository.DB.Preload("Authors").Preload("Links").Preload("Editor").Preload("Languages").Where("isbn10 = ? OR isbn13 = ?", isbn, isbn).First(&foundBook).Error; err != nil {
		return Models.Book{}, err
	}
	return foundBook, nil
}

// func (BookRepository *BookRepository) CreateBook(book Models.Book) Models.Book {
	// var newBook Models.Book
// 	BookRepository.DB.Create(&book).Scan(&newBook)
// 	// BookRepository.DB.Model(&book).Association("Authors")
// 	// BookRepository.DB.Save(&book)
// 	return newBook
// }

func (BookRepository *BookRepository) FirstOrCreateEditor(editorName string) Models.Editor {
	var editor Models.Editor
	BookRepository.DB.FirstOrCreate(&editor, &Models.Editor{Name: editorName})
	return editor
}

func (BookRepository *BookRepository) FirstOrCreateLinks(newLinks []string, bookId uint ) []Models.Link {
	var links []Models.Link
	for _, newLink := range newLinks {
		var link Models.Link
		BookRepository.DB.FirstOrCreate(&link, &Models.Link{Url: newLink, BookRef: bookId})
		links = append(links, link)
	}
	return links
}

func (BookRepository *BookRepository) FirstOrCreateLanguages(book Models.Book, newLanguages []string) []Models.Language {
	var languages []Models.Language
	for _, newLanguage := range newLanguages {
		var link Models.Language
		BookRepository.DB.FirstOrCreate(&link, &Models.Language{Name: newLanguage})
		languages = append(languages, link)
	}
	// err
	BookRepository.DB.Model(&book).Association("Languages").Append(languages)
	return languages
}

func (BookRepository *BookRepository) FirstOrCreateGenres(book Models.Book, newGenres []string) []Models.Genre {
	var genres []Models.Genre
	for _, newGenre := range newGenres {
		var genre Models.Genre
		BookRepository.DB.FirstOrCreate(&genre, &Models.Genre{Name: newGenre})
		genres = append(genres, genre)
	}
	BookRepository.DB.Model(&book).Association("Genres").Append(genres)
	return genres
}

func (BookRepository *BookRepository) FirstOrCreateAuthors(book Models.Book, newAuthors []string) []Models.Author {
	var authors []Models.Author
	for _, newAuthor := range newAuthors {
		var author Models.Author

		BookRepository.DB.FirstOrCreate(&author, &Models.Author{Name: newAuthor})
		authors = append(authors, author)
	}
	BookRepository.DB.Model(&book).Association("Authors").Append(authors)
	return authors
}

func (BookRepository *BookRepository) FindByISBN(isbn string) (Models.Book, error) {
	var foundBook Models.Book
	result := BookRepository.DB.First(&foundBook, "ISBN = ?", isbn)
	if result.RowsAffected == 0 {
		return foundBook, errors.New("not found")
	}
	return foundBook, nil
}

func (BookRepository *BookRepository) FindById(id string) (Models.Book, error) {
	var foundBook Models.Book
	result := BookRepository.DB.First(&foundBook, "id = ?", id)
	if result.RowsAffected == 0 {
		return foundBook, errors.New("book not found")
	}
	return foundBook, nil
}

func (BookRepository *BookRepository) FindByUserID(userId string) []Models.Book {
	var foundBooks []Models.Book
	result := BookRepository.DB.Where("user_refer = ?", userId).Find(&foundBooks)
	if result.Error != nil {
        fmt.Printf("Error: %s", result.Error)
    }
	// utils.PrettyDisplay(foundBooks)
	return foundBooks
}

func (BookRepository *BookRepository) DeleteBookById(bookId string) (Models.Book, error){
	var deletedBook Models.Book

	if err := BookRepository.DB.First(&deletedBook, "id = ?", bookId).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return Models.Book{}, errors.New("Book not found")
        }
    }

	if err := BookRepository.DB.Delete(&deletedBook).Error; err != nil {
        return Models.Book{}, errors.New("problem on delete")
    }
	return deletedBook, nil
}

func (BookRepository *BookRepository) UpdateBookById(book Models.Book) (Models.Book, error){
	var newBook Models.Book

	if err := BookRepository.DB.First(&newBook, "id = ?", book.ID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return Models.Book{}, errors.New("Book not found")
        }
    }

	if err := BookRepository.DB.Model(&newBook).Updates(book).Error; err != nil {
        return Models.Book{}, errors.New("problem on update")
    }
	return newBook, nil
}



