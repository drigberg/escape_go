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

func printWithNewline(text string) {
	fmt.Print(text, "\n")
}

type Room struct {
	description string
	queryText string
	options []string
	results map[string]string
}


func (room Room) query() string {
	printWithNewline(room.description)
	printWithNewline(room.queryText)

	for i := range room.options {
		fmt.Printf("%d: %s", i, room.options[i])
	}

	printWithNewline("")

	res, _ := reader.ReadString('\n')

	// use first byte of response
	action := res[0:1]

	return room.results[action]
}

func getUserName() {
	fmt.Print("Enter username: ")
	username, _ = reader.ReadString('\n')
	fmt.Printf("Hello, %s", username)
}

func main() {
	getUserName()

	roomResults := make(map[string]string)
	roomResults["0"] = "You went to the door."

	room := Room{"You are in a room.", "What do you want to do?", []string{"Go to the door"}, roomResults}

	res := room.query()

	fmt.Print(res, "\n")
}
