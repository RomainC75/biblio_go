package migration

import (
	"fmt"
	"log"

	// articleModels "gitub.com/RomainC75/biblio/internal/modules/article/models"
	bookModels "gitub.com/RomainC75/biblio/internal/modules/book/models"
	userModels "gitub.com/RomainC75/biblio/internal/modules/user/models"
	"gitub.com/RomainC75/biblio/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &bookModels.Author{}, &bookModels.Book{})
	// err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("migration done ...")
}
