package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var forks [5]bool = [5]bool{false, false, false, false, false} // is the fork being used
var m = sync.Mutex{}
var waiter *sync.Cond = sync.NewCond(&m)

func eat(philosopher int) {
	for true {
		var left, right = philosopher, (philosopher + 1) % 5
		waiter.L.Lock()
		if !forks[left] && !forks[right] { // the forks are available
			forks[left], forks[right] = true, true //pick them up
			waiter.L.Unlock()
			break
		} else {
			waiter.Wait()
			waiter.L.Unlock()
		}
	}
}

func think(philosopher int) {
	var left, right = philosopher, (philosopher + 1) % 5
	waiter.L.Lock()
	forks[left], forks[right] = false, false //put forks down
	waiter.L.Unlock()
	waiter.Broadcast()
}

func sleep() {
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Second)
}

func philosopher(i int) {
	for true {
		eat(i)
		fmt.Printf("Philosopher %d eating\n", i)
		sleep()
		think(i)
		fmt.Printf("Philpsopher %d thinking\n", i)
		sleep()
	}
}

func main() {
	for i := 0; i < 5; i++ {
		go philosopher(i)
	}
	select {} // run forever
}
