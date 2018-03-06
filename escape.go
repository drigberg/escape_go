package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	username string
	reader = bufio.NewReader(os.Stdin)
)

type Location struct {
	description string
	queryText string
	options [][]string
}

func isValidIndex(num int, array [][]string) bool {
	if (num >= 0 && num < len(array)) {
		return true
	}
	return false
}

func printWithNewline(text string) {
	fmt.Print(text, "\n")
}


func (room Location) query() string {
	printWithNewline(room.description)

	for i := range room.options {
		printWithNewline("")
		fmt.Printf("%d: %s", i, room.options[i][0])
	}

	printWithNewline("")

	var choiceInt = -1
	var convErr error
	for isValidIndex(choiceInt, room.options) == false || convErr != nil {
		printWithNewline(room.queryText)

		choiceStr, _ := reader.ReadString('\n')
		choiceInt, convErr = strconv.Atoi(choiceStr[0:1])
	}

	return room.options[choiceInt][1]
}

func getUserName() {
	fmt.Print("Enter username: ")
	username, _ = reader.ReadString('\n')
	fmt.Printf("Welcome, %s", username)
}

func main() {
	getUserName()
	room := Location{"You are in the center of the room.", "What do you want to do?", [][]string{{"Go to the door", "You went to the door."}, {"Look out the window", "The window is blocked up."}}}
	res := room.query()
	printWithNewline(res)
}
