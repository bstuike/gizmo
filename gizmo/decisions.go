package gizmo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	back = "8"
	exit = "9"
)

var lg int
var result string
var choice string
var choiceL2 string
var choiceL3 string
var reader = bufio.NewReader(os.Stdin)

// The MainOptions function uses a switch statement to direct the user to a chosen task.
func MainOptions() {
	lg = lcid()
	for choice != exit {
		welcome(lg)
		mainMenu(lg)
		choice = getInput(language[5][lg])
		choice = strings.Title(strings.Replace(choice, "\r\n", "", -1))

		switch choice {
		case "0":
			orca()
		case "1":
			password()
		case "2":
			unlock()
		case "3":
			entity()
		case "4":
			computer()
		case "5":
			printer()
		case "6":
			group()
		case "7":
			advancedOptions()
		case "E":
			lg = 0
		case "F":
			lg = 1
		case "9":
		default:
			fmt.Println("\nInvalid choice - Please try another selection")
		}
	}
	defer l.Close()
}

// The advancedOptions function uses a switch statement to launch the appropriate function.
func advancedOptions() {
	for choiceL2 != back {
		atPrompt()
		advancedMenu(lg)
		choiceL2 = getInput(language[5][lg])
		choiceL2 = strings.Title(strings.Replace(choiceL2, "\r\n", "", -1))

		switch choiceL2 {
		case "1":
			logoff()
		case "2":
			atPrompt()
			reboot()
		case "3":
			testConnection()
		case "4":
			disableCard()
		case "5":
			processOptions()
		case "6":
			serviceOptions()
		case "E":
			lg = 0
		case "F":
			lg = 1
		case "8":
		case "9":
			choiceL2 = back
			choice = exit
		default:
			fmt.Println("\nInvalid choice - Please try another selection")
		}
	}
}

// The processOptions function uses a switch statement to take action based on user input.
func processOptions() {
	for choiceL3 != back {
		processMenu(lg) // Display Menu
		choiceL3 = getInput(language[30][lg])
		choiceL3 = strings.Title(strings.Replace(choiceL3, "\r\n", "", -1))

		switch choiceL3 {
		case "1":
			result = localPC()
			fmt.Print(result)
			enterKey()
		case "2":
		case "E":
			lg = 0
		case "F":
			lg = 1
		case "8":
		case "9":
			choiceL2 = back
			choiceL3 = back
			choice = exit
		default:
			fmt.Println("\nInvalid choice - Please try another selection")
		}
	}
}

// The serviceOptions function uses a switch statement to take action based on user input.
func serviceOptions() {
	for choiceL3 != back {
		serviceMenu(lg) // Display Menu
		choiceL3 = getInput(language[30][lg])
		choiceL3 = strings.Title(strings.Replace(choiceL3, "\r\n", "", -1))

		switch choiceL3 {
		case "1":
		case "2":
		case "3":
			restartService()
		case "4":
			stopService()
		case "E":
			lg = 0
		case "F":
			lg = 1
		case "8":
		case "9":
			choiceL2 = back
			choiceL3 = back
			choice = exit
		default:
			fmt.Println("\nInvalid choice - Please try another selection")
		}
	}
}