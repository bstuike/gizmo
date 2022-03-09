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

type ADObject struct {
	sam, name, display, description, path, lastLogon string
	// Embedded structures
	employee Employee
	machine  Machine
}

type Employee struct {
	email, department, title, description, province, path, directory, lastLogon string
}

type Machine struct {
	fqdn string
}

type Values struct {
	accountCode, badPasswordCount, accountExpires, passwordExpires, passwordLastSet string
	// Embedded structure
	derived Derived
}

type Derived struct {
	accountLocked, accountDisabled, accountExpired, passwordExpired, passwordNeverExpires string
}

// Constants for searching Active Directory.
const (
	filterSAM  = "(sAMAccountName=%s)"
	filterName = "(name=%s)"
)

// Open declarations.
var result *ldap.SearchResult
var item *ldap.Entry
var index, place int

// Valued declarations.
var time = strings.TrimSpace(powerShellRVS("(Get-date).TofileTime()"))
var ps, _ = exec.LookPath("powershell")
var controllers, lockedDCs []string

//var lockedDCs []string
var ado = ADObject{}
var uav = Values{}

// The getInput function takes a string prompt and asks the user for input.
func getInput(prompt string) string {
	fmt.Print("\n ", prompt)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	return userInput
}

// The checkError function executes the builtin panic function if an error is detected.
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
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

// The localPC function gets the name of the local computer.
func localPC() string {
	pc, _ := os.Hostname()
	checkError(err)
	return pc
}

// The cliu function gets the name of the local user.
func cliu() string {
	var account string
	var accountf string
	var accountl string

	person, err := user.Current()
	checkError(err)
	account = person.Name
	s := strings.Split(account, " ")
	accountf = strings.TrimSpace(s[1])
	accountl = strings.TrimSpace(strings.TrimSuffix(s[0], ","))
	account = accountf + " " + accountl
	return account
}

// The query function searches AD though an established LDAP connection.
func query(link *ldap.Conn, filter string, search []string, object string) {
	// Filters must start and finish with ()!
	filterDN := fmt.Sprintf(filter, ldap.EscapeFilter(object))
	searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, filterDN, search, []ldap.Control{})

	result, err = link.Search(searchReq)
	if err != nil {
		fmt.Println("Failed to query LDAP: ", err)
		return
	}
}

// The searchBadPassword function scans a specified DC to determine how many bad password attempts have occured.
func searchBadPassword(drip *ldap.SearchResult, controller string) {
	var bpcOriginal string
	var bpc int

	for _, item = range drip.Entries {
		bpcOriginal = item.GetAttributeValue("badPwdCount")
		bpc = intFromString(bpcOriginal)
		if bpc > 3 {
			controllers = append(controllers[:index], fgRed+controller+"         "+bpcOriginal+"         Locked")
			lockedDCs = append(lockedDCs[:place], controller)
			place++
		} else if bpc < 4 && bpc > 0 {
			controllers = append(controllers[:index], fgYellow+controller+"         "+bpcOriginal+"         Okay")
		} else {
			controllers = append(controllers[:index], fgGreen+controller+"         "+bpcOriginal+"         Good")
		}
	}
}

// The assignObjectValues function builds two objects and uses the contained information to determine user statuses such as password expiration and locked out account.
func assignObjectValues(nond *ldap.SearchResult) {
	for _, item = range nond.Entries {
		ado = ADObject{item.GetAttributeValue("sAMAccountName"), item.GetAttributeValue("name"), item.GetAttributeValue("displayName"), item.GetAttributeValue("description"), item.GetAttributeValue("canonicalName"), strings.TrimSpace(convertLargeInt(item.GetAttributeValue("lastLogonTimestamp"))), Employee{item.GetAttributeValue("mail"), item.GetAttributeValue("department"), item.GetAttributeValue("title"), item.GetAttributeValue("description"), item.GetAttributeValue("st"), item.GetAttributeValue("canonicalName"), item.GetAttributeValue("homeDirectory"), strings.TrimSpace(convertLargeInt(item.GetAttributeValue("lastLogonTimestamp")))}, Machine{item.GetAttributeValue("dNSHostName")}}
		uav = Values{item.GetAttributeValue("userAccountControl"), item.GetAttributeValue("badPwdCount"), item.GetAttributeValue("accountExpires"), item.GetAttributeValue("msDS-UserPasswordExpiryTimeComputed"), strings.TrimSpace(convertLargeInt(item.GetAttributeValue("pwdLastSet"))), Derived{"False", "False", "False", "False", "False"}}

		if uav.badPasswordCount > "3" || uav.accountCode == "2" {
			uav.derived.accountLocked = "True"
		}

		if uav.accountExpires == "0" {
			uav.accountExpires = "Never"
		} else {
			uav.accountExpires = strings.TrimSpace(convertLargeInt(item.GetAttributeValue("accountExpires")))
			if time > uav.accountExpires {
				uav.derived.accountExpired = "True"
			}
		}

		if uav.passwordExpires == "0" {
			uav.passwordExpires = "Never"
			uav.derived.passwordNeverExpires = "True"
		} else {
			uav.passwordExpires = strings.TrimSpace(convertLargeInt(item.GetAttributeValue("msDS-UserPasswordExpiryTimeComputed")))
			if time > uav.passwordExpires {
				uav.derived.passwordExpired = "True"
			}
		}

		if ado.employee.email != "" {
			printUserValues()
		} else {
			printComputerValues()
		}
	}
}

// The printUserValues function prints all the previously collected user information to the terminal.
func printUserValues() {
	clear()
	fmt.Print(" ", fgGreen, language[130][lg])
	fmt.Println(fgBrightYellow, ado.display, colorReset)
	fmt.Println("\n", language[107][lg], "                   :", ado.sam)
	fmt.Println("", language[108][lg], "                  :", ado.name)
	fmt.Println(" Display Name                :", ado.display)
	fmt.Println(" Email Address               :", ado.employee.email)
	fmt.Println(" Department                  :", ado.employee.department)
	fmt.Println(" Title                       :", ado.employee.title)
	fmt.Println(" Description                 :", ado.employee.description)
	fmt.Println(" Province                    :", ado.employee.province)
	fmt.Println(" AD Path                     :", ado.employee.path)
	fmt.Println(" Home Directory              :", ado.employee.directory)
	fmt.Println(" Last Logon                  :", ado.employee.lastLogon)
	fmt.Println(" Account Locked Out          :", uav.derived.accountLocked)
	fmt.Println(" Account Disabled            :", uav.derived.accountDisabled)
	fmt.Println(" Account Expired             :", uav.derived.accountExpired)
	fmt.Println(" Account Expires             :", uav.accountExpires)
	fmt.Println(" Password Last Set           :", uav.passwordLastSet)
	fmt.Println(" Password Expires            :", uav.passwordExpires)
	fmt.Println(" Password Expired            :", uav.derived.passwordExpired)
	fmt.Println(" Password Never Expires      :", uav.derived.passwordExpired)
}

// The printUserValues function prints all the previously collected computer information to the terminal.
func printComputerValues() {
	clear()
	fmt.Print(" ", fgGreen, language[105][lg])
	fmt.Println(fgBrightYellow, ado.name, colorReset)
	fmt.Println("\n Full Name           :", ado.name)
	fmt.Println(" AD Path             :", ado.employee.path)
	fmt.Println(" FQDN                :", ado.machine.fqdn)
	fmt.Println(" Description         :", ado.employee.description)
	fmt.Println(" Last Logon          :", ado.employee.lastLogon)
}

// The printControllerValues function prints all the previously collected bad password information accross all DC's to the terminal.
func printControllerValues() {
	clear()
	fmt.Println("     DC Name         Attempts     Status")
	fmt.Println(" ---------------     --------     ------")
	for _, s := range controllers {
		fmt.Println(" " + s)
	}
}

// The callForUnlock function asks the user if an unlock should be attempted and responds accordingly.
func callForUnlock() {
	if (len(lockedDCs)) > 0 {
		fmt.Println("\n"+fgCyan, userName+colorReset, language[93][lg])
		answer := strings.ToUpper(getInput(language[94][lg]))
		if answer == "Y" {
			unlock()
			fmt.Println("\n", language[97][lg]+fgCyan, userName)
		}
	} else {
		fmt.Println("\n", colorReset+language[91][lg]+fgCyan, userName, colorReset+language[92][lg])
	}
}

// The unlock function attempts to unlock the user using the Unlock-ADAccount PowerShell function.
func unlock() {
	for _, s := range lockedDCs {
		ldapMultiConnect(s)
		powerShellEXE("Unlock-ADAccount -Server " + s + fqdn + " -Identity " + userName)
	}
}

// The testConnection function will test the connection to a computer.
func testConnection() {
	var pingTest string
	var pingGroup = [3]int{}
	var pingResult, avgSpeed int

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

// The restartService function will restart a named service or services.
func restartService(serviceName string) {
	powerShellEXE("Restart-Service -Name " + serviceName)
}

// The stopService function will stop a named service or services.
func stopService(serviceName string) {
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

// The admLocationPrompt function returns a number corresponding to the chosen province.
func admLocationPrompt() int {
	//clear()
	locationTitle()
	regionMenu()
	return intFromString(getInput("Select your location: "))
}

// The atPrompt function displays the computer connected to the application.
func atPrompt() {
	clear()
	fmt.Print(" " + language[15][lg] + ":") // Advanced tools are currently linked to
	fmt.Println(fgGreen, computerName)
	fmt.Println()
}

// The usPrompt function loads the User Search dialog and prompts for a name to search.
func usPrompt() {
	clear()
	usTitle()
	fmt.Println("\n [ex. beo, beordma, beo*dma...]")
	userName = getInput(language[65][lg])
}

// The csPrompt function loads the Computer Search dialog and prompts for a name to search.
func csPrompt() {
	clear()
	csTitle()
	fmt.Println("\n [ex. ncdec652445, 652445, 65244*, 10.141.12.58, ncdec652445.cfia-acia.inspection.gc.ca...]")
	computerName = getInput(language[71][lg])
}

// The psPrompt function loads the Printer Search dialog and prompts for a name to search.
func psPrompt() {
	clear()
	psTitle()
	fmt.Println("\n [ex. P141022202, 10.136.52.188...]")
	printerName = getInput(language[75][lg])
}

// The gsPrompt function loads the Group Search dialog and prompts for a name to search.
func gsPrompt() {
	clear()
	gsTitle()
	groupName = getInput(language[78][lg])
}

// The timezone function returns the current difference from UTC.
func timezone() string {
	return strings.TrimSpace(powerShellRVS("Get-TimeZone | Select -exp BaseUtcOffset | Select -exp Hours"))
}

// The intFromString function converts a String value to an Integer.
func intFromString(convertable string) int {
	freshInt, _ := strconv.Atoi(convertable)
	return freshInt
}

// The convertLargeInt function converts a Large Integer to a String value.
func convertLargeInt(value string) string {
	return powerShellRVS("(Get-Date 1/1/1601).AddDays(" + value + "/864000000000).AddHours(" + timezone() + ")")
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
