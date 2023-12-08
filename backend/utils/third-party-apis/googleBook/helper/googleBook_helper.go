package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gocolly/colly/v2"
	"gitub.com/RomainC75/biblio/config"
	"gitub.com/RomainC75/biblio/utils"
	responses "gitub.com/RomainC75/biblio/utils/third-party-apis/googleBook/responses"
)


type SearchGoogleBookChan struct {
	Response responses.GoogleApiResponse
	Err error
}

var apiURL = "https://www.googleapis.com/"

// isbn : 10/13
func SearchGoogleBook(queryStr string, out chan SearchGoogleBookChan) {
	go func(){
		defer close(out)
		resource := "/books/v1/volumes"
		params := url.Values{}
		params.Add("q", queryStr)
		secret := config.Get().Google.Key
		params.Add("key", secret)
		
		u, _ := url.ParseRequestURI(apiURL)
		u.Path = resource
		u.RawQuery = params.Encode()
		urlStr := fmt.Sprintf("%v", u)
		fmt.Println("-> ", u)
		fmt.Println("-> ", urlStr)
		resp, err := http.Get(urlStr)

		if err != nil {
			out <- SearchGoogleBookChan{responses.GoogleApiResponse{}, err}
		}
		defer resp.Body.Close()
		fmt.Println("CODE : ", resp.StatusCode)
		decoder := json.NewDecoder(resp.Body)
		var response responses.GoogleApiResponse
		err = decoder.Decode(&response)

		if err != nil {
			out <- SearchGoogleBookChan{responses.GoogleApiResponse{}, err}
		}
		fmt.Println("===================")
		utils.PrettyDisplay(response)
		// fmt.Println("-----> RESP : ", resp.Body)
		out <- SearchGoogleBookChan{response, nil}
	}()
}

func SearchByReqStandard(queryStr string) (responses.GoogleApiResponse, error) {
	apiURL := "https://openlibrary.org/"
	resource := "/search.json"
	params := url.Values{}
	params.Add("q", queryStr)
	params.Add("limit", "10")
	params.Add("fields", "key,cover_i,title,subtitle,author_name,name")
	params.Add("mode", "everything")
	params.Add("_spellcheck_count", "0")

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	fmt.Println("-> ", u)
	fmt.Println("-> ", urlStr)
	resp, err := http.Get(urlStr)

	if err != nil {
		return responses.GoogleApiResponse{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response responses.GoogleApiResponse
	err = decoder.Decode(&response)

	if err != nil {
		return responses.GoogleApiResponse{}, err
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
