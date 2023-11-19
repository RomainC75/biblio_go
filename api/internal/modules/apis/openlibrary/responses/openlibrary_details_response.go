package responses

type Details struct {
	Title string `json:"title"`
	Publisher []string `json:"publishers"`
	ReleaseStrDate string `json:"publish_date"`
	Dimensions string  `json:"physical_dimensions"`
	Weight string `json:"weight"`
	NumberOfPages int `json:"number_of_pages"`
}

type Root struct {
	ThumbnailUrl  string `json:"thumbnail_url"`
	Details Details `json:"details"`
}

type SearchResponseDetails map[string]Root