package main

import (
	"fmt"
	"os"
	"github.com/sandertv/go-raknet"
)

func main() {
	address := os.Args[1]
	if _, err := raknet.Ping(address); err != nil {
		fmt.Printf("Failed to ping %v\n\nerr:  %v", address, err)
		os.Exit(1)
	}
	fmt.Printf("Success to ping %v\n", address)
}