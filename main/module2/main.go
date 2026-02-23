package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(a, v0, a0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + a0
	}
}

func main() {

	fmt.Println("Please introduce 3 float separated by spaces. Your integer will represents , acceleration, initial velocity and initial displacement")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("> ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	if strings.TrimSpace(strings.ToLower(userInput)) == "q" {
		fmt.Println("Bye!")
		return
	}

	userInputParts := strings.Fields(strings.TrimSuffix(userInput, "\n"))
	if len(userInputParts) != 3 {
		fmt.Println("Invalid prompt")
		return
	}

	aIvIa := make([]float64, 3, 3)
	for i, p := range userInputParts {
		v, err := strconv.ParseFloat(p, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		aIvIa[i] = v
	}

	a, iV, iA := aIvIa[0], aIvIa[1], aIvIa[2]
	fn := GenDisplaceFn(a, iV, iA)

	fmt.Println("Please enter a number for time")
	userTimeInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(userTimeInput, "\n")), 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Displacement is %g", fn(t))
	return

}
