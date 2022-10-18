package server

import (
	"bufio"
	"fmt"
	"os"

	ginserver "github.com/shunta0213/test_go_auth/server/gin_server"
)

func ServerType() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Choose Server Type: ")

	// Server Type
	for {
		fmt.Println("1: gin")
		fmt.Println("2: gRPC")
		scanner.Scan()

		in := scanner.Text()

		switch in {
		case "1":
			fmt.Println("Starting gin server...")
			ginserver.RunGinServer()
		case "2":
			// fmt.Println("Starting gRPC server...")
			fmt.Println("Currently gRPC is not available.")
			fmt.Println("Please use gin server")
		default:
			fmt.Printf("%s is not available option", in)
			fmt.Println("Enter Correctly.")
		}
	}

}
