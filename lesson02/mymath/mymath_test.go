package mymath_test

import (
	"github.com/maolenlahe/lessons/lesson02/mymath"
	"testing"
)

func TestCalculate(t *testing.T) {
	v, _ := mymath.Calculate(3, "*", 6)
	if v != 18 {
		t.Error("Expected 18, got ", v)
	}
}

func TestAverage(t *testing.T) {
	v := mymath.Average(3, 4, 5, 6, 7)
	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
}
