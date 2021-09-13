package main

/*
* Create a channel of int and then create a goroutine to add a value to the channel
* and then print the channel value in the main function
 */
import "fmt"

func print(value int, ch chan int) {
	ch <- value
}

func main() {
	ch := make(chan int)
	go print(10, ch)
	fmt.Println(<-ch)
}
