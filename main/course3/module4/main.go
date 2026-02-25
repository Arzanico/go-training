package main

import (
	"fmt"
	"sync"
)

type chopStick struct {
	sync.Mutex
}

type Philo struct {
	n               int
	leftCs, rightCs *chopStick
}

// Host: semaphore of size 2 (at most 2 philosophers eating concurrently)
type Host struct {
	permits chan struct{}
}

func NewHost(maxConcurrent int) *Host {
	h := &Host{permits: make(chan struct{}, maxConcurrent)}
	// Fill with tokens
	for i := 0; i < maxConcurrent; i++ {
		h.permits <- struct{}{}
	}
	return h
}

func (h *Host) Acquire() { <-h.permits }
func (h *Host) Release() { h.permits <- struct{}{} }

func (p Philo) eat(host *Host, wg *sync.WaitGroup) {
	defer wg.Done()

	for meal := 0; meal < 3; meal++ {
		// Ask permission from host (blocks if already 2 eating)
		host.Acquire()

		// Pick up chopsticks (any order; not lowest-numbered trick)
		p.leftCs.Lock()
		p.rightCs.Lock()

		// Must print after obtaining necessary locks
		fmt.Printf("starting to eat %d\n", p.n)

		// Must print before releasing locks
		fmt.Printf("finishing eating %d\n", p.n)

		// Put down chopsticks
		p.rightCs.Unlock()
		p.leftCs.Unlock()

		// Tell host we're done (free a slot)
		host.Release()
	}
}

func main() {
	// Create chopsticks
	cSticks := make([]*chopStick, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(chopStick)
	}

	// Create host (own goroutine not strictly needed for semaphore,
	// but the host "executes independently" conceptually here)
	host := NewHost(2)

	// Create philosophers and run
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		p := Philo{
			n:       i + 1, // 1..5
			leftCs:  cSticks[i],
			rightCs: cSticks[(i+1)%5],
		}
		go p.eat(host, &wg)
	}

	wg.Wait()
}
