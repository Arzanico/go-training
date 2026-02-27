package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type seat struct{}

type chopStick struct {
	sync.Mutex
}

type request struct {
	uuid  string
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
		uuid:  uuid.NewString(),
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
	requestQueue := make([]request, 0)

	for {
		if occupiedSeatCounter < maxConcurrent && len(requestQueue) > 0 {
			r := requestQueue[0]
			requestQueue = requestQueue[1:]

			occupiedSeatCounter++

			r.reply <- struct{}{}
			close(r.reply)
			fmt.Printf("[HOST] DEQ   req=%s  action=GRANT_FROM_QUEUE  occupied=%d->%d  qlen=%d\n",
				r.uuid, occupiedSeatCounter, occupiedSeatCounter+1, len(requestQueue))
		}

		select {
		case r := <-o.requestSeatCh:
			fmt.Printf("[HOST] RECV  req=%s  occupied=%d  qlen=%d\n",
				r.uuid, occupiedSeatCounter, len(requestQueue))
			if occupiedSeatCounter < maxConcurrent {
				occupiedSeatCounter++

				r.reply <- struct{}{}
				close(r.reply)
				fmt.Printf("[HOST] ACT   req=%s  action=%s  occupied=%d  qlen=%d\n",
					r.uuid, "GRANT", occupiedSeatCounter, len(requestQueue))
				continue
			}
			requestQueue = append(requestQueue, r)
			fmt.Printf("[HOST] ACT   req=%s  action=%s  occupied=%d  qlen=%d\n",
				r.uuid, "QUEUED", occupiedSeatCounter, len(requestQueue))
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
