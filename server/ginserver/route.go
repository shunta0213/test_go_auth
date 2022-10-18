package ginserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {

	// For Check Api connections
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello!",
		})
	})

	// SignIn
	r.GET("/signin", func(ctx *gin.Context) {})

	// SingUp
	r.POST("/signup", func(ctx *gin.Context) {})
}
