package database

import (
	"fmt"
	"log"

	"gitub.com/RomainC75/biblio/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	cfg := config.Get()

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
		return
	}

	DB = db
}
