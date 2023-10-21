package routing

import (
	"fmt"
	"log"

	"gitub.com/RomainC75/biblio/pkg/configu"
)

func Serve() {
	configs := configu.Get()

	// really usefull ?????
	r := GetRouter()
	err := r.Run(fmt.Sprintf("%v:%v", configs.Server.Host, configs.Server.Port))

	if err != nil {
		log.Fatal("Error in routing !")
		return
	}
}
