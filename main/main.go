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

var actions = map[string]func(Animal) string{
	"eat": func(a Animal) string {
		return a.eat()
	},
	"move": func(a Animal) string {
		return a.move()
	},
	"speak": func(a Animal) string {
		return a.speak()
	},
}

type Animal interface {
	eat() string
	move() string
	speak() string
}

func newAnimal(animalType string, name string) (Animal, error) {
	a, ok := animals[strings.TrimSpace(strings.ToLower(animalType))]
	if !ok {
		return Animal(nil), fmt.Errorf("%s not recognized. Available options are: bird, snake and cow\n", animalType)
	}
	return a(name), nil
}

type bird struct {
	name string
}

func newBird(name string) Animal {
	return bird{
		name: name,
	}
}

func (o bird) eat() string {
	return fmt.Sprintf("worms")
}

func (o bird) move() string {
	return fmt.Sprintf("fly")
}

func (o bird) speak() string {
	return fmt.Sprintf("peep")
}

type snake struct {
	name string
}

func newSnake(name string) Animal {
	return snake{
		name: name,
	}
}

func (o snake) eat() string {
	return fmt.Sprintf("mice")
}

func (o snake) move() string {
	return fmt.Sprintf("slither")
}

func (o snake) speak() string {
	return fmt.Sprintf("hsss")
}

type cow struct {
	name string
}

func newCow(name string) Animal {
	return cow{
		name: name,
	}
}

func (o cow) eat() string {
	return fmt.Sprintf("grass")
}

func (o cow) move() string {
	return fmt.Sprintf("walk")
}

func (o cow) speak() string {
	return fmt.Sprintf("moo")
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
		"1: eat\n" +
		"2: move\n" +
		"3: speak\n")
	fmt.Println("the command <exit> will stop the program\n")

	userAnimals := make(map[string]Animal)
	for {
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
		if commands := strings.Split(input, " "); len(commands) < 3 {
			fmt.Println("command nos valid, check your spelling")
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
			fmt.Println(acc(a))
		default:
			continue

		}

	}

}
