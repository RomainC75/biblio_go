package repositories

import (
	"errors"

	bookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
	"gorm.io/gorm"
)


func (BookRepository *BookRepository) UpdateAuthors(authors []bookModel.Author) ([]bookModel.Author, error){
	updatedAuthors := make([]bookModel.Author,0)

	for _, author := range authors{
		var newAuthor bookModel.Author
		if err := BookRepository.DB.First(&newAuthor, "id = ?", author).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// return bookModel.Book{}, errors.New("Book not found")
				break
			}
		}
	
		if err := BookRepository.DB.Model(&newAuthor).Updates(author).Error; err != nil {
			// return bookModel.Book{}, errors.New("problem on update")
			break
		}
		updatedAuthors = append(updatedAuthors, newAuthor)
	}
	
	return updatedAuthors, nil
}