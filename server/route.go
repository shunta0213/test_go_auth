package server

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
}
