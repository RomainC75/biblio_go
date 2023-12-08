package bootstrap

import (
	"gitub.com/RomainC75/biblio/config"
	"gitub.com/RomainC75/biblio/data/database"

	// "gitub.com/RomainC75/biblio/pkg/html"
	routing "gitub.com/RomainC75/biblio/api/routing"
	// "gitub.com/RomainC75/biblio/pkg/sessions"
	// "gitub.com/RomainC75/biblio/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	// sessions.Start(routing.GetRouter())

	// static.LoadStatic(routing.GetRouter())
	// html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()
	// other way of using viper

	routing.Serve()
}
