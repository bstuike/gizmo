package tools

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var ps, _ = exec.LookPath("powershell.exe")

// clear function clears the terminal or screen.
// func psClear() {
// 	clearCmd := exec.Command(ps, "/c", "cls")
// 	clearCmd.Run()
// }

// clear function clears the terminal or screen.
func clear() {
	clearCmd := exec.Command("cmd", "/c", "cls")
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

// console function runs standard command prompt commands.
// func console(task string) {
// 	cmd := exec.Command("cmd", "/c", task)

// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	err := cmd.Run()
// 	checkError(err)
// }

// powershellRV function runs a PowerShell command and returns the output as a String.
func powershellRVS(task string) string {
	psCmd := exec.Command(ps, "/c", task)

	psOut, _ := psCmd.Output()
	//checkError(err)
	return string(psOut)
}

// powershellRV function runs a PowerShell command and returns the output as an Integer.
func powershellRVI(task string) int {
	psCmd := exec.Command(ps, "/c", task)

	psOut, _ := psCmd.Output()
	byteToInt, _ := strconv.Atoi(string(psOut))
	return byteToInt
}

// powershellEX function executes a PowerShell command directly.
func powershellEX(task string) {
	psCmd := exec.Command(ps, "/c", task)

	psCmd.Stdout = os.Stdout
	psCmd.Stderr = os.Stderr

	err := psCmd.Run()
	checkError(err)
}

// checkError function executes the builtin panic function if an error is detected.
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// disableCard function will disable a network card on a remote computer.
func disableCard() {
	adapters := powershellRVS("Get-NetAdapter -Name * -IncludeHidden | Format-List -Property deviceid,name")
	fmt.Print(adapters)
	fmt.Print(" " + language[154][lg] + ": ")
	deviceID, _ := reader.ReadString('\n')
	deviceID = strings.Title(strings.Replace(deviceID, "\r\n", "", -1))

	if deviceID == "10" {
		fmt.Println(deviceID)
	} else {
		fmt.Println(deviceID)
	}

	fmt.Println(string(fgRed), language[158][lg], string(colorReset))
	fmt.Println("\n Press the Enter key to continue")
	fmt.Scanln()
}

// testConnection function will test the connection to a computer.
func testConnection() {
	atPrompt()
	fmt.Println("\n"+string(fgYellow), language[150][lg], string(fgWhite), pc)      // Checking connection to
	fmt.Println(string(fgGreen), language[151][lg])                                 //Connection succeeded!
	fmt.Println(string(fgRed), language[152][lg])                                   //Connection failed!
	fmt.Println("\n"+string(fgYellow), language[153][lg]+"...", string(colorReset)) // Testing speed
	powershellEX("ping -a " + pc)
	fmt.Println("\n Press the Enter key to continue")
	fmt.Scanln()
}

// reboot function will reboot a remote computer.
func reboot() {
	powershellEX("Restart-Computer -ComputerName " + pc + " -Force")
}

// logoff function will force a logoff.
func logoff() {
	powershellEX("Restart-Computer -ComputerName " + pc + " -Force")
}

func atPrompt() {
	clear()
	fmt.Print(" " + language[15][lg] + ":") // Advanced tools are currently linked to
	fmt.Println(string(fgGreen), pc)
}
