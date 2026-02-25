package main

import (
	"fmt"
	"sync"
)

var askForEatChannel = make(chan struct{})

var permissionChannel = make(chan bool)

var philoChannel = make(chan struct{})

type chopStick struct {
	sync.Mutex
}

type Philo struct {
	n               int
	leftCs, rightCs *chopStick
}

func (o Philo) eat(wg *sync.WaitGroup) {
	counter := 0
	for {
		if counter == 3 {
			break
		}
		askForEatChannel <- struct{}{}
		select {
		case p := <-permissionChannel:
			if p {
				o.leftCs.Lock()
				o.rightCs.Lock()
				fmt.Printf("starting to eat %d\n", o.n)
				fmt.Printf("finishing to eat %d\n", o.n)
				o.rightCs.Unlock()
				o.leftCs.Unlock()
				philoChannel <- struct{}{}
				counter++
			}
		}
	}
	wg.Done()
}

func main() {

	cSticks := make([]*chopStick, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(chopStick)
	}

	go host()

	philo := make([]Philo, 5)
	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		philo[i] = Philo{
			n:       i + 1,
			leftCs:  cSticks[i],
			rightCs: cSticks[(i+1)%5],
		}

		go philo[i].eat(&wg)
	}
	wg.Wait()

}

func host() {
	counter := 0
	for {

		select {
		case <-askForEatChannel:
			if counter < 2 {
				permissionChannel <- true
				counter++
			} else {
				permissionChannel <- false
			}
		case <-philoChannel:
			counter--
		}
	}
}
