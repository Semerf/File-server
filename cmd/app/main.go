package main

import (
	"fmt"

	"github.com/Semerf/File-server/internal/server"
)

func main() {
	fmt.Println("Hello, world!")
	go server.Server()
}
