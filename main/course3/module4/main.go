package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
}

func (o Philo) eat() {
	for i := 0; i < 5; i++ {
		o.leftCS.Lock()
		o.rightCS.Lock()
		fmt.Println("eating")
		o.rightCS.Unlock()
		o.leftCS.Unlock()
	}
}

func main() {
	cSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(ChopS)
	}

	philo := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philo[i] = &Philo{
			leftCS:  cSticks[i],
			rightCS: cSticks[(i+1)%5],
		}
	}

	for i := 0; i < 5; i++ {
		go philo[i].eat()
	}

}
