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
	// 1. Implement server logic using gRPC
	// 2. Set up server to handling incoming client connections
	// 3. Implement messages handling and response
}
