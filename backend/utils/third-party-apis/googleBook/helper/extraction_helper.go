package helper

import (
	"errors"

	googleResponse "gitub.com/RomainC75/biblio/utils/third-party-apis/googleBook/responses"
)

func SelectBook(isbn string, items []googleResponse.Item) (googleResponse.Item, error) {
	for _, item := range items{
		isbns := item.VolumeInfo.Isbns
		for _, isbnStruct := range isbns{
			if isbnStruct.Identifier == isbn{
				return item, nil
			}
		}
	}
	return googleResponse.Item{}, errors.New("book not found in google response")
}


func ExtractIsbn( industryIds []googleResponse.IndustryIdentifiers )(isbn10 string, isbn13 string){

	for _, identifier := range industryIds {
		if identifier.Type == "ISBN_10"{
			isbn10 = identifier.Identifier
		}else if identifier.Type == "ISBN_13"{
			isbn13 = identifier.Identifier
		}
	}
	return
}