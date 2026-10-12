package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan []int)
	resCh := make(chan []int)
	go getSlice(ch, &wg)
	go getSqare(ch, resCh, &wg)
	fmt.Print(<-resCh)
}

func getSlice(ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	n := 10
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.IntN(100)
	}
	ch <- slice
	close(ch)
}
func getSqare(ch chan []int, rCh chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := make([]int, 10)
	for slice := range ch {
		for num := range slice {
			res := num * num
			result = append(result, res)
		}
	}
	rCh <- result
	close(rCh)
}

//dz
