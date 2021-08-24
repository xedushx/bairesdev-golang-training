package main

import (
	"bufio"
	ageFilter "com/bairesdev/training/packer/exercises/age_filter"
	"com/bairesdev/training/packer/exercises/calculator"
	"com/bairesdev/training/packer/exercises/inventory"
	"fmt"
	"os"
)

func enqueue(queue []int, element int) []int {
	queue = append(queue, element)
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []int) []int {
	element := queue[0]
	fmt.Println("Dequeued:", element)
	return queue[1:]
}

func main() {

	// 1.
	fmt.Println("********** 1. ************")
	fmt.Println(ageFilter.AgeFilter(5, 20, []int{2, 3, 5, 7, 11, 13}))

	// 3.
	fmt.Println("\n\n********** 3. ************")
	p1 := inventory.NewProduct(1, "computer")
	p2 := inventory.NewProduct(2, "mouse")
	p3 := inventory.NewProduct(3, "tv")
	p4 := inventory.NewProduct(3, "keyboard")
	p5 := inventory.NewProduct(0, "computer")

	localInventory := inventory.Inventory{}

	localInventory.AddItem("A1001", p1)
	localInventory.AddItem("A2001", p2)
	localInventory.AddItem("A3001", p3)
	localInventory.AddItem("A4001", p4)
	localInventory.AddItem("A5001", p5)

	for k, v := range localInventory.GetProducts() {
		fmt.Printf("Product: %-10v -- %-10v \n", k, v.Name)
	}

	// 4.
	fmt.Println("\n\n********** 4. ************")
	var queue []int

	queue = enqueue(queue, 15)
	queue = enqueue(queue, 12)
	queue = enqueue(queue, 125)

	fmt.Println("Queue:", queue)

	queue = dequeue(queue)
	fmt.Println("Queue:", queue)

	queue = enqueue(queue, 30)
	fmt.Println("Queue:", queue)

	queue = dequeue(queue)
	fmt.Println("Queue:", queue)

	// 2.
	fmt.Println("\n\n********** 2. ************")
	expressions := make([]string, 0)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Calculate => ")
		for scanner.Scan() {
			expressions = append(expressions, scanner.Text())
			res := calculator.Calculate(expressions)
			if res != nil {
				fmt.Printf("RESULT: %v = %v \n", expressions[0], res[0])
			}
			expressions = expressions[:0]
			fmt.Print("Calculate => ")
		}
	}
}
