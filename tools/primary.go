package tools

import (
	"fmt"
	"strconv"
)

// Colour pallette
var colorReset = "\033[0m"
var fgRed = "\033[31m"
var fgGreen = "\033[32m"
var fgYellow = "\033[33m"

// var fgBlue = "\033[34m"
// var fgPurple = "\033[35m"
// var fgCyan = "\033[36m"
var fgWhite = "\033[37m"
var bgRed = "\033[41m"

// var fgBrightRed = "\033[91m"
var fgBrightGreen = "\033[92m"
var fgBrightYellow = "\033[93m"
var fgBrightBlue = "\033[94m"
var fgBrightMagenta = "\033[95m"
var fgBrightCyan = "\033[96m"
var fgBrightWhite = "\033[97m"

// lcid function determines the base language of the operating system.
func lcid() int {
	oslang := 0
	display := powerShellRVS("Get-Culture | select -exp LCID")
	fre, _ := strconv.Atoi(string(display))

	if fre == 3084 {
		oslang = 1
	}
	return oslang
}

// TestDomain function finds the connection speeds of the available Domain Controllers.
func TestDomain() {
	clear()
	// dcSpeed := 1000
	// var fastestDC int

	fmt.Println(string(fgBrightGreen), "\nFinding fastest Domain Controllers...")
	fmt.Println(string(fgBrightYellow), "Testing", string(fgBrightMagenta), fqdn, string(fgBrightYellow), "Domain Controller speed...", string(fgBrightWhite))
	for _, s := range cfia {
		//pingReply := powershellRVS("Test-Connection -ComputerName " + s + fqdn + " -Count 1 -ErrorAction SilentlyContinue | Select-Object responsetime,address")
		pingReplyAdd := powerShellRVS("Test-Connection -ComputerName " + s + fqdn + " -Count 1 -ErrorAction SilentlyContinue | Select-Object address")
		//pingReply := powershellRVS("Test-Connection -TargetName " + s)
		fmt.Println(pingReplyAdd)
		if pingReplyAdd == "CFQCH3AWPDCP002.cfia-acia.inspection.gc.ca" {
			fmt.Print("yes")
		}
	}
	fmt.Println("\nPress the Enter key to continue")
	fmt.Scanln()
}

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

// userName function asks the user for a username and pulls the account information from Active Directory. It also give quick hints & warnings about the account (ex. if expired, disabled, etc.).
func userName() {
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
