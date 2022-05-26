package route

import (
	"gin+jwt+postgres/controller"

	"github.com/gin-gonic/gin"
)

var (
	authController controller.AuthController = controller.NewAuthController()
)

func SetAuthRoute(r *gin.Engine) {
	authRoute := r.Group("/auth")
	{

		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", authController.Register)

	}

}
