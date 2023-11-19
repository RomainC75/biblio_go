package responses

type Details struct {
	Dimensions string  `json:"physical_dimensions"`
}

type Author struct {
	Url string `json:"url"`
	Name string `json:"name"`
}

type Identifiers map[string][]string

type Data struct {
	Title string `json:"title"`
	Authors []Author `json:"authors"`
	NumberOfPages int `json:"number_of_pages"`
	Weight string `json:"weight"`
	Identifiers Identifiers `json:"identifiers"`
	Publisher []map[string]string `json:"publishers"`

	ReleaseStrDate string `json:"publish_date"`
	Subjects []map[string]string `json:"subjects"`
	SubjectPlaces []map[string]string `json:"subject_places"`
	SubjectPeople []map[string]string `json:"subject_people"`

	CoverUrls  map[string]string `json:"thumbnail_url"`
	
}

type SearchResponseData map[string]Data
