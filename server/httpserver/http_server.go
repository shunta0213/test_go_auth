package httpserver

import (
	"fmt"
	"net/http"

	cLog "github.com/shunta0213/test_go_auth/log"
)

func RunHttpServer(ssl bool) {
	// Routes
	http.HandleFunc("/", HealthCheck)

	// Run HTTP Server
	if ssl {
		cLog.Warning("SSL is unavailable now. Use http instead.")
	} else {
		cLog.Println("Welcome to net/http Server.")
		http.ListenAndServe(":8080", nil)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Server.")
}
