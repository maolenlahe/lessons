package mymath

import (
	"errors"
	"fmt"
	"sort"
)

/*
https://tour.golang.org/flowcontrol/9
Based on the operator variable (+, -, *, /) return the calculated value.
*/
func Calculate(x int, operator string, y int) (int, error) {
	switch operator {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		return x / y, nil
	case "%":
		return x % y, nil
	}

	return 0, errors.New("operator: (+, -, *, /, %)")
}

// @todo Modify this function to accept another parameter "mean", "median" or "mode" * Write your tests first
func Average(args ...int) int {
	sort.Ints(args)
	var sum int = 0
	for _, value := range args {
		sum += value
	}
	return sum / len(args)
}
