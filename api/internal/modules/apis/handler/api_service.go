package handler

import (
	"fmt"

	openLibraryHelper "gitub.com/RomainC75/biblio/internal/modules/apis/third-party/openlibrary/helper"
	// googleBook "gitub.com/RomainC75/biblio/internal/modules/apis/third-party/googleBook/helper"
)


func SearchInApis(Isbn string)  {
	// _, err := openLibraryHelper.SearchByReqDetails(Isbn)
	_, err := openLibraryHelper.SearchByReqData(Isbn)
	// _, err := googleBook.SearchGoogleBook(Isbn)
	if err != nil {
		fmt.Println("ERRROR : ", err.Error())
	}
	// fmt.Println(res)

	
}


//details


//data


//google