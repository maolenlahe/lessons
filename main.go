package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("Init ran")
}

func main() {
	var x int = 27
	var s string
	s = strconv.Itoa(x)
	ss := "hi"
	fmt.Println("Hello,"+s, "World!", x, ss)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter numbers to add: ")
	snum1, _ := reader.ReadString('\n')
	snum2, _ := reader.ReadString('\n')
	num1, err1 := strconv.Atoi(strings.Trim(snum1, "\n"))
	if err1 != nil {
		fmt.Println(err1.Error())
		panic(err1)
	}

	num2, err2 := strconv.Atoi(strings.Trim(snum2, "\n"))
	if err2 != nil {
		fmt.Println(err2.Error())
		panic(err2)
	}

	sum := add(num1, num2)
	fmt.Print(sum)

	//if text == "hi" {
	//	fmt.Println(text)
	//} else {
	//	fmt.Println("wrong")
	//}

}

func add(x int, y int) int {
	return x + y
}

/*
Based on the operator variable (+, -, *, /) return the calculated value.
*/
func calculate(x int, operator string, y int) {

}
