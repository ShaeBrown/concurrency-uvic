package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var forks [5]sync.Mutex = [5]sync.Mutex{sync.Mutex{},
	sync.Mutex{}, sync.Mutex{}, sync.Mutex{}}

func eat(philosopher int) {
	var left, right = philosopher, (philosopher + 1) % 5
	var first, second int
	if left < right {
		first = left
		second = right
	} else {
		first = right
		second = left
	}
	forks[first].Lock()
	forks[second].Lock()
}

func think(philosopher int) {
	var left, right = philosopher, (philosopher + 1) % 5
	forks[left].Unlock()
	forks[right].Unlock()
}

func sleep() {
	r := rand.Int(100)
	time.Sleep(time.Duration(r) * time.Millisecond)
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
