package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func delayFunc(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Execution of goroutine has been completed")
	wg.Done()
}

func waitGroupExample() {
	fmt.Println("Waitgroup")
	// block the part of the code until a go routine is successfully executed
	// here we wont see console of go func because our main function has been completed before the completion of goroutine
	// go delayFunc()
	// this we where waitgroup comes into rescue
	var wg sync.WaitGroup
	wg.Add(1)
	go delayFunc(&wg)
	wg.Wait()
	fmt.Println("Finished executing")
}

var availUrls = []string{"https://google.com", "https://instagram.com", "https://twitter.com", "https://tiktok.com"}

func fetchStatus(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	for _, url := range availUrls {
		wg.Add(1)
		go func(url string) {
			response, err := http.Get(url)
			if err != nil {
				fmt.Println(w, "Error is", err)
			}
			fmt.Println(w, "Response", response.Status, response.StatusCode)
		}(url)
	}
	wg.Wait()
}

func main() {
	// waitGroupExample()fmt.Println(w, "Response", response)
	fmt.Println("I am above the goroutine response")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
