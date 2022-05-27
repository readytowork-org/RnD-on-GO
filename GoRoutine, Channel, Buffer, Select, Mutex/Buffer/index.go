package main

import "fmt"

func main() {
	// this will give error as the channel does not have any capacity to run
	// firstChannel := make(chan string)
	// firstChannel <- "First one"
	// fmt.Println(<-firstChannel)
	// so instead we, specify the capacity

	// channel has only one capacity but we have two messages so this will enter into deadlock
	// as one  slot is already used by first message so we will incraese the slots
	// firstChannel := make(chan string, 1)

	firstChannel := make(chan string, 2)
	firstChannel <- "First one"
	firstChannel <- "Second one"

	// length of the buffered channel is shown as
	fmt.Println(len(firstChannel))
	fmt.Println(<-firstChannel)
	// gives only one output so to print the second one we need to print again
	fmt.Println(<-firstChannel)

	// capacity of the buffered channel is shown as
	fmt.Println(cap(firstChannel))
	// length of the buffered channel is O because both the channels are written
	fmt.Println(len(firstChannel))
}
