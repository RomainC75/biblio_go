package bootstrap

import (
	"gitub.com/RomainC75/biblio/pkg/configu"
	"gitub.com/RomainC75/biblio/pkg/database"

	// "gitub.com/RomainC75/biblio/pkg/html"
	"gitub.com/RomainC75/biblio/pkg/routing"
	// "gitub.com/RomainC75/biblio/pkg/sessions"
	// "gitub.com/RomainC75/biblio/pkg/static"
)

func Serve() {
	configu.Set()

	database.Connect()

	routing.Init()

	// sessions.Start(routing.GetRouter())

	// static.LoadStatic(routing.GetRouter())
	// html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()
	// other way of using viper

	routing.Serve()
}
