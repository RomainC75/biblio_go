package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	// -> ID, CreatedAt, UpdatedAt, DeletedAt
	Isbn10      string    `gorm:"varchar:191;unique_index"`
	Isbn13      string    `gorm:"varchar:191;unique_index"`

	Title string `gorm:"varchar:191;"`
	Description string `gorm:"varchar:191;"`
	ReleaseYear uint `gorm:"integer"`

	SeriesNumber int // 0 if not part of a series
	MaxSeriesNumber int // 0 if not finished

	IsPersoEdited bool // case : personnal data

	WeightG uint
	DimensionX float64
	DimensionY float64
	DimensionZ float64
	NumberOfPages int

	// manyToOne
	Editor    Editor `gorm:"foreignKey:EditorRef"`
	EditorRef uint // manyToOne
	
	Links []Link `gorm:"foreignKey:BookRef"`
	// manyToMany
	Languages []Language `gorm:"many2many:book_language;"`
	Genres []Genre `gorm:"many2many:book_genre;"`
	Authors []Author `gorm:"many2many:book_author;"`
}

// One to Many
type Editor struct {
	gorm.Model
  	Name string 	`gorm:"varchar:191;"`
	Books []Book `gorm:"foreignKey:EditorRef"`
}

type Link struct {
	gorm.Model
  	Url string 	`gorm:"varchar:300;"`
	BookRef uint 
}

// Many To Many

type Language struct {
	gorm.Model
  	Name string 	`gorm:"varchar:191;"`
	Books []Book   `gorm:"many2many:book_language;"`
}

type Genre struct {
	gorm.Model
  	Name string 	`gorm:"varchar:191;"`
	Books []Book   `gorm:"many2many:book_genre;"`
}


type Author struct {
	gorm.Model
  	Name string 	`gorm:"varchar:191;"`
	Books []Book   `gorm:"many2many:book_author;"`
}
