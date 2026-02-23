package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (cow Cow) Speak() string {
	return cow.noise
}

func (cow Cow) Eat() string {
	return cow.food
}

func (cow Cow) Move() string {
	return cow.locomotion
}

func (bird Bird) Speak() string {
	return bird.noise
}

func (bird Bird) Eat() string {
	return bird.food
}

func (bird Bird) Move() string {
	return bird.locomotion
}

func (snake Snake) Speak() string {
	return snake.noise
}

func (snake Snake) Eat() string {
	return snake.food
}

func (snake Snake) Move() string {
	return snake.locomotion
}

var Animals map[string]Animal = map[string]Animal{}

func newCow() Cow {
	return Cow{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
}

func newBird() Bird {
	return Bird{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
}

func newSnake() Snake {
	return Snake{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
}

func log_created() {
	log.Print("Created it!")

}
func main() {
	input_scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n> ")
		input_scanner.Scan()
		raw_input := input_scanner.Text()
		inp := strings.Split(raw_input, " ")
		if inp[0] == "newanimal" || inp[0] == "n" {
			switch inp[2] {
			case "cow":
				Animals[inp[1]] = newCow()
				log_created()
			case "bird":
				Animals[inp[1]] = newBird()
				log_created()
			case "snake":
				Animals[inp[1]] = newSnake()
				log_created()
			}
		} else {
			if inp[0] == "query" || inp[0] == "q" {
				req_animal := Animals[inp[1]]
				if req_animal != nil {
					switch inp[2] {
					case "speak":
						log.Println(req_animal.Speak())
					case "move":
						log.Println(req_animal.Move())
					case "eat":
						log.Println(req_animal.Eat())
					}
				} else {
					log.Println("- BAD Query - Try again")
				}
			} else {

				log.Println("- BAD Query - Try again")
			}
		}
	}
}
