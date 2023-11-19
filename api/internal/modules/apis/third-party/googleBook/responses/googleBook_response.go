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
}

type GoogleApiResponse struct {
	Kind string `json:"kind"`
	TotalItems int `json:"totalItems"`
	Items []Item `json:"items"`
}