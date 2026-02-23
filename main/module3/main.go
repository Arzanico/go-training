package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (o Animal) Eat() {
	fmt.Println(o.food)
}

func (o Animal) Move() {
	fmt.Println(o.locomotion)
}

func (o Animal) Speak() {
	fmt.Println(o.noise)
}

func newAnimal(name string) (Animal, error) {
	switch name {
	case "cow":
		return Animal{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		}, nil
	case "bird":
		return Animal{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		}, nil
	case "snake":
		return Animal{
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		}, nil
	default:
		return Animal{}, fmt.Errorf("animal not found")
	}
}

func main() {
	userInput := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		userInput.Scan()
		parts := strings.Fields(userInput.Text())

		animal, err := newAnimal(parts[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch strings.ToLower(parts[1]) {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
		default:
			fmt.Println("Unknown action")
			continue

		}
	}
}
