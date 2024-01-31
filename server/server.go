package main

import (
	"log"
	"net/http"
	"time"
)

var (
	ch1 = make(chan int)

	workers = 50
)

func workerFxn() {

	for i := range ch1 {
		log.Println(i)

		// imitation of some work
		time.Sleep(time.Second * 5)
	}
}

func handleHello(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	w.Write([]byte("Welcome"))

	// send the req id to ch1 channel
	go func ()  {
		ch1<- time.Now().Nanosecond()
	}()

}

func main() {

	// launch the workers
	for i := 0; i < workers; i++ {
		go workerFxn()

	}

	http.HandleFunc("/hello", handleHello)

	log.Println("Starting server in 8080..")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error running on port 8080")
	}
}
