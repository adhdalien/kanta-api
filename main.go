package main

import (
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("server failed to listen: %v", err)
	}
}
