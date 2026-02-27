package main

import (
	"fmt"
	"sync"
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	id              int
	leftCS, rightCS *Chopstick
}

func (p Philosopher) eat(wg *sync.WaitGroup, host chan bool) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		fmt.Printf("finishing eating %d\n", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<-host
	}
}

func main() {
	host := make(chan bool, 2)
	var wg sync.WaitGroup

	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:      i + 1,
			leftCS:  chopsticks[i],
			rightCS: chopsticks[(i+1)%5],
		}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(&wg, host)
	}
	wg.Wait()
}
