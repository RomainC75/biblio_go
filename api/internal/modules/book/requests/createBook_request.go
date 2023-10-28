package request

import (
	BookModel "gitub.com/RomainC75/biblio/internal/modules/book/models"
)

type CreateBookRequest struct {
	Title   string          `form:"title" binding:"required,min=3,max=191"`
	ISBN    string          `form:"name" binding:"omitempty,min=10,max=17"`
	Authors []string        `form:"authors" binding:"required"`
	Genre   BookModel.Genre `form:"genre" binding:"required,oneof=S.F fantastic love"`
}
