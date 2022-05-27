package main

import (
	"fmt"
	"sync"
)

// mutex is used as a  locking mechanish to ensure that only one goroutine / concurrent process is
// accessing the critical section of code at any point of time

var (
	mutex       sync.Mutex
	totalAmount int
)

func deposit(value int, wg *sync.WaitGroup) {
	// acquiring the mutex using mutex.Lock(),
	// this will enter into critical section and run until its proper excecution
	mutex.Lock()
	fmt.Printf("Deposit %d to account with balance %d\n", value, totalAmount)
	totalAmount += value
	// if this method is not called then it will enter into deadlock and nothing will be executed
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdraw %d to account with balance %d\n", value, totalAmount)
	totalAmount -= value
	mutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Hello World")
	totalAmount = 1000
	var wg sync.WaitGroup
	wg.Add(2)
	go withdraw(200, &wg)
	go deposit(700, &wg)
	wg.Wait()
	fmt.Printf("Available balance is %d\n", totalAmount)
}
