package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("Calculator 1.0")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter numbers to add: ")
	snum1, _ := reader.ReadString('\n')
	operator, _ := reader.ReadString('\n')
	snum2, _ := reader.ReadString('\n')

	num1, err1 := strconv.Atoi(strings.Trim(snum1, "\n"))
	if err1 != nil {
		fmt.Println(err1.Error())
		panic(err1)
	}

	operator = strings.Trim(operator, "\n")

	num2, err2 := strconv.Atoi(strings.Trim(snum2, "\n"))
	if err2 != nil {
		fmt.Println(err2.Error())
		panic(err2)
	}

	result := calculate(num1, operator, num2)
	fmt.Print(result)
}

/*
https://tour.golang.org/flowcontrol/9
Based on the operator variable (+, -, *, /) return the calculated value.
*/
func calculate(x int, operator string, y int) int {
	switch operator {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	case "%":
		return x % y
	}

	return 0
}
