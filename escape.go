/**
 * Module dependencies
 */

package main

import (
	"escape_go/helpers"
	"escape_go/structs"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"os/exec"
)

/**
 * Module
 */

var (
	username string
	reader = bufio.NewReader(os.Stdin)
	clear map[string]func()
	locations map[string]structs.Location
)


// give options at location, get response
func query(location structs.Location) string {
	helpers.PrintWithNewline(location.Description)

	for i := range location.Options {
		helpers.PrintWithNewline("")
		fmt.Printf("%d: %s", i, location.Options[i]["action"])
	}

	helpers.PrintWithNewline("")

	var choiceInt = -1
	var convErr error
	for helpers.IsValidIndex(choiceInt, location.Options) == false || convErr != nil {
		helpers.PrintWithNewline("")
		helpers.PrintWithNewline(location.Query)

		choiceStr, _ := reader.ReadString('\n')
		choiceInt, convErr = strconv.Atoi(choiceStr[0:1])
	}

	helpers.CallClear()

	helpers.PrintWithNewline(location.Options[choiceInt]["result"])
	helpers.PrintWithNewline("")

	return location.Options[choiceInt]["new_location"]
}

// get username
func getUserName() {
	fmt.Print("Enter username: ")
	username, _ = reader.ReadString('\n')
	fmt.Printf("Welcome, %s", username)
}

// move between locations based on user choices
func explore(locations map[string]structs.Location, start string) {
	key := start
	for key != "END" {
		key = query(locations[key])
	}

	helpers.PrintWithNewline("You won. Nice job. Now go to http://danielrigberg.com to learn more about the guy who got bored and made this silly game.")
}

func init() {
	// read location yaml file
	locations = helpers.GetLocations("./static/locations.yaml")

	// define 'clear' command for each operating system
	helpers.Clear = make(map[string]func())
	helpers.Clear["linux"] = func() {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
	}
	helpers.Clear["windows"] = func() {
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
	}

	helpers.Clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	helpers.CallClear()
	getUserName()

	// play the game
	explore(locations, "center")
}
