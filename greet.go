package main

import (
	"fmt"
	"os"
	"strings"
)

func main () {
	if len(os.Args) < 2 {
		fmt.Println("Please provide your name as an argument.")
		return
	}

	name := strings.Join(os.Args[1:], " ")
	fmt.Printf("Hello, %s! Welcome to Go programming.\n", name)
}

// Run the code with "go run greet.go YourName" :-)