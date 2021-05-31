package tools

import (
	"fmt"
)

// Colour pallette
var colorReset = "\033[0m"
var fgRed = "\033[31m"
var fgGreen = "\033[32m"
var fgYellow = "\033[33m"
var fgBlue = "\033[34m"
var fgPurple = "\033[35m"
var fgCyan = "\033[36m"
var fgWhite = "\033[37m"
var bgRed = "\033[41m"
var fgBrightRed = "\033[91m"
var fgBrightGreen = "\033[92m"
var fgBrightYellow = "\033[93m"
var fgBrightBlue = "\033[94m"
var fgBrightMagenta = "\033[95m"
var fgBrightCyan = "\033[96m"
var fgBrightWhite = "\033[97m"

// orca function will verify if the specified user is an ORCA member or not.
func orca() {
	fmt.Println("\nYou chose 0")
}

// password function is used to reset a user password in AD. It asks for a new password, if the user must change password at next logon, for a confirmation and if the user wants to check if the account is locked out.
func password() {
	fmt.Println("\nYou chose 1")
}

// unlock function will verify if an account is locked out. If yes, it will propose to unlock it.
func unlock() {
	fmt.Println("\nYou chose 2")
}

// user function asks the user for a username and pulls the account information from Active Directory. It also give quick hints & warnings about the account (ex. if expired, disabled, etc.).
func user() {
	fmt.Println("\nYou chose 3")
}

// computer function asks the user for a computer name and pulls the machine information from Active Directory. It also give quick hints & warnings about the account (ex. if expired, disabled, etc.).
func computer() {
	fmt.Println("\nYou chose 4")
}

// printer function will ask for printer name, will retrieve the information from AD and test it. Optionally, you can retreive the full list of CFIA printers.
func printer() {
	fmt.Println("\nYou chose 5")
}

// group function asks for a group name and then searches Active Directory.
func group() {
	fmt.Println("\nYou chose 6")
}
