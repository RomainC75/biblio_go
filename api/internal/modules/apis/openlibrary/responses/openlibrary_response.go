package responses

type FoundBook struct {
	Authors []string `json:"author_name"`
	Key     string   `json:"key"`
	Title   string   `json:"title"`
}

type SearchResponse struct {
	Docs          []FoundBook `json:"docs"`
	NumFound      int         `json:"numFound"`
	NumFoundExact bool        `json:"numFoundExact"`
	Num_found     int         `json:"num_found"`
	Offset        any         `json:"offset"`
	Q             string      `json:"q"`
	Start         int         `json:"start"`
}

// func ToBookModel(searchResponse SearchResponse, Authors ) (BookModel.Book, error) {
// 	if len(searchResponse.Docs == 0) {
// 		var bookModel BookModel.Book
// 		return bookModel, nil
// 	}
// 	return BookModel.Book{
// 		Title:   searchResponse.Docs[0].Title,
// 		ISNB:    searchResponse.Q,
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
