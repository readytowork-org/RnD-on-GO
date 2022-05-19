package main

import (
	"errors"
	"fmt"
)

func main() {
	total, err := sumFrom(50, 10)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Sum is", total)
	}

	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("result is", result)
	}
}
func sumFrom(start int, end int) (int, error) {

	if start > end {
		return 0, errors.New("Starting value should not be grater than ending value.")
	}

	result := 0

	for i := start; i <= end; i++ {
		result += i
	}
	return result, nil
}
func divide(dividend float64, divisor float64) (float64, error) {

	if divisor == 0 {
		return 0, errors.New("Divisor should not be 0")
	}

	result := dividend / divisor

	return result, nil
}
