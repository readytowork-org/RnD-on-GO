package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		reRender("Hello")
		wg.Done()
	}()
	// go reRender("Good Bye")
	wg.Wait()
}
func reRender(item string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Processed", i, item)
	}
}
