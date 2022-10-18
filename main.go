package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shunta0213/test_go_auth/server"
)

func main() {

	// Run Gin Server
	r := gin.Default()
	server.SetRoutes(r)
	r.Run()

}
