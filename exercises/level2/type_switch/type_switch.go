package main

import "fmt"

type User struct {
	Name string
}

func checkType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("Type: int, Value:", a.(int))
	case string:
		fmt.Println("Type: string, Value:", a.(string))
	case float64:
		fmt.Println("Type: float64, Value:", a.(float64))
	case User:
		fmt.Println("Type: User, Value:", a.(User))
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	checkType("Eduardo")
	checkType(32)
	checkType(User{"Eduardo Pacheco"})
	checkType(11.25)
	checkType(false)
}
