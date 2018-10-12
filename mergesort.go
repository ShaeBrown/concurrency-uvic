package main

import (
	"fmt"
	"math/rand"
)

func merge(a []int, b []int) []int {
	var r = make([]int, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}

func Mergesort(items []int) chan []int {
	out := make(chan []int)
	go func() {
		if len(items) < 2 {
			out <- items
			return
		}

		var middle = len(items) / 2

		left := Mergesort(items[:middle])
		right := Mergesort(items[middle:])

		a := <-left
		b := <-right

		out <- merge(a, b)
	}()
	return out
}

func randIntArray(length int) []int {
	input := make([]int, length)
	for i := 0; i < length; i += 1 {
		input[i] = rand.Int()
	}
	return input
}

func main() {
	input := randIntArray(10000)
	result := Mergesort(input)
	solution := <-result
	fmt.Println(solution)
}
