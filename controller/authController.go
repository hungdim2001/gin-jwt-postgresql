package controller

import (
	"fmt"
	"gin+jwt+postgres/dto"
	"gin+jwt+postgres/helper"
	"gin+jwt+postgres/model"
	"gin+jwt+postgres/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController() AuthController {
	return &authController{
		authService: service.NewAuthService(),
		jwtService:  service.NewJwtService(),
	}
}
func (h *authController) Login(c *gin.Context) {

	var loginDTO dto.Login
	errDTO := c.ShouldBind(&loginDTO)
	if errDTO != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildResponse(false, "no information to login ", helper.EmptyObj{}))
		return
	}
	authResult := h.authService.Login(loginDTO)
	fmt.Printf("authResult dto: %v", authResult)
	if v, ok := authResult.(model.User); ok {
		// claims := token.Claims.(jwt.MapClaims)
		// fmt.Printf("\n claims: %v\n", claims["exp"])
		v.Token, v.TokenExp = h.jwtService.GenerateToken(strconv.Itoa(v.ID))
		c.JSON(http.StatusOK, helper.BuildResponse(true, "login successfully", v))
		return
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildResponse(false, "user or password incorrect", helper.EmptyObj{}))

}
func (h *authController) Register(c *gin.Context) {
	var registerDTO dto.Register
	errDTO := c.ShouldBind(&registerDTO)
	fmt.Printf("err:%v", errDTO)
	if errDTO != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildResponse(false, "no information to register ", helper.EmptyObj{}))
		return
	}
	fmt.Printf("\nregisterDTO:%v\n", registerDTO)

	if h.authService.CheckDuplicate(registerDTO) {
		c.JSON(http.StatusConflict, helper.BuildResponse(false, "user already exist ", helper.EmptyObj{}))
		return
	}
	h.authService.Register(registerDTO)
	c.JSON(http.StatusCreated, helper.BuildResponse(true, "user is created", helper.EmptyObj{}))

}
