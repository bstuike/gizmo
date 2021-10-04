package QuickTools

import (
	"fmt"
)

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
	fmt.Print("\n " + language[5][lg])             // Please make a selection:
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
	fmt.Print("\n " + language[5][lg])            // Please make a selection:
}

// serviceMenu function displays the Service Tools menu.
func serviceMenu(lg int) {
	clear()
	subTitle(24)
	fmt.Println("\n     1 - " + language[39][lg]) // Get service list
	fmt.Println("     2 - " + language[40][lg])   // Start service(s)
	fmt.Println("     3 - " + language[41][lg])   // Restart service(s)
	fmt.Println("     4 - " + language[42][lg])   // Stop Service(s)
	fmt.Println("\n     8 - " + language[13][lg]) // Back
	fmt.Println("     9 - " + language[14][lg])   // Exit
	fmt.Print("\n " + language[30][lg])           // Choose an operation:
}

// processMenu function displays the Process Tools menu.
func processMenu(lg int) {
	clear()
	subTitle(23)
	fmt.Println("\n     1 - " + language[28][lg]) // Get process list
	fmt.Println("     2 - " + language[29][lg])   // Terminate process
	fmt.Println("\n     8 - " + language[13][lg]) // Back
	fmt.Println("     9 - " + language[14][lg])   // Exit
	fmt.Print("\n " + language[30][lg])           // Choose an operation:
}

// advancedTitle surrounds the Advanced Tools title with yellow stars.
func advancedTitle() {
	fmt.Println("\n"+fgYellow, "**********************")
	fmt.Println(" **", fgWhite, language[18][lg], fgYellow, "**")
	fmt.Println(" **********************", colorReset)
}

// mainTitle surrounds the Main Menu title with yellow stars.
func mainTitle() {
	if lg == 0 {
		fmt.Println(fgYellow, "*****************")
		fmt.Println(" **", fgWhite, language[175][lg], fgYellow, "**")
		fmt.Println(" *****************", colorReset)
	} else {
		fmt.Println(fgYellow, "**********************")
		fmt.Println(" **", fgWhite, language[175][lg], fgYellow, "**")
		fmt.Println(" **********************", colorReset)
	}
}

// subTitle surrounds the Service ot Process Tools title with yellow stars.
func subTitle(num int) {
	fmt.Println()
	if lg == 0 {
		fmt.Println(fgYellow, "*********************")
		fmt.Println(" **", fgWhite, language[num][lg], fgYellow, "**")
		fmt.Println(" *********************", colorReset)
	} else if lg == 1 && num == 23 {
		fmt.Println(fgYellow, "*********************************")
		fmt.Println(" **", fgWhite, language[num][lg], fgYellow, "**")
		fmt.Println(" *********************************", colorReset)
	} else {
		fmt.Println(fgYellow, "********************************")
		fmt.Println(" **", fgWhite, language[num][lg], fgYellow, "**")
		fmt.Println(" ********************************", colorReset)
	}
}
