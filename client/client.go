package main

import (
	"fmt"
	"net/http"
	"sync"
)

const (
	url          = "http://localhost:8080/hello"
	totalRequest = 10_000
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= totalRequest; i++ {
		wg.Add(1)
		go sendRequest(i, &wg)
	}

	wg.Wait()

	fmt.Println("All requests completed")
}

func sendRequest(requestID int, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Request %d failed: %v\n", requestID, err)
		return
	}

	fmt.Printf("Request %d completed %s\n", requestID, resp.Status)

}
