package middleware

import (
	"gin+jwt+postgres/helper"
	"gin+jwt+postgres/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthJWT(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authToken := c.GetHeader("Authorization")
		tokenFromHeader := authToken[len(BEARER_SCHEMA):]
		jwtService := service.NewJwtService()
		token, _ := jwtService.ValidateToken(tokenFromHeader)
		if token.Valid {
			c.Next()
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildResponse(false, "token is invalid", helper.EmptyObj{}))
	}
}
