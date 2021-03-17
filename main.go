package main

import (
	"fmt"
)
func main() {
	port := 3000
	port, err := startWebServer(port)
	fmt.Println(port, err)
}

func startWebServer(port int) (int, error) {
	fmt.Println("starting server")
	fmt.Println("server started on port", port)
	return port, nil
}