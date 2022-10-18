package ginserver

import (
	"github.com/gin-gonic/gin"
	cLog "github.com/shunta0213/test_go_auth/log"
)

func RunGinServer(ssl bool) {
	r := gin.Default()

	cLog.Printf("Setting routers...")
	SetRoutes(r)
	cLog.Printf("Finished settings routes!")

	if ssl {
		cLog.Println("Currently https server is unavailable.")
		cLog.Println("Use http server instead.")
		// r.RunTLS(":8080", file path to certificates)
	} else {
		r.Run()
	}
}
