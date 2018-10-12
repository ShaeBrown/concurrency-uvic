package main

import (
	"fmt"
	"math/rand"
	"time"
)

var queue = make(chan struct{}, 2)

func take_time() {
	r := rand.Intn(1000)
	time.Sleep(time.Duration(r) * time.Millisecond)
}

func barber() {
	for true {
		<-queue
		fmt.Printf("Customer left. There is now %d customers\n", len(queue))
		take_time()
	}
}

func customers() {
	for true {
		take_time()
		select {
		case queue <- struct{}{}: // could make a customer class, and pass the info
			fmt.Printf("Customer came in. There is now %d customers\n", len(queue))
		default:
			fmt.Println("Customer left because barbershop is full")
		}
	}
}

func main() {
	go customers()
	go barber()
	select {}
}
