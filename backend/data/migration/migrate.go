package migration

import (
	"fmt"
	"log"

	Models "gitub.com/RomainC75/biblio/data/models"

	"gitub.com/RomainC75/biblio/data/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(
		&Models.User{}, 
		&Models.Editor{}, 
		&Models.Book{}, 
		&Models.Link{}, 
		&Models.Language{}, 
		&Models.Genre{}, 
		&Models.Author{},
	)
	// err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("migration done ...")
}
