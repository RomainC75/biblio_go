package helper

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	DataResponse "gitub.com/RomainC75/biblio/utils/third-party-apis/openlibrary/responses"
)

func GetDimensions(fullStr string)([]float64, error){
	r, _ := regexp.Compile("[0-9.]+")
	dimensions := r.FindAllString(fullStr, -1)
	if len(dimensions) != 3 {
		return []float64{}, errors.New(fmt.Sprintf("error converting the openLibrary dimensions string - wrong number of numbers found : %s\n", fullStr))
	}
	floatDimensions := []float64{}

	for _, c := range dimensions{
		f, err := strconv.ParseFloat(c, 64)
		if err != nil{
			return []float64{}, errors.New(fmt.Sprintf("error converting the openLibrary dimensions string - convert : %s\n", fullStr))
		}
		floatDimensions = append(floatDimensions, f)
	}
	return floatDimensions, nil
}

func GetWeight(fullStr string)(uint, error){
	r, _ := regexp.Compile("[0-9]+")
	weightStr := r.FindString(fullStr)
	weight, err := strconv.ParseUint(weightStr, 10, 64)
	if err != nil {
		return uint(0), err
	}
	return uint(weight), nil
}

func GetIsbnsFromData( identifiers DataResponse.Identifiers ) (isbn10 string, isbn13 string){	
	value10, ok := identifiers["isbn_10"]
	if ok {
		isbn10 = value10[0]
	}
	value13, ok := identifiers["isbn_13"]
	if ok {
		isbn10 = value13[0]
	}
	return 
}