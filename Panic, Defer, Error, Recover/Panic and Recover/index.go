package main

import (
	"fmt"
)

func main() {
	// defer cleanUp()
	// a, b := 10, 0
	// finalValue := a / b
	// fmt.Println(finalValue)

	// defer cleanUp()
	// str := "abc"
	// n, err := strconv.ParseInt(str, 2, 5)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(n)
	fmt.Println("start")
	startPanicking()
	fmt.Println("end")

}

func startPanicking() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	panic("Brace yourself, panic started")
	//wont exceute this line because of panic func
	fmt.Println("after panicking completed")
}

// recover helps to revoer from panic
// func cleanUp() {
// 	if r := recover(); r != nil {
// 		fmt.Println("recovered", r)
// 	}
// }
