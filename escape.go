/**
 * Module dependencies
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"runtime"
	"os/exec"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

/**
 * Module
 */

var (
	username string
	reader = bufio.NewReader(os.Stdin)
	clear map[string]func()
	locations map[string]Location
)

/**
 * Structs
 */

type Location struct {
	Description string
	QueryText string
	Options []map[string]string
}

/**
 * Helpers
 */

func CallClear() {
	function, ok := clear[runtime.GOOS]
	if ok {
		fmt.Print(ok)
			function()
	} else {
			panic("Your platform is unsupported! Can't clear the terminal screen.")
	}
}

func getLocations() {
	yamlData, err := ioutil.ReadFile("./locations.yaml")

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(yamlData), &locations)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", locations)
}

func printWithNewline(text string) {
	fmt.Print(text, "\n")
}

func isValidIndex(num int, array []map[string]string) bool {
	if (num >= 0 && num < len(array)) {
		return true
	}
	return false
}

/**
 * Functionality
 */

// give options at location, get response
func (location Location) query() string {
	printWithNewline(location.Description)

	for i := range location.Options {
		printWithNewline("")
		fmt.Printf("%d: %s", i, location.Options[i]["action"])
	}

	printWithNewline("")

	var choiceInt = -1
	var convErr error
	for isValidIndex(choiceInt, location.Options) == false || convErr != nil {
		printWithNewline(location.QueryText)

		choiceStr, _ := reader.ReadString('\n')
		choiceInt, convErr = strconv.Atoi(choiceStr[0:1])
	}

	CallClear()

	printWithNewline(location.Options[choiceInt]["result"])
	printWithNewline("")

	return location.Options[choiceInt]["new_location"]
}

// get username
func getUserName() {
	fmt.Print("Enter username: ")
	username, _ = reader.ReadString('\n')
	fmt.Printf("Welcome, %s", username)
}

// move between locations based on user choices
func explore(locations map[string]Location, start string) {
	key := start
	for key != "END" {
		key = locations[key].query()
	}

	printWithNewline("You won. Nice job. Now go to http://danielrigberg.com to learn more about the guy who got bored and made this silly game.")
}

func init() {
	// read location yaml file
	getLocations()

	// define 'clear' command for each operating system
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
	CallClear()
	getUserName()

	// play the game
	explore(locations, "center")
}
