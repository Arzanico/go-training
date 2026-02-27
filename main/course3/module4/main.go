package main

import (
	"fmt"
	"sync"
)

type seat struct{}

type chopStick struct {
	sync.Mutex
}

type request struct {
	reply chan struct{}
}

type host struct {
	requestSeatCh chan request
	releaseSeatCh chan seat
}

func newHost() *host {
	h := &host{
		requestSeatCh: make(chan request),
		releaseSeatCh: make(chan seat),
	}

	go h.run(2)
	return h

}

func (o *host) acquire() {
	r := request{
		reply: make(chan struct{}),
	}
	o.requestSeatCh <- r
	<-r.reply
}

func (o *host) release() {
	o.releaseSeatCh <- seat{}
}

func (o *host) run(maxConcurrent int) {

	occupiedSeatCounter := 0
	philosQueue := make([]request, 0)

	for {
		if occupiedSeatCounter < maxConcurrent && len(philosQueue) > 0 {
			p := philosQueue[0]
			philosQueue = philosQueue[1:]

			occupiedSeatCounter++

			p.reply <- struct{}{}
			close(p.reply)
		}

		select {
		case r := <-o.requestSeatCh:
			if occupiedSeatCounter < maxConcurrent {
				occupiedSeatCounter++
				r.reply <- struct{}{}
				close(r.reply)
				continue
			}
			philosQueue = append(philosQueue, r)
		case <-o.releaseSeatCh:
			if occupiedSeatCounter > 0 {
				occupiedSeatCounter--
			}
		}
	}
}

type philosopher struct {
	n               int
	leftCS, rightCS *chopStick
}

func (o *philosopher) eat(h *host, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		h.acquire()
		o.leftCS.Lock()
		o.rightCS.Lock()
		fmt.Printf("start eating %d\n", o.n)
		o.rightCS.Unlock()
		o.leftCS.Unlock()
		fmt.Printf("end eating %d\n", o.n)
		h.release()
	}
}

func main() {

	chopSticks := make([]*chopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(chopStick)
	}

	philosophers := make([]philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = philosopher{
			n:       i,
			leftCS:  chopSticks[i],
			rightCS: chopSticks[(i+1)%5],
		}
	}

	h := newHost()
	wg := new(sync.WaitGroup)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(h, wg)
	}
	wg.Wait()

}
