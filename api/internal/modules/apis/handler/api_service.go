package handler

import (
	"fmt"

	openLibraryHelper "gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/helper"
)


func SearchInApis(Isbn string)  {
	res, err := openLibraryHelper.SearchByReq(Isbn)
	if err != nil {
		fmt.Println("ERRROR : ", err.Error())
	}
	fmt.Println(res)
}


