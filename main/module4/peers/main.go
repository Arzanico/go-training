package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	animals = make(map[string]Animal, 0)
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

func NewAnimal(kind string) Animal {
	switch kind {
	case "cow":
		return new(Cow)

	case "bird":
		return new(Bird)

	case "snake":
		return new(Snake)
	}

	return nil
}

type Cow struct{}

func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (a Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (a Bird) Eat() {
	fmt.Println("worms")
}

func (a Bird) Move() {
	fmt.Println("fly")
}

func (a Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (a Snake) Eat() {
	fmt.Println("mice")
}

func (a Snake) Move() {
	fmt.Println("slither")
}

func (a Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		s.Scan()

		t := strings.TrimSpace(s.Text())
		if t != "" {
			ss := strings.Split(t, " ")
			if len(ss) < 3 {
				continue
			}

			// comm, info := ss[0], ss[1], ss[2]
			comm := ss[0]
			switch comm {
			case "newanimal":
				name, kind := ss[1], ss[2]

				animals[name] = NewAnimal(kind)

			case "query":
				name, req := ss[1], ss[2]

				v, ok := animals[name]
				if !ok {
					fmt.Printf("animal '%s' not found\n", name)
					continue
				}

				switch req {
				case "eat":
					v.Eat()

				case "move":
					v.Move()

				case "speak":
					v.Speak()

				default:
					fmt.Printf("request '%s' not found\n", req)
					continue
				}
			}
		} else {
			// exit if user entered an empty string
			break
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
