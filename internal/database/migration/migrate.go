package migration

import (
	"fmt"
	"log"

	// articleModels "gitub.com/RomainC75/biblio/internal/modules/article/models"
	userModels "gitub.com/RomainC75/biblio/internal/modules/user/models"
	"gitub.com/RomainC75/biblio/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{})
	// err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("migration done ...")
}
