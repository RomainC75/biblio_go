package dto

type CreateBookRequest struct {
	Title   string          `form:"title" binding:"required,min=3,max=191"`
	ISBN    string          `form:"name" binding:"omitempty,min=10,max=17"`
	Authors []string        `form:"authors" binding:"required"`
	// GenreCode   BookModel.GenreCode `form:"genreCode" binding:"required,numeric"`
	// LanguageCode   BookModel.LanguageCode `form:"languageCode" binding:"required,numeric"`
	// GenreCode   BookModel.GenreCode `form:"genre" binding:"required,oneof=S.F fantastic love"`
}
