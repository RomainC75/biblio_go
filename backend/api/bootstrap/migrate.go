package bootstrap

import (
	"gitub.com/RomainC75/biblio/config"
	"gitub.com/RomainC75/biblio/data/database"
	"gitub.com/RomainC75/biblio/data/migration"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
