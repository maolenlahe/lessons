package main

import (
	"bufio"
	"fmt"
	"github.com/maolenlahe/lessons/lesson02/mymath"
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

	result, err := mymath.Calculate(num1, operator, num2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	arr := []int{1, 2, 3, 4}
	avg := mymath.Average(arr...)
	fmt.Println(avg)
}
