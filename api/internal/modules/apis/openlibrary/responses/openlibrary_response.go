package responses

type FoundBook struct {
	author []string
	key    string
	title  string
}

type Search struct {
	docs          []FoundBook
	numFound      int
	numFoundExact bool
	num_found     int
	// offset
	q     string
	start int
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
