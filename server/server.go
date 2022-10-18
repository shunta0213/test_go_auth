package server

import (
	"bufio"
	"os"

	cLog "github.com/shunta0213/test_go_auth/log"
	"github.com/shunta0213/test_go_auth/server/ginserver"
)

// Run server from stdin
func ServerType() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Get Server Type
		cLog.Println("Choose Server Type: ")
		cLog.Println("1: gin http server")
		cLog.Println("2: gin https server")
		cLog.Println("3: gRPC")
		scanner.Scan()
		in := scanner.Text()

		// run functions according to got text
		switch in {

		// Run gin http server
		// DO NOT use in the production env.
		case "1":
			cLog.Println("Starting gin http server...")
			cLog.Warning("You are using http, everyone can see the body of API request.")
			cLog.Warning("Don't use in the production environment.")
			ginserver.RunGinServer(false)

		// Run gin 'https' server
		case "2":
			cLog.Println("Starting gin https server..")
			ginserver.RunGinServer(true)

		// run gRPC server
		case "3":
			cLog.Println("Currently gRPC is not available.")
			cLog.Println("Please use gin server")

		// When unavailable choice
		default:
			cLog.Printf("%s is not available option", in)
			cLog.Println("Enter Correctly.")
		}
	}

}
