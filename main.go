package main

import (
	"gin+jwt+postgres/database"
	"gin+jwt+postgres/route"

	"github.com/gin-gonic/gin"
)

// var (
// 	authController controller.AuthController = controller.NewAuthController()
// )

func main() {
	defer database.CloseDB()
	database.Connect()
	r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.POST("/", func(c *gin.Context) { // test validate token
	// 	const BEARER_SCHEMA = "Bearer "
	// 	authToken := c.GetHeader("Authorization")
	// 	tokenFromHeader := authToken[len(BEARER_SCHEMA):]
	// 	jwtService := service.NewJwtService()
	// 	token, _ := jwtService.ValidateToken(tokenFromHeader)
	// 	if token.Valid {
	// 		claims := token.Claims.(jwt.MapClaims)
	// 		fmt.Printf("\n claims: %v\n", claims["exp"])
	// 		c.JSON(200, gin.H{
	// 			"success": true,
	// 			"message": "token is valid",
	// 		})
	// 		return
	// 	}
	// 	c.JSON(400, gin.H{
	// 		"success": false,
	// 		"message": "token is invalid",
	// 	})
	// 	return
	// })

	// authRoute := r.Group("/auth")
	// {

	// 	authRoute.POST("/login", authController.Login)
	// 	authRoute.POST("/register", authController.Register)
	// }
	route.UseRoute(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
