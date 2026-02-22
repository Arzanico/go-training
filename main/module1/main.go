package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func BubbleSort(userInput []int) {

	for i := 0; i < len(userInput)-1; i++ {
		for j, v := range userInput {
			if j == i {
				continue
			}
			if userInput[i] > v {
				Swap(userInput, i)
			}
		}
	}

}

func Swap(l []int, index int) {
	l[index], l[index+1] = l[index+1], l[index]
}

func main() {
	fmt.Println("Pleas, enter a sequence of 10 number, no less no more than that.")

	userInputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n> ")
		textUserInput, err := userInputReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		if textUserInput == "exit" {
			break
		}

		textUserInput = strings.TrimSuffix(textUserInput, "\n")
		if len(textUserInput) != 10 {
			fmt.Println("Wrong input - check length")
			continue
		}

		bubleSortedUserInput := make([]int, 0)
		sanitizedTextUserInput := strings.Join(strings.Fields(textUserInput), "")
		for _, v := range sanitizedTextUserInput {
			if !unicode.IsDigit(v) {
				fmt.Println("Wrong input - check non numeric characters")
				continue
			}

			n, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Println(err)
				continue
			}

			bubleSortedUserInput = append(bubleSortedUserInput, n)
		}

		BubbleSort(bubleSortedUserInput)

		fmt.Println(bubleSortedUserInput)
	}
}
