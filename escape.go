package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"runtime"
	"os/exec"
)

var (
	username string
	reader = bufio.NewReader(os.Stdin)
	clear map[string]func()
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


func (location Location) query() string {
	printWithNewline(location.description)

	for i := range location.options {
		printWithNewline("")
		fmt.Printf("%d: %s", i, location.options[i][0])
	}

	printWithNewline("")

	var choiceInt = -1
	var convErr error
	for isValidIndex(choiceInt, location.options) == false || convErr != nil {
		printWithNewline(location.queryText)

		choiceStr, _ := reader.ReadString('\n')
		choiceInt, convErr = strconv.Atoi(choiceStr[0:1])
	}

	CallClear()

	printWithNewline(location.options[choiceInt][1])
	printWithNewline("")

	return location.options[choiceInt][2]
}

func getUserName() {
	fmt.Print("Enter username: ")
	username, _ = reader.ReadString('\n')
	fmt.Printf("Welcome, %s", username)
}

func explore(locations map[string]Location, start string) {
	key := start
	for key != "END" {
		key = locations[key].query()
	}

	printWithNewline("You won. Nice job. Now go to http://danielrigberg.com to learn more about the guy who got bored and made this silly game.")
}

func CallClear() {
	function, ok := clear[runtime.GOOS]
	if ok {
		fmt.Print(ok)
			function()
	} else {
			panic("Your platform is unsupported! Can't clear the terminal screen.")
	}
}

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
	}
	clear["windows"] = func() {
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
	}

	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	fmt.Print(runtime.GOOS)
	CallClear()

	getUserName()
	locations := make(map[string]Location)
	locations["center"] = Location{"You are in the center of the room.", "What do you want to do?", [][]string{{"Go to the door", "You go to the door.", "door"}, {"Look out the window", "The window is blocked up.", "window"}}}
	locations["window"] = Location{"You are at the window. It is boarded up with metal sheets.", "What do you want to do?", [][]string{{"Go to the door", "You go to the door.", "door"}, {"Bang on the window", "It rattles.", "window"}}}
	locations["door"] = Location{"You are at the door.", "What do you want to do?", [][]string{{"Try to open it", "It opens.", "END"}, {"Go to the window", "You go to the window.", "window"}}}

	explore(locations, "center")
}
