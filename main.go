package main

import (
	"fmt"
	"os"

	"github.com/thakursaurabh1998/simple-database/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	err := server.StartTCPServer()
	return err
}
