package main

import (
	"bufio"
	"fmt"
	"os"
)

type Room struct {
	description string
	queryText string
	options []string
	results []string
}

func printWithNewline(text string) {
	fmt.Print(text, "\n")
}

func (room Room) query() string {
	printWithNewline(room.description)
	printWithNewline(room.queryText)

	for i := range room.options {
		fmt.Printf("%d: %s", i, room.options[i])
	}

	printWithNewline("")

	action, _ := reader.ReadString('\n')

	return room.results[int(action)]
}

var (
	username string
	reader = bufio.NewReader(os.Stdin)
)

func getUserName() {
	fmt.Print("Enter username: ")
	username, _ = reader.ReadString('\n')
}

func explore(room Room) {

}

func main() {
	getUserName()
	fmt.Printf("Hello, %s", username)
	room := Room{"You are in a room.", "What do you want to do?", []string{"Go to the door"}, []string{"You go to the door."}}

	action := room.query()

	fmt.Print(action)
}