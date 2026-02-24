package main

import (
	"fmt"
	"sync"
)

/*
A race condition is a characteristic of a code base of a program on which the condition of two processes of
the program compete trying to change the state of the same resource.
In other words, a race condition appears when two process of a same program tries to modify the
state of the same resource.
Race conditions are also related with interleaving. The order
of execution of the process, when they are cocurrent is not deterministic, so in race condition the final state
of a resource that is access from, at least two cocurrent precess, is unpredictable.
*/

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(race())
	}
}

func race() int {
	var wg sync.WaitGroup
	x := 0

	wg.Add(2)
	go routine1(&wg, &x)
	go routine2(&wg, &x)

	wg.Wait()
	return x
}

func routine1(wg *sync.WaitGroup, x *int) {
	*x = 1
	wg.Done()
}

func routine2(wg *sync.WaitGroup, x *int) {
	*x = 2
	wg.Done()
}
