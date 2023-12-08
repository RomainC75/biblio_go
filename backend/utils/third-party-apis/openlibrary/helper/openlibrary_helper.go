package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	responses "gitub.com/RomainC75/biblio/utils/third-party-apis/openlibrary/responses"
)

var apiURL = "https://openlibrary.org/"

type SearchByReqDataChan struct{
	Response responses.SearchResponseData
	Err error
}

type SearchByReqDetailsChan struct {
	Response responses.SearchResponseDetails
	Err error
}

// isbn : 10/13
func SearchByReqData(queryStr string, out chan SearchByReqDataChan)  {
	go func(){
		defer close(out)
		resource := "/api/books"
		params := url.Values{}
		params.Add("bibkeys", queryStr)
		params.Add("jscmd", "data")
		params.Add("format", "json")
		

		u, _ := url.ParseRequestURI(apiURL)
		u.Path = resource
		u.RawQuery = params.Encode()
		urlStr := fmt.Sprintf("%v", u)
		fmt.Println("-> ", u)
		fmt.Println("-> ", urlStr)
		resp, err := http.Get(urlStr)

		if err != nil {
			out <- SearchByReqDataChan{responses.SearchResponseData{}, err}
		}
		defer resp.Body.Close()
		fmt.Println("CODE : ", resp.StatusCode)
		decoder := json.NewDecoder(resp.Body)
		var response responses.SearchResponseData
		err = decoder.Decode(&response)

		if err != nil {
			out <- SearchByReqDataChan{responses.SearchResponseData{}, err}
		}
		out <- SearchByReqDataChan{response, nil}
		
	}()
}

func SearchByReqDetails(queryStr string, out chan SearchByReqDetailsChan) {
	go func(){
		defer close(out)
		resource := "/api/books"
		params := url.Values{}
		params.Add("bibkeys", queryStr)
		params.Add("jscmd", "details")
		params.Add("format", "json")
		

		u, _ := url.ParseRequestURI(apiURL)
		u.Path = resource
		u.RawQuery = params.Encode()
		urlStr := fmt.Sprintf("%v", u)
		fmt.Println("-> ", u)
		fmt.Println("-> ", urlStr)
		resp, err := http.Get(urlStr)

		if err != nil {
			out <- SearchByReqDetailsChan{responses.SearchResponseDetails{}, err}
		}
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		var response responses.SearchResponseDetails
		err = decoder.Decode(&response)

		if err != nil {
			out <- SearchByReqDetailsChan{responses.SearchResponseDetails{}, err}
		}
		
		out <- SearchByReqDetailsChan{response, nil}
	}()
}

