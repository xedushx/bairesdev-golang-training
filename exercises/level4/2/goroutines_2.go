package main

/*
* Create a function goroutine that will execute an anonymous function to just print the number “1”,
* in the main function print the number “0” and also add a time.Sleep() to wait 2 seconds
 */
import (
	"fmt"
	"time"
)

func printMessage() {
	fmt.Println("1")
}

func main() {
	fmt.Println("0")
	go printMessage()
	time.Sleep(2 * time.Second)
}
