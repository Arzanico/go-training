package main

import (
	"fmt"
	"sync"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(race())
	}
}

func race() int {
	var wg sync.WaitGroup
	x := 0

	for i := 0; i < 1000; i++ {
		wg.Add(2)

		go func() {
			x = 1
			wg.Done()
		}()

		go func() {
			x = 2
			wg.Done()
		}()
	}

	wg.Wait()
	return x
}
