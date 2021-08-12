package tools

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

//var pc = powershellRVS("hostname")
var pc = "ABCALGC678441D"
var reader = bufio.NewReader(os.Stdin)
var fqdn = ".cfia-acia.inspection.gc.ca"

// OSsLanguage function determines the base language of the operating system.
func OSLanguage() {
	display := powershellRVS("(Get-WmiObject -class win32_operatingsystem).oslanguage")
	eng := display > "1033"
	if eng {
		lg = 0
	} else {
		lg = 1
	}
}

// TestDomain function
func TestDomain() {
	clear()
	// dcSpeed := 1000
	// var fastestDC int

	fmt.Println(string(fgBrightGreen), "\nFinding fastest Domain Controllers...")
	fmt.Println(string(fgBrightYellow), "Testing", string(fgBrightMagenta), fqdn, string(fgBrightYellow), "Domain Controller speed...", string(fgBrightWhite))
	for _, s := range cfia {
		//pingReply := powershellRVS("Test-Connection -ComputerName " + s + fqdn + " -Count 1 -ErrorAction SilentlyContinue | Select-Object responsetime,address")
		pingReplyAdd := powershellRVS("Test-Connection -ComputerName " + s + fqdn + " -Count 1 -ErrorAction SilentlyContinue | Select-Object address")
		//pingReply := powershellRVS("Test-Connection -TargetName " + s)
		fmt.Println(pingReplyAdd)
		if pingReplyAdd == "CFQCH3AWPDCP002.cfia-acia.inspection.gc.ca" {
			fmt.Print("yes")
		}
	}
	fmt.Println("\nPress the Enter key to continue")
	fmt.Scanln()
}

// mainTasks function uses a switch statement to direct the user to a chosen task.
func MainTasks() {
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
			user()
		case "4":
			computer()
		case "5":
			printer()
		case "6":
			group()
		case "7":
			advancedTasks()
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

// advancedTasks function uses a switch statement to launch the appropriate function.
func advancedTasks() {
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
			processTasks()
		case "6":
			serviceTasks()
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

// processTasks function uses a switch statement to take action based on user input.
func processTasks() {
	choiceL3 = ""
	for choiceL3 != back {
		processMenu(lg)
		choiceL3, _ = reader.ReadString('\n')
		choiceL3 = strings.Title(strings.Replace(choiceL3, "\r\n", "", -1))

		switch choiceL3 {
		case "1":
			result = powershellRVS("Get-process -ComputerName " + pc + fqdn)
			fmt.Print(result)
			fmt.Println("\nPress the Enter key to continue")
			fmt.Scanln()
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

// serviceTasks function uses a switch statement to take action based on user input.
func serviceTasks() {
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
