package ginserver

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	cLog "github.com/shunta0213/test_go_auth/log"
)

func RunGinServer(ssl bool, db *sql.DB) {
	r := gin.Default()

	cLog.Println("Successfully connected to database.")

	cLog.Println("Setting routers...")
	SetRoutes(r, db)
	cLog.Println("Finished settings routes!")

	if ssl {
		cLog.Println("Currently https server is unavailable.")
		cLog.Println("Use http server instead.")
		// r.RunTLS(":8080", file path to certificates)
	} else {
		r.Run()
	}
}
