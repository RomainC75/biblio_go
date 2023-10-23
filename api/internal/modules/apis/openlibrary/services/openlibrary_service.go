package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gocolly/colly/v2"
	"gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/responses"
)

// func New() *UserService {
// 	return &UserService{
// 		userRepository: UserRepository.New(),
// 	}
// }

func SearchByReq(queryStr string) (int, error) {
	baseURL := "https://openlibrary.org/"
	resource := "/search.json"
	params := url.Values{}
	params.Add("q", queryStr)
	params.Add("limit", "10")
	params.Add("fields", "key,cover_i,title,subtitle,author_name,name")
	params.Add("mode", "everything")
	params.Add("_spellcheck_count", "0")

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	fmt.Println("-> ", u)
	fmt.Println("-> ", urlStr)
	resp, err := http.Get(urlStr)

	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response responses.SearchResponse
	err = decoder.Decode(&response)

	if err != nil {
		return 0, err
	}
	fmt.Println("--->", response)
	// fmt.Println("-----> RESP : ", resp.Body)
	return 1, nil
}

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
