package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var animals = map[string]func(string) Animal{
	"bird":  newBird,
	"snake": newSnake,
	"cow":   newCow,
}

var actions = map[string]func(Animal){
	"Eat": func(a Animal) {
		a.Eat()
	},
	"Move": func(a Animal) {
		a.Move()
	},
	"Speak": func(a Animal) {
		a.Speak()
	},
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func newAnimal(animalType string, name string) (Animal, error) {
	a, ok := animals[strings.TrimSpace(strings.ToLower(animalType))]
	if !ok {
		return Animal(nil), fmt.Errorf("%s not recognized. Available options are: bird, snake and cow\n", animalType)
	}
	return a(name), nil
}

type bird struct {
	name  string
	eat   string
	move  string
	speak string
}

func newBird(name string) Animal {
	return bird{
		name:  name,
		eat:   "worms",
		move:  "fly",
		speak: "peep",
	}
}

func (o bird) Eat() {
	fmt.Println(o.eat)
}

func (o bird) Move() {
	fmt.Println(o.move)
}

func (o bird) Speak() {
	fmt.Println(o.speak)
}

type snake struct {
	name  string
	eat   string
	move  string
	speak string
}

func newSnake(name string) Animal {
	return snake{
		name:  name,
		eat:   "mice",
		move:  "slither",
		speak: "hsss",
	}
}

func (o snake) Eat() {
	fmt.Println(o.eat)
}

func (o snake) Move() {
	fmt.Println(o.move)
}

func (o snake) Speak() {
	fmt.Println(o.speak)
}

type cow struct {
	name  string
	eat   string
	move  string
	speak string
}

func newCow(name string) Animal {
	return cow{
		name:  name,
		eat:   "grass",
		move:  "walk",
		speak: "moo",
	}
}

func (o cow) Eat() {
	fmt.Println(o.eat)
}

func (o cow) Move() {
	fmt.Println(o.move)
}

func (o cow) Speak() {
	fmt.Println(o.speak)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hi, Welcome to Friend from the Farm")
	fmt.Println("Here you can choose a set of animals from the list of available animal")
	fmt.Println("You will be able to provide a name for each animal")
	fmt.Println("After that you will be able to get some information from that animal")
	fmt.Println("Usage: You have 3 commands:\n" +
		"newAnimal" +
		"query\n" +
		"exit\n")
	fmt.Println("newAnimal allow you to create a new animal, for this comand you need to provide a name and a type for the animals as follows" +
		"newAnimal <name> <type>")
	fmt.Println("query allows you to get animal information, for this comand you need to provide the name of the animal you want the information from and an action as follows" +
		"query <name> action")
	fmt.Println("THere are 3 abailable action for each animal\n" +
		"1: Eat\n" +
		"2: Move\n" +
		"3: Speak\n")
	fmt.Println("the command <exit> will stop the program\n")

	userAnimals := make(map[string]Animal)
	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("unknown error", err)
			return
		}

		input = strings.ToLower(strings.TrimSpace(input))
		if input == "exit" {
			fmt.Println("shutting down....")
			return
		}

		var command, animalName, property string
		if commands := strings.Fields(input); len(commands) < 3 {
			fmt.Println("command not valid, check your spelling")
			continue
		} else {
			command = commands[0]
			animalName = commands[1]
			property = commands[2]
		}

		switch command {
		case "newanimal":
			a, newAnimalErr := newAnimal(property, animalName)
			if newAnimalErr != nil {
				fmt.Printf(" error %s", newAnimalErr.Error())
				continue
			}
			userAnimals[animalName] = a
			fmt.Println("Created it!")
		case "query":
			a, ok := userAnimals[animalName]
			if !ok {
				fmt.Printf("animal %s not found\n", animalName)
				continue
			}

			acc, ok := actions[property]
			if !ok {
				fmt.Printf("action %s not found\n", property)
				continue
			}
			acc(a)
		default:
			continue

		}

	}

}
