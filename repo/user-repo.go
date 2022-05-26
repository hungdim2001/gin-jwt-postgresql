package repo

import (
	"gin+jwt+postgres/database"
	"gin+jwt+postgres/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user model.User)
	FindAccount(account string) interface{}
	UpdateUser()
	DeleteUser()
	CheckDuplicate(user model.User) (tx *gorm.DB)
	ProfileUser(id int) model.User
}
type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
func (h *userRepository) InsertUser(user model.User) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic("err: hashPassword")
	}
	user.Password = string(hashPassword)
	database.DB.Save(&user)

}
func (h *userRepository) FindAccount(account string) interface{} {
	var userFound model.User
	err := database.DB.Where("email= ? OR username= ?", account, account).Take(&userFound).Error
	if err != nil {
		return nil
	}
	return userFound
}
func (h *userRepository) UpdateUser() {

}
func (h *userRepository) DeleteUser() {

}
func (h *userRepository) CheckDuplicate(user model.User) (tx *gorm.DB) {
	var userFound model.User
	return database.DB.Where("email = ? OR username = ?", user.Email, user.Username).Take(&userFound)
}
func (h *userRepository) ProfileUser(id int) model.User {
	var userFound model.User
	database.DB.Find(&userFound, id)
	return userFound
}
