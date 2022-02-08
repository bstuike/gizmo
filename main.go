/*
main:
- Author: Byron Stuike
- Date: 2021-08-11
*/

/*
go run main.go
*/

package main

import (
	Q "gizmo/gizmo"
)

// The main function launches the program and executes the main selection of program abilities.
func main() {
	// Q.TestDomain()
	Q.LDAPConnect()
	Q.MainChoices()
}
