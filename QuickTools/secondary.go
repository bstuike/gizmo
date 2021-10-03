package QuickTools

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var ps, _ = exec.LookPath("powershell.exe")

// clear function clears the terminal or screen.
func clear() {
	fmt.Println(colorReset)
	clearCmd := exec.Command("cmd", "/c", "cls")
	clearCmd.Stdout = os.Stdout

	err := clearCmd.Run()
	checkError(err)
}

// atPrompt function displays the computer connected to the application.
func atPrompt() {
	clear()
	fmt.Println()
	fmt.Print(" " + language[15][lg] + ":") // Advanced tools are currently linked to
	fmt.Println(fgGreen, localPC())
}

// localPC function gets the name of the local computer.
func localPC() string {
	pc, err := os.Hostname()
	checkError(err)
	return pc
}

func cliu() string {
	person, err := user.Current()
	checkError(err)
	return person.Name
}

func carryon() {
	fmt.Println("\nPress"+string(fgCyan), "Enter"+string(colorReset), "to continue")
	fmt.Scanln()
}

// powershellEXE function executes a PowerShell command directly.
func powerShellEXE(task string) {
	psCmd := exec.Command(ps, task)

	psCmd.Stdout = os.Stdout
	psCmd.Stderr = os.Stderr

	err := psCmd.Run()
	checkError(err)
}

// powershellRVS function runs a PowerShell command and returns the output as a String.
func powerShellRVS(task string) string {
	psCmd := exec.Command(ps, task)
	psOut, _ := psCmd.Output()
	return string(psOut)
}

// checkError function executes the builtin panic function if an error is detected.
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
// console function runs standard command prompt commands.
func console(task string) {
	cmd := exec.Command("cmd", "/c", task)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	checkError(err)
}
*/

/*
// powershellRVI function runs a PowerShell command and returns the output as an Integer.
func powershellRVI(task string) int {
	psCmd := exec.Command(ps, "/c", task)

	psOut, _ := psCmd.Output()
	byteToInt, _ := strconv.Atoi(string(psOut))
	return byteToInt
}
*/

// testConnection function will test the connection to a computer.
func testConnection() {
	atPrompt()
	fmt.Println("\n"+fgYellow, language[150][lg], fgWhite, localPC()) // Checking connection to
	fmt.Println(fgGreen, language[151][lg])                           // Connection succeeded!
	fmt.Println(fgRed, language[152][lg])                             // Connection failed!
	fmt.Println("\n"+fgYellow, language[153][lg]+"...", colorReset)   // Testing speed
	powerShellEXE("ping -a " + localPC())
	carryon()
}

// disableCard function will disable a network card on a remote computer.
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
	carryon()
}

// reboot function will reboot a remote computer.
func reboot() {
	powerShellEXE("Restart-Computer -ComputerName " + localPC() + " -Force")
}

// logoff function will force a logoff.
func logoff() {
	powerShellEXE("shutdown /l /m \\" + localPC() + " /t 0")
}
