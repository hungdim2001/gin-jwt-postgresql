package route

import (
	"gin+jwt+postgres/controller"

	"github.com/gin-gonic/gin"
)

var (
	userController controller.UserController = controller.NewUserController()
)

func SetUserRoute(r *gin.Engine) {
	userRoute := r.Group("/user")
	{
		userRoute.POST("/", userController.GetUser)

	}

}
