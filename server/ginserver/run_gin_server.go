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
		// When ssl is true, run ssl gin server.
		// But now ssl is not available.
		// This is mainly for production environment.

		cLog.Println("Currently https server is unavailable.")
		cLog.Println("Use http server instead.")
		// r.RunTLS(":8080", file path to certificates)
	} else {
		// When ssl is false, run normal gin server.
		// Normal gin server should not be used in production environment.

		r.Run()
	}
}
