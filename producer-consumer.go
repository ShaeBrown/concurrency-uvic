package main

import "fmt"

var quit = make(chan string)

func consume(buffer chan int) {
	for num := range buffer {
		fmt.Println("Consumed ", num)
		if num == 9999 {
			quit <- "done"
		}
	}
}

func produce(buffer chan int) {
	num := 0
	for num < 10000 {
		buffer <- num
		num = num + 1
	}
}

func main() {
	buffer_size := 10
	ch := make(chan int, buffer_size)
	go produce(ch)
	go consume(ch)
	select {
	case <-quit:
		return
	} // run forever
}
