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
		Editor  bookModel.Editor 
		Link  bookModel.Link 
		Language  bookModel.Language 
		Genre  bookModel.Genre 
		Author  bookModel.Author
}

func SearchInApis(Isbn string)(){
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
	utils.PrettyDisplay(olDetailsResult.Response[Isbn])
	

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
	googleBookChanResult := <- googleBook
	if googleBookChanResult.Err != nil {
		fmt.Println("==> GHOOGLE error ", googleBookChanResult.Err.Error())
	}
	utils.PrettyDisplay(googleBookChanResult.Response)
	
	// ApiCombinator(
	// 	olDetailsResult.Response[Isbn],
	// 	olDataResult.Response,
	// 	googleBookChanResult.Response,
	// )
}

func ApiCombinator(
	olDetailsData openLibraryResponse.Root,
	olDataData openLibraryResponse.SearchResponseData,
	googleBookData googleBookResponse.GoogleApiResponse,
	) SearchInApisResponse{
		// TODO : filter inside googleBook to find the best item possible 

		googleBook :=  googleBookData.Items[0]

		releaseDateStr := string(googleBook.VolumeInfo.ReleaseDate[0:4])
		releaseDate, _ := strconv.ParseUint(releaseDateStr, 10, 32)

		
		weight, _ := openLibraryHelper.GetWeight(olDetailsData.Details.Weight)

		fullDimensionsStr, err := openLibraryHelper.GetDimensions(olDetailsData.Details.Dimensions)
		if err != nil{
			fullDimensionsStr = []float64{0,0,0}
		}



		return SearchInApisResponse{
			Book: bookModel.Book{
				Isbn10 : googleBook.VolumeInfo.Isbns[0].Identifier,
				Isbn13 : googleBook.VolumeInfo.Isbns[1].Identifier,
				Title : googleBook.VolumeInfo.Title,
				Description : googleBook.VolumeInfo.Description,
				ReleaseYear: uint(releaseDate),
				// SeriesNumber: ,
				// MaxSeriesNumber
				IsPersoEdited: false,
				WeightG: weight,
				
				DimensionX: fullDimensionsStr[0],
				DimensionY: fullDimensionsStr[1],
				DimensionZ: fullDimensionsStr[2],
	
				NumberOfPages: olDetailsData.Details.NumberOfPages,
	
				// EditorRef: olDetailsData.Details.Publisher[0],
				// Links: 
				// Languages
				// Genres
				// Authors
			},
			Editor : olDetailsData.Details.Publisher[0],
			Links : olDetailsData.ThumbnailUrl,
			Language : googleBook.VolumeInfo.Language,
			Genres : olDataData.VolumeInfo.Categories,
			Authors : olDataData.VolumeInfo.Authors,

			}

}
		




//details


//data


//google