// mimic waitgroup functionality using channel

package main

import "fmt"

func task(done chan bool) {
	defer func() { done <- true }() // executes even if processing fails inside

	fmt.Println("Non-blocking function: Something processing inside function")
}

func main() {
	done := make(chan bool)

	go task(done) // running a non-blocking func

	fmt.Println("Main go-routine: Normally this shouldn't have printed.")
	fmt.Println("Main go-routine:  But here it will be printed because of next line in code.")

	<-done // waiting for chan reply making it blocking

}
