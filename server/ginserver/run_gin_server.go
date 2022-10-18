package ginserver

import "github.com/gin-gonic/gin"

func RunGinServer() {
	r := gin.Default()
	SetRoutes(r)
	r.Run()
}
