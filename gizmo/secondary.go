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

var ps, _ = exec.LookPath("powershell")
var serviceName string

// The clear function clears the terminal or screen.
func clear() {
	fmt.Println(colorReset)
	clearCmd := exec.Command("cmd", "/c", "cls")
	clearCmd.Stdout = os.Stdout

	err := clearCmd.Run()
	checkError(err)
}

func getInput(prompt string) string {
	fmt.Print("\n " + prompt)
	userInput, _ := reader.ReadString('\n')
	return userInput
}

// The atPrompt function displays the computer connected to the application.
func atPrompt() {
	clear()
	fmt.Println()
	fmt.Print(" " + language[15][lg] + ":") // Advanced tools are currently linked to
	fmt.Println(fgGreen, localPC())
}

// The localPC function gets the name of the local computer.
func localPC() string {
	pc, err := os.Hostname()
	checkError(err)
	return pc
}

func cliu() string {
	person, err := user.Current()
	checkError(err)
	return person.Username
}

func enterKey() {
	fmt.Println(colorReset, "\nPress", fgCyan+"Enter"+colorReset, "to continue")
	fmt.Scanln()
}

func query() {
	clear()
	// Filters must start and finish with ()!
	searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, filterDN, []string{"sAMAccountName", "name", "displayName", "title", "mail", "department", "description", "st", "lastLogon", "pwdLastSet", "homeDirectory", "userAccountControl", "msDS-UserPasswordExpired", "msDS-UserAccountDisabled"}, []ldap.Control{})

	result, err := l.Search(searchReq)
	if err != nil {
		fmt.Println("Failed to query LDAP: ", err)
		return
	}

	fmt.Println(fgGreen, language[131][lg])

	for _, b := range result.Entries {
		accountCode := b.GetAttributeValue("userAccountControl")
		fmt.Println(accountCode)
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

		fmt.Println(fgBrightYellow, b.GetAttributeValue("displayName"), colorReset)
		fmt.Println("\n\n Logon Name                   :", b.GetAttributeValue("sAMAccountName"))
		fmt.Println(" Full Name                    :", b.GetAttributeValue("name"))
		fmt.Println(" Display Name                 :", b.GetAttributeValue("displayName"))
		fmt.Println(" Title                        :", b.GetAttributeValue("title"))
		fmt.Println(" Email Address                :", b.GetAttributeValue("mail"))
		fmt.Println(" Department                   :", b.GetAttributeValue("department"))
		fmt.Println(" Description                  :", b.GetAttributeValue("description"))
		fmt.Println(" Province                     :", b.GetAttributeValue("st"))
		fmt.Println(" Home Directory               :", b.GetAttributeValue("homeDirectory"))
		fmt.Println(" Last Logon                   :", b.GetAttributeValue("lastLogon"))
		fmt.Println(" Password Last Set            :", strings.TrimSpace(convertLargeInt(b.GetAttributeValue("pwdLastSet"))))
		fmt.Println(" Password Expired             :", pwdexp)
		fmt.Println(" Password Never Expires       :", pwdnvexp)
		fmt.Println(" Account Locked Out           :", acctlo)
		fmt.Println(" Account Disabled             :", acctdis)
	}
}

func convertLargeInt(value string) string {
	time := powerShellRVS("(Get-Date 1/1/1601).AddDays(" + value + "/864000000000).AddHours(-7)")
	return time
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

// powershellRVI function runs a PowerShell command and returns the output as an Integer.
func powerShellRVI(task string) int {
	psCmd := exec.Command(ps, "/c", task)

	psOut, _ := psCmd.Output()
	byteToInt, _ := strconv.Atoi(string(psOut))
	return byteToInt
}

// The checkError function executes the builtin panic function if an error is detected.
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// The testConnection function will test the connection to a computer.
func testConnection() {
	atPrompt()
	fmt.Println("\n", fgYellow, language[150][lg], fgWhite, localPC()) // Checking connection to
	fmt.Println(fgGreen, language[151][lg])                            // Connection succeeded!
	fmt.Println(fgRed, language[152][lg])                              // Connection failed!
	fmt.Println("\n", fgYellow, language[153][lg], "...", colorReset)  // Testing speed
	//responseTime := powerShellRVI("Test-Connection -ComputerName 10.139.18.166 -ErrorAction silentlycontinue | Select-Object -Property responsetime")
	responseTime := powerShellRVS("Test-Connection -ComputerName abcalgc682020p")
	// rtime := powerShellEXE("$AVSpeed = $AVSpeed / $results.count")
	// for {key}, {value} := range {list}
	// for _, time := range results {
	// 	fmt.Println("My animal is:", time)
	// }
	//avgspeed := 4
	fmt.Print(responseTime)
	//powerShellEXE("ping -a " + localPC())
	enterKey()
}

// The disableCard function will disable a network card on a remote computer.
func disableCard() {
	adapters := powerShellRVS("Get-NetAdapter -Name * -IncludeHidden | Format-List -Property deviceid,name")
	fmt.Print(adapters)
	fmt.Print(" " + language[154][lg] + ": ")
	deviceID, _ := reader.ReadString('\n')
	deviceID = strings.Title(strings.Replace(deviceID, "\r\n", "", -1))

	if deviceID == "10" {
		fmt.Println(deviceID)
	} else {
		fmt.Println(deviceID)
	}

	fmt.Println(fgRed, language[158][lg], colorReset)
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
	powerShellEXE("Restart-Computer -ComputerName " + "'" + localPC() + "'" + " -Force")
}

// The logoff function will force a logoff.
func logoff() {
	powerShellEXE("shutdown /l /m \\" + localPC() + " /t 0")
}

/*
func expire() {
	f := powerShellRVS("Get-ADUser -Identity stuikeb -Properties msDS-UserPasswordExpiryTimeComputed | Select -exp {[datetime]::FromFileTime($_.”msDS-UserPasswordExpiryTimeComputed”)}")
	fmt.Println(f)
}
*/
