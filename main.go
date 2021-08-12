/*
main:
- Go version: 1.16.7
- Author: Byron Stuike
- Date: 2021-08-11
*/

// go run main.go

package main

import (
	"fmt"

	T "tools/tools"
)

// main function lauches the program and executes the main selection of program abilities.
func main() {
	T.OSLanguage()
	//T.TestDomain()
	T.MainTasks()
	goodbye()
}

// goodbye function prints a farwell messsage to an exiting user.
func goodbye() {
	fmt.Println("\n Thank you for using the Quick Tools System")
}
