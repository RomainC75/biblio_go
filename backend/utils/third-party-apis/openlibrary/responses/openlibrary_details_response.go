package responses

type DetailsPart struct {
	Title string `json:"title"`
	Publisher []string `json:"publishers"`
	ReleaseStrDate string `json:"publish_date"`
	Dimensions string  `json:"physical_dimensions"`
	Weight string `json:"weight"`
	NumberOfPages int `json:"number_of_pages"`
}

type Root struct {
	ThumbnailUrl  string `json:"thumbnail_url"`
	Details DetailsPart `json:"details"`
}

type SearchResponseDetails map[string]Root


// func ToBookModel(searchResponse SearchResponse, Authors ) (BookModel.Book, error) {
// 	if len(searchResponse.Docs == 0) {
// 		var bookModel BookModel.Book
// 		return bookModel, nil
// 	}
// 	return BookModel.Book{
// 		Title:   searchResponse.Docs[0].Title,
// 		ISBN:    searchResponse.Q,
// 		Authors: []*models.Author{
// 			Name: searchResponse.Docs[0].Authors[0],
// 		},
// 	}
// }

// if need to transform the response !!
// func ToUser(user userModel.User) User {
// 	return User{
// 		ID:    user.ID,
// 		Name:  user.Name,
// 		Email: user.Email,
// 		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
// 	}
// }
