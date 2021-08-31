package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func parseArgs(c []string) (float64, float64) {
	num1, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		panic(err)
	}
	num2, err := strconv.ParseFloat(c[2], 64)
	if err != nil {
		panic(err)
	}
	return num1, num2
}

func Calculate(e []string) []float64 {
	var result []float64 = nil
	for _, v := range e {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		c := strings.Split(v, " ")
		if len(c)-1 < 2 {
			panic("error: some arguments are not supplied. (Ex: 50 - 20)")
		}
		num1, num2 := parseArgs(c)
		switch c[1] {
		case "+":
			result = append(result, num1+num2)
		case "-":
			result = append(result, num1-num2)
		case "*":
			result = append(result, num1*num2)
		case "/":
			if num2 == 0.0 {
				panic("error: you tried to divide by zero.")
			}
			result = append(result, num1/num2)
		default:
			panic("error: no allocated operation")
		}
	}
	return result
}
