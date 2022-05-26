package controller

import (
	"fmt"
	"gin+jwt+postgres/helper"
	"gin+jwt+postgres/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUser(c *gin.Context)
}
type userController struct {
	jwtService  service.JWTService
	userService service.UserService
}

func NewUserController() UserController {
	return &userController{
		jwtService:  service.NewJwtService(),
		userService: service.NewUserService(),
	}
}
func (h *userController) GetUser(c *gin.Context) {
	const BEARER = "Bearer "
	authToken := c.GetHeader("Authorization")
	tokenFromHeader := authToken[len(BEARER):]
	fmt.Printf("token from header:%v", tokenFromHeader)
	token, err := h.jwtService.ValidateToken(tokenFromHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildResponse(false, "invalid token", helper.EmptyObj{}))
		return
	}
	claim := token.Claims.(jwt.MapClaims)
	idString := claim["id"].(string)
	id, _ := strconv.Atoi(idString)
	user := h.userService.GetUser(id)
	c.JSON(http.StatusOK, helper.BuildResponse(true, "get user successfully", user))
}
