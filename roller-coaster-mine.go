/*
Here are some additional details:
• Passengers should invoke board and unboard.
• The car should invoke load, run and unload.
• Passengers cannot board until the car has invoked load
• The car cannot depart until C passengers have boarded.
• Passengers cannot unboard until the car has invoked unload.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/deckarep/golang-set"
)

type Passenger struct {
	id int
}

var C = 10
var n = 15

var car = mapset.NewSet()
var boarding = make(chan Passenger, C)
var unboarding = make(chan Passenger, C)
var lock = sync.Mutex{}

func train() {
	for true {
		fmt.Println("Train loading")
		for car.Cardinality() < C {
			p := <-boarding
			car.Add(p)
		}
		fmt.Println("Train has departed")
		r := rand.Intn(1000)
		time.Sleep(time.Duration(r) * time.Millisecond)

		fmt.Println("Train unloading")
		for car.Cardinality() > 0 {
			p := car.Pop().(Passenger)
			unboarding <- p
		}
	}
}

func passengers() {
	i := 0
	for true {
		select {
		case <-unboarding:
		default:
			p := Passenger{i}
			boarding <- p
			i += 1
		}
	}
}

func main() {
	go train()
	go passengers()
	select {}
}
