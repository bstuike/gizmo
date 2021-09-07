/*
main:
- Author: Byron Stuike
- Date: 2021-08-11
*/

// go run main.go

package main

import (
	"fmt"
	"sync"

	T "tools/tools"
)

// goodbye function prints a farwell messsage to an exiting user.
func goodbye(wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("\n Thank you for using the Quick Tools System")
	fmt.Println()
}

// pause function inserts a delay into the completion of a function.
func pause(amount int) {
	wg := new(sync.WaitGroup)
	wg.Add(amount)
	goodbye(wg)
	wg.Wait()
}

// main function lauches the program and executes the main selection of program abilities.
func main() {
	T.OSLanguage()
	//T.TestDomain()
	T.MainTasks()
	pause(1)
}
