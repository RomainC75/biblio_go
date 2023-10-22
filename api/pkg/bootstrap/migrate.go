package bootstrap

import (
	"gitub.com/RomainC75/biblio/internal/database/migration"
	"gitub.com/RomainC75/biblio/pkg/configu"
	"gitub.com/RomainC75/biblio/pkg/database"
)

func Migrate() {
	configu.Set()

	database.Connect()

	migration.Migrate()
}
