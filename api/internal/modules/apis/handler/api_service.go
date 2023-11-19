package handler

import (
	"fmt"

	googleBookHelper "gitub.com/RomainC75/biblio/internal/modules/apis/third-party/googleBook/helper"
	openLibraryHelper "gitub.com/RomainC75/biblio/internal/modules/apis/third-party/openlibrary/helper"
	"gitub.com/RomainC75/biblio/pkg/utils"
)


func SearchInApis(Isbn string)  {
	// _, err := openLibraryHelper.SearchByReqDetails(Isbn)

	olData := make(chan openLibraryHelper.SearchByReqDataChan)
	olDetails := make(chan openLibraryHelper.SearchByReqDetailsChan)
	googleBook := make(chan googleBookHelper.SearchGoogleBookChan)


	openLibraryHelper.SearchByReqData(Isbn, olData)
	openLibraryHelper.SearchByReqDetails(Isbn, olDetails)
	googleBookHelper.SearchGoogleBook(Isbn, googleBook)


	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
	olDataResult := <- olData
	if olDataResult.Err != nil {
		fmt.Println("==> DETAILS error ", olDataResult.Err.Error())
	}
	utils.PrettyDisplay(olDataResult.Response)


	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
	olDetailsResult := <- olDetails
	if olDetailsResult.Err != nil {
		fmt.Println("==> DATA error ", olDetailsResult.Err.Error())
	}
	utils.PrettyDisplay(olDetailsResult.Response)


	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
	googleBookChanResponse := <- googleBook
	if googleBookChanResponse.Err != nil {
		fmt.Println("==> GHOOGLE error ", googleBookChanResponse.Err.Error())
	}
	utils.PrettyDisplay(googleBookChanResponse.Response.Kind)
	
	
}


//details


//data


//google