package dto

type Login struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Password string `json:"-" form:"password" binding:"required"`
}
