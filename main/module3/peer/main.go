package main

import (
	b "bufio"
	f "fmt"
	"os"
	s "strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	f.Print(a.food, "\n")
}

func (a Animal) Move() {
	f.Print(a.locomotion, "\n")
}

func (a Animal) Speak() {
	f.Print(a.noise, "\n")
}

func main() {
	var cow, bird, snake, temp Animal
	var command []string
	cow.food, cow.locomotion, cow.noise = "grass", "walk", "moo"
	bird.food, bird.locomotion, bird.noise = "worms", "fly", "peep"
	snake.food, snake.locomotion, snake.noise = "mice", "slither", "hsss"

	scanner := b.NewScanner(os.Stdin)

	for true {
		f.Printf(">")
		if scanner.Scan() {
			command = s.Split(s.ToLower(scanner.Text()), " ")
			if len(command) == 2 && (command[0] == "cow" || command[0] == "bird" || command[0] == "snake") && (command[1] == "eat" || command[1] == "move" || command[1] == "speak") {
				switch command[0] {
				case "cow":
					temp = cow
				case "bird":
					temp = bird
				case "snake":
					temp = snake
				}
				switch command[1] {
				case "eat":
					temp.Eat()
				case "move":
					temp.Move()
				case "speak":
					temp.Speak()
				}
			} else {
				f.Print("Invalid Syntax: >[animal] [action]\n")
			}
		} else {
			f.Print("Error Scanning User Input\n")
		}
	}
}
