package services

import (
	"errors"
	"fmt"
	"reflect"
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
		return SearchInApisResponse{}, errors.New(fmt.Sprintf("no book found for isbn %s", isbn))
	}
	fmt.Printf("==> GOOGLE : %+v\n")
	utils.PrettyDisplay(googleBookChanResult.Response)
	

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
	googleBookItem googleBookResponse.Item,
	) SearchInApisResponse{

		// TODO : filter inside googleBook to find the best item possible.

		var releaseDate uint64
		releaseDateStr := ""
		if len(googleBookItem.VolumeInfo.ReleaseDate)>=4{
			releaseDateStr = string(googleBookItem.VolumeInfo.ReleaseDate[0:4])
			releaseDate, _ = strconv.ParseUint(releaseDateStr, 10, 32)
		}

		isbn10, isbn13 := googleBookHelper.ExtractIsbn(googleBookItem.VolumeInfo.Isbns)
		isbn100, isbn130 := openLibraryHelper.GetIsbnsFromData(olDataData[isbn].Identifiers)
		fmt.Println("=>ISBNS", isbn10, isbn13, isbn100, isbn130)

		weight, _ := openLibraryHelper.GetWeight(olDetailsData.Details.Weight)

		fullDimensionsStr, err := openLibraryHelper.GetDimensions(olDetailsData.Details.Dimensions)
		if err != nil{
			fullDimensionsStr = []float64{0, 0, 0}
		}

		var genres []string
		if len(googleBookItem.VolumeInfo.Categories) > 0{
			genres = googleBookItem.VolumeInfo.Categories
		}else{
			genres = []string{}
		}

		var description string
		if len(googleBookItem.VolumeInfo.Description)>0 {
			description = googleBookItem.VolumeInfo.Description
		}else {
			description = googleBookItem.SearchInfo.ShortDescription
		}

		var publisher string
		if len(olDetailsData.Details.Publisher)	> 0{
			publisher = olDetailsData.Details.Publisher[0]
		}else{
			publisher = googleBookItem.Details.VolumeInfo.Publisher
		}

		links := []string{}
		values := reflect.ValueOf(googleBookItem.Details.VolumeInfo.ImageLinks)
		for i := values.NumField()-1 ; i >= 0 ; i-- {
			field := values.Field(i)
			if field.String() != ""{
				links = append(links, field.String())
				break
			}
		}
		if len(links) == 0{
			links = append(links, olDetailsData.ThumbnailUrl)
		}


		return SearchInApisResponse{
			Book: Models.Book{
				Isbn10 : googleBookItem.VolumeInfo.Isbns[0].Identifier,
				Isbn13 : googleBookItem.VolumeInfo.Isbns[1].Identifier,
				Title : googleBookItem.VolumeInfo.Title,
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
			Editor : publisher,
			Links : links,
			Language : []string{ googleBookItem.VolumeInfo.Language },
			// Genres : olDataData.VolumeInfo.Categories,
			Genres : genres,
			Authors : googleBookItem.VolumeInfo.Authors,
		}
}
		




//details


//data


//google