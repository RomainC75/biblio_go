package bootstrap

import (
	"gitub.com/RomainC75/biblio/config"
	"gitub.com/RomainC75/biblio/data/database"
	"gitub.com/RomainC75/biblio/data/seeder"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
