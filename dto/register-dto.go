package dto

type Register struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Username string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"-" form:"password" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required" `
}
