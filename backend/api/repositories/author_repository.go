package repositories

import (
	"errors"

	Models "gitub.com/RomainC75/biblio/data/models"
	"gorm.io/gorm"
)


func (BookRepository *BookRepository) UpdateAuthors(authors []Models.Author) ([]Models.Author, error){
	updatedAuthors := make([]Models.Author,0)

	for _, author := range authors{
		var newAuthor Models.Author
		if err := BookRepository.DB.First(&newAuthor, "id = ?", author).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// return Models.Book{}, errors.New("Book not found")
				break
			}
		}
	
		if err := BookRepository.DB.Model(&newAuthor).Updates(author).Error; err != nil {
			// return Models.Book{}, errors.New("problem on update")
			break
		}
		updatedAuthors = append(updatedAuthors, newAuthor)
	}
	
	return updatedAuthors, nil
}