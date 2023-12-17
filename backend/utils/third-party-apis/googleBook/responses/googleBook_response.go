package responses

type IndustryIdentifiers struct {
	Type string `json:"type"`
	Identifier string `json:"identifier"`
}

type VolumeInfo struct {
	Title string `json:"title"`
	Authors []string `json:"authors"`
	Editor string `json:"publisher"`
	ReleaseDate string `json:"publishedDate"`
	Description string `json:"description"`
	Isbns []IndustryIdentifiers `json:"industryIdentifiers"`
	PageCount int `json:"pageCount"`
	Categories []string `json:"categories"`
	Language string `json:"language"`
}

type SaleInfo struct {
	Country string `json:"country"`
}

type SearchInfo struct {
	ShortDescription string `json:"textSnippet"`
}

type Item struct {
	Kind string `json:"kind"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
	SaleInfo `json:"saleInfo"`
	SearchInfo SearchInfo `json:"searchInfo"`
	SelfLink string `json:"selfLink"`
	Details GoogleDetails `json:"details"`
}

type GoogleApiResponse struct {
	Kind string `json:"kind"`
	TotalItems int `json:"totalItems"`
	Items []Item `json:"items"`
}

type GoogleDetails struct{
	VolumeInfo struct{
		ImageLinks GoogleImageLinks `json:"imageLinks"`
		Categories []string `json:"categories"`
		Publisher string `json:"publisher"`
	} `json:"volumeInfo"`
}

type GoogleImageLinks struct{
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail string `json:"thumbnail"`
	Small string `json:"small"`
	Medium string `json:"medium"`
	Large string `json:"large"`
	ExtraLarge string `json:"extraLarge"`
}