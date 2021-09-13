package main

/*
* Create a goroutine that will execute an anonymous function to print “Hello World”
* and in the main routine print “main function”
 */
import "fmt"

func printMessage() {
	fmt.Println("Hello World")
}

func main() {
	fmt.Println("main function")
	go printMessage()
}
