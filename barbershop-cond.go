package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var barber = sync.NewCond(&sync.Mutex{})
var customers = 0
var max_customers = 10
var quit = make(chan struct{})
var total_customers = 0

func take_time() {
	r := rand.Intn(1000)
	time.Sleep(time.Duration(r) * time.Millisecond)
}

func barbers() {
	for true {
		barber.L.Lock()
		if customers == 0 {
			fmt.Println("Going to sleep, no customers")
			barber.Wait()
		}
		barber.L.Unlock()
		time.Sleep(time.Duration(1000) * time.Millisecond)
		barber.L.Lock()
		customers -= 1
		total_customers += 1
		if total_customers == 100 {
			quit <- struct{}{}
		}
		fmt.Printf("Customer left. There is now %d customers\n", customers)
		barber.L.Unlock()
	}
}

func customer() {
	for true {
		time.Sleep(time.Duration(800) * time.Millisecond)
		barber.L.Lock()
		if customers == max_customers {
			barber.L.Unlock()
			fmt.Println("Customer left because barbershop is full")
			continue
		}
		customers += 1
		fmt.Printf("Customer came in. There is now %d customers\n", customers)
		barber.L.Unlock()
		barber.Broadcast()
	}
}

func main() {
	go customer()
	go barbers()
	select {
	case <-quit:
		return
	}
}
