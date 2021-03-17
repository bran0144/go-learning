package main

import (
	"fmt"
)
func main() {
	port := 3000
	startWebServer(port, 2)
}

func startWebServer(port, numberOfRetries int) {
	fmt.Println("starting server")
	fmt.Println("server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
}