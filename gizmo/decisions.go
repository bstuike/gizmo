package gizmo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lg int
var back = "8"
var exit = "9"
var result string
var choice string
var choiceL2 string
var choiceL3 string
var reader = bufio.NewReader(os.Stdin)
var fqdn = ".cfia-acia.inspection.gc.ca"

// The MainChoices function uses a switch statement to direct the user to a chosen task.
func MainChoices() {
	lg = lcid()
	for choice != exit {
		welcome(lg)
		mainMenu(lg)
		choice, _ = reader.ReadString('\n')
		choice = strings.Title(strings.Replace(choice, "\r\n", "", -1))

		switch choice {
		case "0":
			orca()
		case "1":
			password()
		case "2":
			unlock()
		case "3":
			userName()
		case "4":
			computer()
		case "5":
			printer()
		case "6":
			group()
		case "7":
			advancedChoices()
		case "E":
			lg = 0
		case "F":
			lg = 1
		case "9":
		default:
			fmt.Println("\nInvalid choice - Please try another selection")
		}
	}
}

// The advancedChoices function uses a switch statement to launch the appropriate function.
func advancedChoices() {
	choiceL2 = ""
	for choiceL2 != back {
		atPrompt()
		advancedMenu(lg)
		choiceL2, _ = reader.ReadString('\n')
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
			processChoices()
		case "6":
			serviceChoices()
		case "E":
			lg = 0
		case "F":
			lg = 1
		case "8":
		case "9":
			choiceL2 = "8"
			choice = "9"
		default:
			fmt.Println("\nInvalid choice - Please try another selection")
		}
	}
}

// The processChoices function uses a switch statement to take action based on user input.
func processChoices() {
	choiceL3 = ""
	for choiceL3 != back {
		processMenu(lg)
		choiceL3, _ = reader.ReadString('\n')
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
			choiceL2 = "8"
			choiceL3 = "8"
			choice = "9"
		default:
			fmt.Println("\nInvalid choice - Please try another selection")
		}
	}
}

// The serviceChoices function uses a switch statement to take action based on user input.
func serviceChoices() {
	choiceL3 = ""
	for choiceL3 != back {
		serviceMenu(lg)
		choiceL3, _ = reader.ReadString('\n')
		choiceL3 = strings.Title(strings.Replace(choiceL3, "\r\n", "", -1))

		switch choiceL3 {
		case "1":
		case "2":
		case "3":
		case "4":
		case "E":
			lg = 0
		case "F":
			lg = 1
		case "8":
		case "9":
			choiceL2 = "8"
			choiceL3 = "8"
			choice = "9"
		default:
			fmt.Println("\nInvalid choice - Please try another selection")
		}
	}
}
