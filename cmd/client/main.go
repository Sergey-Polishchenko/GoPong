package main

import (
	"log"

	"gopong/internal/shared"
)

func main() {
	host, err := shared.LoadVar("HOST")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	port, err := shared.LoadVar("PORT")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// TODO:
	// 1. Implement client logic using gRPC
	// 2. Set up connection to server and establish communication
	// 3. Implement game logic, including menu and gameloop
}
