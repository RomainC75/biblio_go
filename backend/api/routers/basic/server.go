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

	// usefull if executed in localhost
	//  err := r.Run(fmt.Sprintf("%v:%v", configs.Server.Host, configs.Server.Port))
	err := r.Run(fmt.Sprintf(":%v", configs.Server.Port))
	fmt.Println("SSSTTTTAAARRRT")

	if err != nil {
		log.Fatal("Error in routing !")
		return
	}
}
