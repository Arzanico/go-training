package main

import (
	"fmt"
	"sync"
)

// --- Chopsticks ---

type chopStick struct{ sync.Mutex }

// --- Philosopher ---

type Philo struct {
	n               int
	leftCs, rightCs *chopStick
}

// --- Host (runs in its own goroutine) ---

type request struct {
	reply chan struct{} // host closes/sends to grant permission
}

type Host struct {
	reqCh  chan request
	doneCh chan struct{}
}

func NewHost(maxConcurrent int) *Host {
	h := &Host{
		reqCh:  make(chan request),
		doneCh: make(chan struct{}),
	}
	go h.run(maxConcurrent)
	return h
}

func (h *Host) run(maxConcurrent int) {
	eating := 0
	var queue []request

	for {
		// If there is capacity and someone is queued, grant immediately.
		if eating < maxConcurrent && len(queue) > 0 {
			r := queue[0]
			queue = queue[1:]
			eating++
			// Grant permission
			r.reply <- struct{}{}
			close(r.reply)
			continue
		}

		select {
		case r := <-h.reqCh:
			if eating < maxConcurrent {
				eating++
				r.reply <- struct{}{}
				close(r.reply)
			} else {
				// queue the request until someone finishes
				queue = append(queue, r)
			}

		case <-h.doneCh:
			// One philosopher finished eating
			if eating > 0 {
				eating--
			}
		}
	}
}

func (h *Host) Acquire() {
	r := request{reply: make(chan struct{})}
	h.reqCh <- r
	<-r.reply // wait until host grants permission
}

func (h *Host) Release() {
	h.doneCh <- struct{}{}
}

// --- Eating logic ---

func (p Philo) eat(host *Host, wg *sync.WaitGroup) {
	defer wg.Done()

	for meal := 0; meal < 3; meal++ {
		// Ask permission from host (host goroutine enforces max 2 concurrent)
		host.Acquire()

		// Pick up chopsticks in any order (not lowest-numbered first).
		// To keep it "any order", we alternate by philosopher number.
		if p.n%2 == 0 {
			p.rightCs.Lock()
			p.leftCs.Lock()
		} else {
			p.leftCs.Lock()
			p.rightCs.Lock()
		}

		// Print after obtaining necessary locks
		fmt.Printf("starting to eat %d\n", p.n)

		// (No sleep needed by the spec; keeping it minimal/deterministic.)

		// Print before releasing locks
		fmt.Printf("finishing eating %d\n", p.n)

		// Put down chopsticks
		p.rightCs.Unlock()
		p.leftCs.Unlock()

		// Tell host we're done
		host.Release()
	}
}

func main() {
	// Create 5 chopsticks
	cSticks := make([]*chopStick, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(chopStick)
	}

	// Host allows at most 2 philosophers to eat concurrently (in its own goroutine)
	host := NewHost(2)

	// Create 5 philosophers (1..5)
	philos := make([]Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = Philo{
			n:       i + 1,
			leftCs:  cSticks[i],
			rightCs: cSticks[(i+1)%5],
		}
	}

	// Run philosophers
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(host, &wg)
	}
	wg.Wait()
}
