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
		cLog.Println("1: gin")
		cLog.Println("2: gRPC")
		scanner.Scan()
		in := scanner.Text()

		// run functions according to got text
		switch in {
		// Run gin server
		case "1":
			cLog.Println("Starting gin server...")
			ginserver.RunGinServer()

		// run gRPC server
		case "2":
			cLog.Println("Currently gRPC is not available.")
			cLog.Println("Please use gin server")

		// When unavailable choice
		default:
			cLog.Printf("%s is not available option", in)
			cLog.Println("Enter Correctly.")
		}
	}

}
