package ginserver

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shunta0213/test_go_auth/server/ginserver/controllers"
	"github.com/shunta0213/test_go_auth/server/ginserver/services"
)

func SetRoutes(r *gin.Engine, DB *sql.DB) {

	// For Check Api connections
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello!",
		})
	})

	userController := controllers.NewUserController(services.NewUserService(DB))
	r.GET("/signin", func(ctx *gin.Context) { userController.SignIn(ctx) })
	r.POST("/signup", func(ctx *gin.Context) { userController.SignUp(ctx) })
}
