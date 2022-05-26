package service

import (
	"gin+jwt+postgres/model"
	"gin+jwt+postgres/repo"
)

type UserService interface {
	GetUser(it int) model.User
}
type userService struct {
	userRepo repo.UserRepository
}

func NewUserService() UserService {
	return &userService{
		userRepo: repo.NewUserRepository(),
	}
}
func (s *userService) GetUser(id int) model.User {
	return s.userRepo.ProfileUser(id)
}
