package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Swap(nums []int, i int) {
	nums[i], nums[i+1] = nums[i+1], nums[i]
}

func BubbleSort(nums []int) {
	n := len(nums)
	for pass := 0; pass < n-1; pass++ {
		for i := 0; i < n-1-pass; i++ {
			if nums[i] > nums[i+1] {
				Swap(nums, i)
			}
		}
	}
}

func main() {
	fmt.Println("Pleas, enter a sequence of number, no more than 10.")

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
		if len(textUserInput) > 10 {
			fmt.Println("Wrong input - check length")
			continue
		}

		sanitizedUserInput := make([]int, 0)
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

			sanitizedUserInput = append(sanitizedUserInput, n)
		}

		BubbleSort(sanitizedUserInput)

		fmt.Println(sanitizedUserInput)
	}
}
