package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// Id       int    `json:"id" gorm:"primary_key:auto_increment"`
	ID       int    `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Username string `json:"user_name"`
	Password string `json:"-"`
	Name     string `json:"name" `
	Token    string `json:"token,omitempty"`
	TokenExp int    `json:"token_exp,omitempty"`
}
