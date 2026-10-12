package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	resCh := make(chan int)
	go getSlice(ch, &wg)
	go getSqare(ch, resCh, &wg)
	for num := range resCh {
		fmt.Print(num)
	}
}

func getSlice(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	n := 10
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.IntN(101)
	}
	for _, v := range slice {
		ch <- v
	}
	close(ch)
}
func getSqare(ch chan int, rCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		rCh <- v * v
	}

	close(rCh)
}

//dz
