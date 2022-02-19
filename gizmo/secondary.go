package gizmo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

// Constants for searching Active Directory.
const (
	filterSAM  = "(sAMAccountName=%s)"
	filterName = "(name=%s)"
)

// Valued declarations
var userSP = []string{"sAMAccountName", "name", "displayName", "title", "mail", "department", "description", "st", "lastLogon", "pwdLastSet", "homeDirectory", "userAccountControl", "canonicalName", "msDS-UserPasswordExpiryTimeComputed"}
var computerSP = []string{"name", "description", "canonicalName", "dNSHostName", "lastLogon", "userAccountControl"}
var result *ldap.SearchResult
var item *ldap.Entry
var ps, _ = exec.LookPath("powershell")
var serviceName string

// The checkError function executes the builtin panic function if an error is detected.
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// The getInput function takes a string prompt and asks the user for input.
func getInput(prompt string) string {
	fmt.Print("\n ", prompt)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	return userInput
}

// The enterKey function pauses the transition to the nest screen until the enter key is pressed.
func enterKey() {
	fmt.Print("\n ", colorReset, language[27][lg])
	fmt.Scanln()
}

// The clear function clears the terminal or screen.
func clear() {
	powerShellEXE("cls")
	fmt.Println(colorReset)
}

func timezone() string {
	utc := strings.TrimSpace(powerShellRVS("Get-TimeZone | Select -exp BaseUtcOffset | Select -exp Hours"))
	return utc
}

// The localPC function gets the name of the local computer.
func localPC() string {
	pc, _ := os.Hostname()
	checkError(err)
	return pc
}

// The cliu function gets the name of the local user.
func cliu() string {
	person, err := user.Current()
	checkError(err)
	return person.Username
}

func intFromString(convertable string) int {
	freshInt, _ := strconv.Atoi(convertable)
	return freshInt
}

func convertLargeInt(value string) string {
	offset := timezone()
	time := powerShellRVS("(Get-Date 1/1/1601).AddDays(" + value + "/864000000000).AddHours(" + offset + ")")
	return time
}

func query(filter string, search []string, object string) {
	clear()
	// Filters must start and finish with ()!
	filterDN := fmt.Sprintf(filter, ldap.EscapeFilter(object))
	searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, filterDN, search, []ldap.Control{})

	result, err = link.Search(searchReq)
	if err != nil {
		fmt.Println("Failed to query LDAP: ", err)
		return
	}
}

func searchValues(nond *ldap.SearchResult) {
	for _, item = range nond.Entries {
		accountCode := item.GetAttributeValue("userAccountControl")
		var acctlo, acctdis, pwdexp, pwdnvexp string = "False", "False", "False", "False"

		switch accountCode {
		case "2":
			acctlo = "True"
		case "16", "514":
			acctdis = "True"
		case "8388608":
			pwdexp = "True"
		case "65536", "66048":
			pwdnvexp = "True"
		default:
		}

		if item.GetAttributeValue("mail") != "" {
			printUserValues()
			fmt.Println(" Password Expired            :", pwdexp)
			fmt.Println(" Password Never Expires      :", pwdnvexp)
			fmt.Println(" Account Locked Out          :", acctlo)
			fmt.Println(" Account Disabled            :", acctdis)
		} else {
			printComputerValues()
		}
	}
}

func printUserValues() {
	fmt.Print("", fgGreen, language[130][lg])
	fmt.Println(fgBrightYellow, item.GetAttributeValue("displayName"), colorReset)
	fmt.Println("\n", language[107][lg], "                   :", item.GetAttributeValue("sAMAccountName"))
	fmt.Println("", language[108][lg], "                  :", item.GetAttributeValue("name"))
	fmt.Println(" Display Name                :", item.GetAttributeValue("displayName"))
	fmt.Println(" Email Address               :", item.GetAttributeValue("mail"))
	fmt.Println(" Department                  :", item.GetAttributeValue("department"))
	fmt.Println(" Title                       :", item.GetAttributeValue("title"))
	fmt.Println(" Description                 :", item.GetAttributeValue("description"))
	fmt.Println(" Province                    :", item.GetAttributeValue("st"))
	fmt.Println(" AD Path                     :", item.GetAttributeValue("canonicalName"))
	fmt.Println(" Home Directory              :", item.GetAttributeValue("homeDirectory"))
	fmt.Println(" Last Logon                  :", strings.TrimSpace(convertLargeInt(item.GetAttributeValue("lastLogon"))))
	fmt.Println(" Password Last Set           :", strings.TrimSpace(convertLargeInt(item.GetAttributeValue("pwdLastSet"))))
	fmt.Println(" Password Expires            :", strings.TrimSpace(convertLargeInt(item.GetAttributeValue("msDS-UserPasswordExpiryTimeComputed"))))
}

func printComputerValues() {
	fmt.Print("", fgGreen, language[105][lg])
	fmt.Println(fgBrightYellow, item.GetAttributeValue("name"), colorReset)
	fmt.Println("\n Full Name           :", item.GetAttributeValue("name"))
	fmt.Println(" AD Path             :", item.GetAttributeValue("conicalName"))
	fmt.Println(" FQDN                :", item.GetAttributeValue("dNSHostName"))
	fmt.Println(" Description         :", item.GetAttributeValue("description"))
	fmt.Println(" Last Logon          :", strings.TrimSpace(convertLargeInt(item.GetAttributeValue("lastLogon"))))
}

// The testConnection function will test the connection to a computer.
func testConnection() {
	var pingTest string
	var pingGroup = [3]int{}
	var pingResult, avgSpeed, index int

	atPrompt()
	fmt.Println(fgYellow, language[149][lg], fgWhite+computerName) // Checking connection to

	for index = 0; index < 3; index++ {
		pingTest = strings.TrimSpace(powerShellRVS("Test-Connection -ComputerName " + computerName + " -Count 1 -ErrorAction SilentlyContinue | Select -exp ResponseTime"))
		pingResult = intFromString(pingTest)
		pingGroup[index] = pingResult
	}
	avgSpeed = (pingGroup[0] + pingGroup[1] + pingGroup[2]) / 3

	if avgSpeed > 0 {
		fmt.Println(fgGreen, language[150][lg]) // Connection succeeded!
	} else {
		fmt.Println(fgRed, language[151][lg]) // Connection failed!
	}
	enterKey()
}

// The disableCard function will disable a network card on a remote computer.
func disableCard() {
	adapters := powerShellRVS("Get-NetAdapter -Name * -IncludeHidden | Format-List -Property deviceid,name")
	fmt.Print(adapters)
	fmt.Print(" " + language[153][lg] + ": ")
	deviceID, _ := reader.ReadString('\n')
	deviceID = strings.Title(strings.Replace(deviceID, "\r\n", "", -1))

	if deviceID == "10" {
		fmt.Println(deviceID)
	} else {
		fmt.Println(deviceID)
	}

	fmt.Println(fgRed, language[158][lg])
	enterKey()
}

func restartService() {
	powerShellEXE("Restart-Service -Name " + serviceName)
}

func stopService() {
	powerShellEXE("Stop-Service -Name " + serviceName)
}

// The reboot function will reboot a remote computer.
func reboot() {
	powerShellEXE("Restart-Computer -ComputerName " + "'" + computerName + "'" + " -Force")
}

// The logoff function will force a logoff.
func logoff() {
	powerShellEXE("shutdown /l /m \\" + computerName + " /t 0")
}

// The atPrompt function displays the computer connected to the application.
func atPrompt() {
	clear()
	fmt.Print(" " + language[15][lg] + ":") // Advanced tools are currently linked to
	fmt.Println(fgGreen, computerName)
	fmt.Println()
}

// The usPrompt function
func usPrompt() {
	clear()
	usTitle()
	fmt.Println("\n [ex. beo, beordma, beo*dma ...]")
	userName = getInput(language[65][lg])
}

// The csPrompt function
func csPrompt() {
	clear()
	csTitle()
	fmt.Println("\n [ex. ncdec652445, 652445, 65244*, 10.141.12.58, ncdec652445.cfia-acia.inspection.gc.ca ...]")
	computerName = getInput(language[71][lg])
}

// The psPrompt function
func psPrompt() {
	clear()
	psTitle()
	fmt.Println("\n [ex. P141022202, 10.136.52.188 ...]")
	printerName = getInput(language[75][lg])
}

// The gsPrompt function
func gsPrompt() {
	clear()
	gsTitle()
	groupName = getInput(language[78][lg])
}

// The powershellEXE function executes a PowerShell command directly.
func powerShellEXE(task string) {
	psCmd := exec.Command(ps, task)
	psCmd.Stdout = os.Stdout
	psCmd.Stderr = os.Stderr

	err := psCmd.Run()
	checkError(err)
}

// The powershellRVS function runs a PowerShell command and returns the output as a String.
func powerShellRVS(task string) string {
	psCmd := exec.Command(ps, task)
	psOut, _ := psCmd.CombinedOutput()
	return string(psOut)
}

/*
// powershellRVI function runs a PowerShell command and returns the output as an Integer.
func powerShellRVI(task string) int {
	psCmd := exec.Command(ps, "/c", task)

	psOut, _ := psCmd.Output()
	byteToInt, _ := strconv.Atoi(string(psOut))
	return byteToInt
}
*/
