package request

type CreateBookRequest struct {
	Title   string   `form:"title" binding:"required,min=3,max=191"`
	ISBN    string   `form:"name" binding:"required,min=10,max=17"`
	Authors []string `form:"authors" binding:"required"`
	Genre   Genre    `form:"genre" binding:"required,oneof=S.F fantastic love"`
}

type Genre string

const (
	ScienceFiction Genre = "S.F"
	Fantastic      Genre = "fantastic"
	Love           Genre = "love"
)
