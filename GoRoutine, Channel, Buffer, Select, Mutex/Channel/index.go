// channels are mainly used for communications between two different go routines
package main

import "fmt"

func assignValues(v chan string) {
	v <- "name"
}

func main() {
	receivedValue1 := make(chan string)
	// defer close(receivedValue)
	go assignValues(receivedValue1)

	receivedValue2 := make(chan int)

	go func(val chan int) {
		val <- 210
	}(receivedValue2)

	value1 := <-receivedValue1
	value2 := <-receivedValue2

	fmt.Println(value1, value2)
}
