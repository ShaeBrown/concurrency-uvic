package main

import (
	"fmt"
	"sync"
)

var M = 10
var num_savages = 12
var servings = 0

type Semaphore chan struct{}

func (sem Semaphore) Wait() {
	<-sem
}

func (sem Semaphore) Signal() {
	sem <- struct{}{}
}

var empty = make(Semaphore)
var full = make(Semaphore)
var mutex = sync.Mutex{}

func cook() {
	for true {
		empty.Wait()
		fmt.Println("Filling pot")
		full.Signal()
	}
}

func savage(i int) {
	for true {
		mutex.Lock()
		if servings == 0 {
			empty.Signal()
			full.Wait()
			servings = M
		}
		servings -= 1
		fmt.Printf("Savage %d ate from the pot\n", i)
		mutex.Unlock()
	}

}

func main() {
	go cook()
	for i := 0; i < num_savages; i++ {
		go savage(i)
	}
	select {}
}
