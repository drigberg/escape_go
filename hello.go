package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	username string
	reader = bufio.NewReader(os.Stdin)
)

func getUserName() {
	fmt.Print("Enter username: ")
	username, _ = reader.ReadString('\n')
}

func main() {
	getUserName()
	fmt.Printf("Hello, %s", username)
}