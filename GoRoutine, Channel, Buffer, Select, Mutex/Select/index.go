package main

import (
	"fmt"
	"time"
)

func main() {
	// select keyword allows us to wait for multiple communication operations of goroutines
	firstChannel := make(chan string)
	secondChannel := make(chan string)
	quit := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		firstChannel <- "Nice"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		secondChannel <- "Hello"
	}()

	// using select keyword to wait on the receiving end of these two channels
	for i := 0; i < 2; i++ {
		select {
		case receivingEndFirst := <-firstChannel:
			fmt.Println("Received from first channel", receivingEndFirst)
		case receivingEndSecond := <-secondChannel:
			fmt.Println("Received from second channel", receivingEndSecond)
		case <-quit:
			return
		}

	}

}
