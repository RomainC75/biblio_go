package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	responses "gitub.com/RomainC75/biblio/internal/modules/apis/third-party/openlibrary/responses"
	"gitub.com/RomainC75/biblio/pkg/utils"
)

var apiURL = "https://openlibrary.org/"

// isbn : 10/13
func SearchByReqData(queryStr string) (responses.SearchResponseData, error) {
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
		return responses.SearchResponseData{}, err
	}
	defer resp.Body.Close()
	fmt.Println("CODE : ", resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	var response responses.SearchResponseData
	err = decoder.Decode(&response)

	if err != nil {
		return responses.SearchResponseData{}, err
	}
	fmt.Println("===================")
	utils.PrettyDisplay(response)
	// fmt.Println("-----> RESP : ", resp.Body)
	return response, nil
}

func SearchByReqDetails(queryStr string) (responses.SearchResponseDetails, error) {
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
		return responses.SearchResponseDetails{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response responses.SearchResponseDetails
	err = decoder.Decode(&response)

	if err != nil {
		return responses.SearchResponseDetails{}, err
	}
	fmt.Println("--->", response)
	// fmt.Println("-----> RESP : ", resp.Body)
	return response, nil
}

