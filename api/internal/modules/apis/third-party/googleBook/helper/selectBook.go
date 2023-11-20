package helper

import (
	"errors"

	googleResponse "gitub.com/RomainC75/biblio/internal/modules/apis/third-party/googleBook/responses"
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
