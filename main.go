package main

import (
	"fmt"
	"errors"
)
func main() {
	port := 3000
	err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) error {
	fmt.Println("starting server")
	fmt.Println("server started on port", port)
	return errors.New("something went wrong")
}