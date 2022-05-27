package main

import "fmt"

func main() {
	fmt.Println("Line first")

	// defer statements delay the execution of the function or method or an
	// anonymous method until the nearby functions returns
	// if two defer occur then last one to go will be executed at first (LIFO)

	// defer fmt.Println("Line second")
	// fmt.Println("Line third")

	// defer fmt.Println("Line Fourth")
	multiply(1, 10)
	multiply(2, 10)
	defer multiply(3, 10)
	multiply(4, 10)
	multiply(5, 10)
	multiply(6, 10)
	multiply(7, 10)
	multiply(8, 10)
	defer multiply(9, 10)
	multiply(10, 10)

}

func multiply(a int, b int) {
	mul := a * b
	fmt.Println("the multiplication of ", a, "and ", b, "is", mul)
}
