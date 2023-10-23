package responses

type FoundBook struct {
	Author []string `json:"authot"`
	Key    string   `json:"key"`
	Title  string   `json:"title"`
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

// if need to transform the response !!
// func ToUser(user userModel.User) User {
// 	return User{
// 		ID:    user.ID,
// 		Name:  user.Name,
// 		Email: user.Email,
// 		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
// 	}
// }
