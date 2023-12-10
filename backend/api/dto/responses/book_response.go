package responses

import (
	"errors"
	"reflect"

	"github.com/google/uuid"
	Models "gitub.com/RomainC75/biblio/data/models"
)

type Author struct {
	ID   uuid.UUID
	Name string
}

type Dimensions struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Book struct {
	// -> ID, CreatedAt, UpdatedAt, DeletedAt
	Id uint `json:"id"`
	Isbn10      string    `json:"isbn_10"`
	Isbn13      string    `json:"isbn_13"`

	Title string `json:"title"`
	Description string `json:"description"`
	ReleaseYear uint `json:"release_year"`

	SeriesNumber int `json:"series_number"`
	MaxSeriesNumber int `json:"max_series_number"`

	IsPersoEdited bool `json:"is_perso_edited"`

	WeightG uint `json:"weight_g"`
	Dimensions Dimensions `json:"dimensions"`
	
	NumberOfPages int `json:"number_of_pages"`

	EditorName    string `json:"editor_name"`
	Links []string `json:"links"`
	Languages []string `json:"languages"`

	Genres []string `gorm:"many2many:book_genre;"`
	Authors []string `json:"authors"`
}


func ToBookResponse(book Models.Book) Book {
	extractedLinks, err := GetEachStringField(book.Links, "Url")
	if err != nil {
		extractedLinks = []string{}
	}
	extractedLanguages, err := GetEachStringField(book.Languages, "Name")
	if err != nil {
		extractedLanguages = []string{}
	}
	extractedAuthors, err := GetEachStringField(book.Authors, "Name")
	if err != nil {
		extractedAuthors = []string{}
	}
	extractedGenres, err :=  GetEachStringField(book.Genres, "Name")
	if err != nil {
		extractedGenres = []string{}
	}
	return Book{
		Id:      book.ID,
		Isbn10:    book.Isbn10,
		Isbn13: book.Isbn13,
		Title:   book.Title,
		Description: book.Description,
		ReleaseYear: book.ReleaseYear,
		SeriesNumber: book.SeriesNumber,
		MaxSeriesNumber: book.MaxSeriesNumber,
		IsPersoEdited: book.IsPersoEdited,

		WeightG: book.WeightG,
		Dimensions: Dimensions{
			X: book.DimensionX,
			Y: book.DimensionY,
			Z: book.DimensionZ,
		},

		NumberOfPages: book.NumberOfPages,
		EditorName: book.Editor.Name,
		Links: extractedLinks,
		Languages: extractedLanguages,
		Authors: extractedAuthors,
		Genres: extractedGenres,
	}
}


func GetEachStringField(data interface{}, key string) ([]string, error){
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
        return nil, errors.New("data is not a slice")
    }
	res := []string{}
    for i := 0; i < v.Len(); i++ {
        elem := v.Index(i)
        field := elem.FieldByName(key)
        if !field.IsValid() {
            return nil, errors.New("field does not exist")
        }
        if field.Kind() != reflect.String {
            return nil, errors.New("field is not a string")
        }
        res = append(res, field.String())
    }
    return res, nil
}
