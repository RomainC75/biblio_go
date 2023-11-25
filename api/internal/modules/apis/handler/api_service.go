package handler

import (
	"fmt"
	"strconv"

	googleBookHelper "gitub.com/RomainC75/biblio/internal/modules/apis/third-party/googleBook/helper"
	googleBookResponse "gitub.com/RomainC75/biblio/internal/modules/apis/third-party/googleBook/responses"
	openLibraryHelper "gitub.com/RomainC75/biblio/internal/modules/apis/third-party/openlibrary/helper"
	openLibraryResponse "gitub.com/RomainC75/biblio/internal/modules/apis/third-party/openlibrary/responses"

	bookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"

	"gitub.com/RomainC75/biblio/pkg/utils"
)

type SearchInApisResponse struct {
		Book  bookModel.Book 
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

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
	olDataResult := <- olData
	if olDataResult.Err != nil {
		fmt.Println("==> DETAILS error ", olDataResult.Err.Error())
	}
	// utils.PrettyDisplay(olDataResult)

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
	olDetailsResult := <- olDetails
	if olDetailsResult.Err != nil {
		fmt.Println("==> DATA error ", olDetailsResult.Err.Error())
	}
	// utils.PrettyDisplay(olDetailsResult)

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
	googleBookChanResult := <- googleBook
	if googleBookChanResult.Err != nil {
		fmt.Println("==> GHOOGLE error ", googleBookChanResult.Err.Error())
	}
	// utils.PrettyDisplay(googleBookChanResult)
	
	compilatedResponse := ApiCombinator(
		isbn,
		olDetailsResult.Response[isbn],
		olDataResult.Response,
		googleBookChanResult.Response,
	)
	utils.PrettyDisplay(compilatedResponse)
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
		fmt.Println("=========================================")
		fmt.Println("=========================================")
		fmt.Println("=>", isbn10, isbn13, isbn100, isbn130)
		fmt.Println("=========================================")

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
			Book: bookModel.Book{
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