package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gocolly/colly/v2"
	responses "gitub.com/RomainC75/biblio/internal/modules/apis/openlibrary/responses"
	"gitub.com/RomainC75/biblio/pkg/utils"
)

// isbn : 10/13
func SearchByReqDetails(queryStr string) (responses.SearchResponseDetails, error) {
	baseURL := "https://openlibrary.org/"
	resource := "/api/books"
	params := url.Values{}
	params.Add("bibkeys", queryStr)
	params.Add("jscmd", "details")
	params.Add("format", "json")
	

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	fmt.Println("-> ", u)
	fmt.Println("-> ", urlStr)
	resp, err := http.Get(urlStr)

	if err != nil {
		return responses.SearchResponseDetails{}, err
	}
	defer resp.Body.Close()
	fmt.Println("CODE : ", resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	var response responses.SearchResponseDetails
	err = decoder.Decode(&response)

	if err != nil {
		return responses.SearchResponseDetails{}, err
	}
	fmt.Println("===================")
	utils.PrettyDisplay(response)
	// fmt.Println("-----> RESP : ", resp.Body)
	return response, nil
}

func SearchByReqStandard(queryStr string) (responses.SearchResponseStandard, error) {
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
		return responses.SearchResponseStandard{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response responses.SearchResponseStandard
	err = decoder.Decode(&response)

	if err != nil {
		return responses.SearchResponseStandard{}, err
	}
	fmt.Println("--->", response)
	// fmt.Println("-----> RESP : ", resp.Body)
	return response, nil
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
