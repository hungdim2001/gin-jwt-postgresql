package route

import "github.com/gin-gonic/gin"

func UseRoute(r *gin.Engine) {
	// r := gin.Default()
	// auth route
	SetAuthRoute(r)
	SetUserRoute(r)
}
