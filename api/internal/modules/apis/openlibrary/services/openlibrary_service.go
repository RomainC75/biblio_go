package services

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

// func New() *UserService {
// 	return &UserService{
// 		userRepository: UserRepository.New(),
// 	}
// }

func Search(queryStr string) (int, error) {
	// res, err := http.Get(url)

	c := colly.NewCollector()
	url := "https://openlibrary.org/works/OL17860744W/A_Court_of_Mist_and_Fury"

	c.OnHTML(".work-title-and-author", func(e *colly.HTMLElement) {
		// Extract data from HTML elements
		title := e.ChildText("h1")
		author := e.ChildText("h2 > a")

		// Clean up the extracted data
		// quote = strings.TrimSpace(quote)
		// author = strings.TrimSpace(author)
		// tags = strings.TrimSpace(tags)

		// Print the scraped data
		fmt.Printf("Quote: %s\nAuthor: %s\n\n", title, author)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return 1, nil
}
