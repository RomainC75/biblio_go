package dto

type RegisterRequest struct {
	Email    string `form:"email" binding:"required,email,min=3,max=100"`
	Password string `form:"password" binding:"required,min=5"`
}