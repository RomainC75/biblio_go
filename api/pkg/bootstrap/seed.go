package bootstrap

import (
	"gitub.com/RomainC75/biblio/internal/database/seeder"
	"gitub.com/RomainC75/biblio/pkg/configu"
	"gitub.com/RomainC75/biblio/pkg/database"
)

func Seed() {
	configu.Set()

	database.Connect()

	seeder.Seed()
}
