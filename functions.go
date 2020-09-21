package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 2000
	startWebServer(port)
	multipleParams(2, 4, "mark", "kim")

	total := sum(1, 5)
	fmt.Println(total)

	//if want both params
	isValidPort, error := checkPort(port)
	fmt.Println(error)
	fmt.Println(isValidPort)

	// if interested in one response
	_, responseError := checkPort(3000)
	fmt.Println(responseError)
}

func startWebServer(port int) {
	fmt.Println("Starting Server")
	// Start Server
	fmt.Println("Server Started on port ", port)
}

func multipleParams(age, height int, name, location string) {
	fmt.Println(age, height, name, location)
}

func sum(a, b int) int {
	return a + b
}

func checkPort(port int) (bool, error) {
	if port < 3000 {
		return false, errors.New("invalid port number")
	}
	return true, nil
}
