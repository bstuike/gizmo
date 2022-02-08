package gizmo

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-ldap/ldap/v3"
)

const (
	// Colour palette
	colorReset = "\033[0m"
	fgRed      = "\033[31m"
	fgGreen    = "\033[32m"
	fgYellow   = "\033[33m"
	fgBlue     = "\033[34m"
	fgPurple   = "\033[35m"
	fgCyan     = "\033[36m"
	fgWhite    = "\033[37m"
	bgRed      = "\033[41m"

	fgBrightRed     = "\033[91m"
	fgBrightGreen   = "\033[92m"
	fgBrightYellow  = "\033[93m"
	fgBrightBlue    = "\033[94m"
	fgBrightMagenta = "\033[95m"
	fgBrightCyan    = "\033[96m"
	fgBrightWhite   = "\033[97m"

	// Constants for binding and searching the LDAP server
	fqdn     = ".CFIA-ACIA.inspection.gc.ca"
	ldapBind = "CN=Byron Stuike,OU=AB,OU=Administrative Objects,DC=CFIA-ACIA,DC=inspection,DC=gc,DC=ca"
	baseDN   = "DC=CFIA-ACIA,DC=inspection,DC=gc,DC=ca"
)

var ldapURL = "ldaps://CFONK1AWVDCP007" + fqdn
var l, err = ldap.DialURL(ldapURL)
var ldapUser = "Byron Stuike"
var ldapPassword string
var filterDN = fmt.Sprintf("(CN=%s)", ldap.EscapeFilter(ldapUser))

// The TestDomain function finds the connection speeds of the available Domain Controllers.
func TestDomain() {
	clear()
	// dcSpeed := 1000
	// var fastestDC int

	fmt.Println(fgBrightGreen, "\nFinding fastest Domain Controllers...")
	fmt.Println(fgBrightYellow, "Testing", fgBrightMagenta, fqdn, fgBrightYellow, "Domain Controller speed...", fgBrightWhite)
	for _, s := range cfia {
		//pingReply := powershellRVS("Test-Connection -ComputerName " + s + fqdn + " -Count 1 -ErrorAction SilentlyContinue | Select-Object responsetime,address")
		pingReplyAdd := powerShellRVS("Test-Connection -ComputerName " + s + fqdn + " -Count 1 -ErrorAction SilentlyContinue | Select-Object address")
		//pingReply := powershellRVS("Test-Connection -TargetName " + s)
		fmt.Println(pingReplyAdd)
		if pingReplyAdd == "CFQCH3AWPDCP002.cfia-acia.inspection.gc.ca" {
			fmt.Print("yes")
		}
	}
	enterKey()
}

// The LDAPConnect function connects to the best available Domain Controllers.
func LDAPConnect() {
	checkError(err)

	err = l.Bind(ldapBind, ldapPassword)
	checkError(err)

	//defer l.Close()
}

// The lcid function determines the base language of the operating system.
func lcid() int {
	oslang := 0
	display := powerShellRVS("Get-Culture | select -exp LCID")
	fre, _ := strconv.Atoi(display)

	if fre == 3084 {
		oslang = 1
	}
	return oslang
}

// The orca function will verify if the specified user is an ORCA member or not.
func orca() {
	fmt.Println("\nYou chose 0")
	enterKey()
}

// The password function is used to reset a user password in AD. It asks for a new password, if the user must change password at next logon, for a confirmation and if the user wants to check if the account is locked out.
func password() {
	fmt.Println("\nEnter current password")
	OldPassword := getInput()
	fmt.Println("\nEnter new password")
	NewPassword := getInput()
	fmt.Println("\nYou chose 1")
	passwdModReq := ldap.NewPasswordModifyRequest("", OldPassword, NewPassword)
	if _, err = l.PasswordModify(passwdModReq); err != nil {
		log.Fatalf("failed to modify password: %v", err)
	}
	enterKey()
}

// The unlock function will verify if an account is locked out. If yes, it will propose to unlock it.
func unlock() {
	fmt.Println("\nYou chose 2")
	enterKey()
}

// The userName function asks the user for a username and pulls the account information from Active Directory. It also gives quick hints & warnings about the account (ex. if expired, disabled, etc.).
func userName() {
	fmt.Println("\nYou chose 3")
	query()
	enterKey()
}

// The computer function asks the user for a computer name and pulls the machine information from Active Directory. It also gives quick hints & warnings about the account (ex. if expired, disabled, etc.).
func computer() {
	fmt.Println("\nYou chose 4")
	enterKey()
}

// The printer function will ask for printer name, will retrieve the information from AD and test it. Optionally, you can retrieve the full list of CFIA printers.
func printer() {
	fmt.Println("\nYou chose 5")
	enterKey()
}

// The group function asks for a group name and then searches Active Directory.
func group() {
	fmt.Println("\nYou chose 6")
	enterKey()
}
