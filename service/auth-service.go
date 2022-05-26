package service

import (
	"gin+jwt+postgres/dto"
	"gin+jwt+postgres/model"
	"gin+jwt+postgres/repo"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(loginDTO dto.Login) interface{}
	CheckDuplicate(registerDTO dto.Register) bool

	Register(registerDTO dto.Register)
}
type authService struct {
	userRepo repo.UserRepository
}

func NewAuthService() AuthService {
	return &authService{
		userRepo: repo.NewUserRepository(),
	}
}
func (h *authService) Login(loginDTO dto.Login) interface{} {
	res := h.userRepo.FindAccount(loginDTO.Account)
	if v, ok := res.(model.User); ok {
		isValidPassword := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(loginDTO.Password))

		if isValidPassword != nil {
			return false
		}
		return res
	}
	return false
}
func (h *authService) CheckDuplicate(registerDTO dto.Register) bool {
	var user model.User
	smapping.FillStruct(&user, smapping.MapFields(&registerDTO))

	res := h.userRepo.CheckDuplicate(user)
	return (res.Error == nil)
}

func (h *authService) Register(registerDTO dto.Register) {
	var user model.User
	smapping.FillStruct(&user, smapping.MapFields(&registerDTO))
	h.userRepo.InsertUser(user)
}
