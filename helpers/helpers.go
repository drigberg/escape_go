/**
 * Module dependencies
 */

package helpers

import (
	"escape_go/structs"
	"fmt"
	"log"
	"runtime"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

/**
 * Module
 */

var (
	Clear map[string]func()
)

func GetLocations(filepath string) map[string]structs.Location {
	var locations map[string]structs.Location

	yamlData, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(yamlData), &locations)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return locations
}

func CallClear() {
	function, ok := Clear[runtime.GOOS]
	if ok {
		fmt.Print(ok)
			function()
	} else {
			panic("Your platform is unsupported! Can't clear the terminal screen.")
	}
}

func PrintWithNewline(text string) {
	fmt.Print(text, "\n")
}

func IsValidIndex(num int, array []map[string]string) bool {
	if (num >= 0 && num < len(array)) {
		return true
	}
	return false
}
