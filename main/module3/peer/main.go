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

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

var (
	cow   = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird  = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake = Animal{food: "mice", locomotion: "slither", noise: "hsss"}
)

func requestAnimal() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")

	if scanner.Scan() {
		input := scanner.Text()

		words := strings.Fields(input)

		if len(words) < 2 {
			fmt.Println("Please enter both animal and info (e.g., 'cow eat')")
			return
		}

		animal := words[0]
		request := words[1]

		switch animal {
		case "cow":
			switch request {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			}

		case "bird":
			switch request {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			}

		case "snake":
			switch request {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			}

		default:
			fmt.Println("Unknown")
		}
	}
}

func main() {
	for {
		requestAnimal()
	}
}
