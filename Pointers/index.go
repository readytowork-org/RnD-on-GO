package main

import "fmt"

func updatedValue(str *string) {
	*str = "Value has been updated!"
}

func updatedValueAgain(str string) {
	str = "Value has been updated, again!"
}

func firstMain() {
	x := 7
	// Reference of x, i.e.  the location where 7 is stored
	fmt.Println(&x)
	// e.g. 0x1400012c008

	// not the actual value y is equals to the location of x i.e. reference of x
	y := &x
	fmt.Println(y)

	// dereferencing the value where y is pointing to i.e. modifying the block of value stored in y
	// i.e the block is 7 and actual value is somewhat like 0x1400012c008. So this will update the  value of x
	*y = 8

	fmt.Println(x, y)
}

func secondMain() {
	valueToBeUpdated := "Hello guys,"
	// updatedValue(&valueToBeUpdated)
	updatedValueAgain(valueToBeUpdated)
	fmt.Println(valueToBeUpdated)
}

func main() {
	// firstMain()
	// secondMain()
	var value = "Good"
	var pointer *string = &value
	fmt.Println(pointer)
	fmt.Println(*pointer)
	*pointer = "Go has become bad"
	fmt.Println(value)
}
