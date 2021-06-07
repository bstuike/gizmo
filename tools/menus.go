package tools

import (
	"fmt"
)

// welcome function displays the program name and author information
func welcome(lg int) {
	clear()
	fmt.Print(" F - Français")
	fmt.Println("\tE - English")
	fmt.Println("\n" + string(bgRed) + "                                           ")
	fmt.Println(" ----------------------------------------- ")
	fmt.Println(" ----------- Quick Tools GIZMO ----------- ")
	fmt.Println(" ----------------------------------------- ")
	fmt.Println("                                          ", string(colorReset))
	fmt.Println(string(fgBrightBlue) + "\n " + language[174][lg] + ":" + string(fgWhite) + " Marc-Antoine Beord (marc-antoine.beord@ssc-spc.gc.ca)")
	fmt.Println(string(fgBrightCyan), language[1][lg]+":"+string(fgWhite)+" Byron Stuike (byron.stuike@inspection.gc.ca)")
	fmt.Println(string(fgBrightYellow), language[2][lg]+":"+string(fgWhite)+" Byron Stuike (byron.stuike@inspection.gc.ca)")
	fmt.Println("\n " + language[4][lg] + string(fgGreen) + " Username")
	fmt.Println()
}

// mainMenu function displays the complete list of initial options.
func mainMenu(lg int) {
	mainTitle()
	fmt.Println("\n     0 - " + language[160][lg]) // ORCA Status Verification
	fmt.Println("     1 - " + language[6][lg])     // RESET Password
	fmt.Println("     2 - " + language[7][lg])     // LOCKED OUT account
	fmt.Println("\n     3 - " + language[8][lg])   // USER Information
	fmt.Println("     4 - " + language[9][lg])     // COMPUTER Information
	fmt.Println("     5 - " + language[10][lg])    // PRINTER Information
	fmt.Println("     6 - " + language[11][lg])    // GROUP Information
	fmt.Println("\n     7 - " + language[12][lg])  // Advanced Computer Tools
	fmt.Println("\n     9 - " + language[14][lg])  // Exit
	fmt.Print("\n " + language[5][lg])             // Please make a selection
}

// advancedMenu function displays the Advanced Tools menu.
func advancedMenu(lg int) {
	advancedTitle()
	fmt.Println("\n     1 - " + language[19][lg]) // Force Logoff
	fmt.Println("     2 - " + language[20][lg])   // Restart Computer
	fmt.Println("     3 - " + language[21][lg])   // Test Network Connection
	fmt.Println("\n     4 - " + language[22][lg]) // Disable Network Card
	fmt.Println("     5 - " + language[23][lg])   // Process Tools
	fmt.Println("     6 - " + language[24][lg])   // Service Tools
	fmt.Println("\n     8 - " + language[13][lg]) // Back
	fmt.Println("     9 - " + language[14][lg])   // Exit
	fmt.Print("\n " + language[5][lg])            // Please make a selection
}

// serviceMenu function displays the Service Tools menu
func serviceMenu(lg int) {
	clear()
	subTitle(24)
	fmt.Println("\n     1 - " + language[39][lg]) // Get service list
	fmt.Println("     2 - " + language[40][lg])   // Start service(s)
	fmt.Println("     3 - " + language[41][lg])   // Restart service(s)
	fmt.Println("     4 - " + language[42][lg])   // Stop Service(s)
	fmt.Println("\n     8 - " + language[13][lg]) // Back
	fmt.Println("     9 - " + language[14][lg])   // Exit
	fmt.Print("\n " + language[30][lg])           // Please make a selection
}

// processMenu function displays the Process Tools menu.
func processMenu(lg int) {
	clear()
	subTitle(23)
	fmt.Println("\n     1 - " + language[28][lg]) // Get process list
	fmt.Println("     2 - " + language[29][lg])   // Terminate process
	fmt.Println("\n     8 - " + language[13][lg]) // Back
	fmt.Println("     9 - " + language[14][lg])   // Exit
	fmt.Print("\n " + language[30][lg])           // Please make a selection
}

func advancedTitle() {
	fmt.Println("\n"+string(fgYellow), "**********************")
	fmt.Println(" **", string(fgWhite), language[18][lg], string(fgYellow), "**")
	fmt.Println(" **********************", string(colorReset))
}

func mainTitle() {
	if lg == 0 {
		fmt.Println(string(fgYellow), "*****************")
		fmt.Println(" **", string(fgWhite), language[175][lg], string(fgYellow), "**")
		fmt.Println(" *****************", string(colorReset))
	} else {
		fmt.Println(string(fgYellow), "**********************")
		fmt.Println(" **", string(fgWhite), language[175][lg], string(fgYellow), "**")
		fmt.Println(" **********************", string(colorReset))
	}
}

func subTitle(num int) {
	if lg == 0 {
		fmt.Println(string(fgYellow), "*********************")
		fmt.Println(" **", string(fgWhite), language[num][lg], string(fgYellow), "**")
		fmt.Println(" *********************", string(colorReset))
	} else if lg == 1 && num == 23 {
		fmt.Println(string(fgYellow), "*********************************")
		fmt.Println(" **", string(fgWhite), language[num][lg], string(fgYellow), "**")
		fmt.Println(" *********************************", string(colorReset))
	} else {
		fmt.Println(string(fgYellow), "********************************")
		fmt.Println(" **", string(fgWhite), language[num][lg], string(fgYellow), "**")
		fmt.Println(" ********************************", string(colorReset))
	}
}
