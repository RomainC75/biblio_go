package services

import (
	"errors"
	"fmt"
	"strconv"

	"gitub.com/RomainC75/biblio/utils"
	googleBookHelper "gitub.com/RomainC75/biblio/utils/third-party-apis/googleBook/helper"
	googleBookResponse "gitub.com/RomainC75/biblio/utils/third-party-apis/googleBook/responses"
	openLibraryHelper "gitub.com/RomainC75/biblio/utils/third-party-apis/openlibrary/helper"
	openLibraryResponse "gitub.com/RomainC75/biblio/utils/third-party-apis/openlibrary/responses"

	Models "gitub.com/RomainC75/biblio/data/models"
)

type SearchInApisResponse struct {
		Book  Models.Book 
		Editor  string
		Links  []string
		Language []string
		Genres  []string
		Authors  []string
}

func SearchInApis(isbn string)(SearchInApisResponse, error){
	olData := make(chan openLibraryHelper.SearchByReqDataChan)
	olDetails := make(chan openLibraryHelper.SearchByReqDetailsChan)
	googleBook := make(chan googleBookHelper.SearchGoogleBookChan)

	openLibraryHelper.SearchByReqData(isbn, olData)
	openLibraryHelper.SearchByReqDetails(isbn, olDetails)
	googleBookHelper.SearchGoogleBook(isbn, googleBook)

	olDataResult := <- olData
	if olDataResult.Err != nil {
		fmt.Println("==> DETAILS error ", olDataResult.Err.Error())
	}

	olDetailsResult := <- olDetails
	if olDetailsResult.Err != nil {
		fmt.Println("==> DATA error ", olDetailsResult.Err.Error())
	}

	googleBookChanResult := <- googleBook
	if googleBookChanResult.Err != nil {
		fmt.Println("==> GHOOGLE error ", googleBookChanResult.Err.Error())
	}
	fmt.Printf("==> GOOGLE : %+v\n")
	utils.PrettyDisplay(googleBookChanResult.Response)
	
	fmt.Printf("==> ITEMS :  : %+d\n", googleBookChanResult.Response.TotalItems)
	if googleBookChanResult.Response.TotalItems == 0 {
		return SearchInApisResponse{}, errors.New(fmt.Sprintf("no book found for isbn %s", isbn))
	}

	compilatedResponse := ApiCombinator(
		isbn,
		olDetailsResult.Response[isbn],
		olDataResult.Response,
		googleBookChanResult.Response,
	)
	
	return compilatedResponse, nil
}

func ApiCombinator(
	isbn string, 
	olDetailsData openLibraryResponse.Root,
	olDataData openLibraryResponse.SearchResponseData,
	googleBookData googleBookResponse.GoogleApiResponse,
	) SearchInApisResponse{

		// TODO : filter inside googleBook to find the best item possible.
		googleBook, _ := googleBookHelper.SelectBook(isbn, googleBookData.Items)

		var releaseDate uint64
		releaseDateStr := ""
		if len(googleBook.VolumeInfo.ReleaseDate)>=4{
			releaseDateStr = string(googleBook.VolumeInfo.ReleaseDate[0:4])
			releaseDate, _ = strconv.ParseUint(releaseDateStr, 10, 32)
		}

		isbn10, isbn13 := googleBookHelper.ExtractIsbn(googleBook.VolumeInfo.Isbns)
		isbn100, isbn130 := openLibraryHelper.GetIsbnsFromData(olDataData[isbn].Identifiers)
		fmt.Println("=>ISBNS", isbn10, isbn13, isbn100, isbn130)

		weight, _ := openLibraryHelper.GetWeight(olDetailsData.Details.Weight)

		fullDimensionsStr, err := openLibraryHelper.GetDimensions(olDetailsData.Details.Dimensions)
		if err != nil{
			fullDimensionsStr = []float64{0, 0, 0}
		}

		var genres []string
		if len(googleBook.VolumeInfo.Categories) > 0{
			genres = googleBook.VolumeInfo.Categories
		}else{
			genres = []string{}
		}

		var description string
		if len(googleBook.VolumeInfo.Description)>0 {
			description = googleBook.VolumeInfo.Description
		}else {
			description = googleBook.SearchInfo.ShortDescription
		}

		return SearchInApisResponse{
			Book: Models.Book{
				Isbn10 : googleBook.VolumeInfo.Isbns[0].Identifier,
				Isbn13 : googleBook.VolumeInfo.Isbns[1].Identifier,
				Title : googleBook.VolumeInfo.Title,
				Description : description,
				ReleaseYear: uint(releaseDate),
				// SeriesNumber: ,
				// MaxSeriesNumber
				IsPersoEdited: false,
				WeightG: weight,
				
				DimensionX: fullDimensionsStr[0],
				DimensionY: fullDimensionsStr[1],
				DimensionZ: fullDimensionsStr[2],
	
				NumberOfPages: olDetailsData.Details.NumberOfPages,
			},
			Editor : olDetailsData.Details.Publisher[0],
			Links : []string{olDetailsData.ThumbnailUrl},
			Language : []string{ googleBook.VolumeInfo.Language },
			// Genres : olDataData.VolumeInfo.Categories,
			Genres : genres,
			Authors : googleBook.VolumeInfo.Authors,
		}
}
		




//details


//data


//google