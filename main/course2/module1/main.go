package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(nums []int, i int) {
	nums[i], nums[i+1] = nums[i+1], nums[i]
}

func BubbleSort(nums []int) {
	n := len(nums)
	for loop := 0; loop < n-1; loop++ {
		for i := 0; i < n-1-loop; i++ {
			if nums[i] > nums[i+1] {
				Swap(nums, i)
			}
		}
	}
}

func main() {
	fmt.Println("Please, enter a sequence of number, no more than 10.")

	userInputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n> ")
		textUserInput, err := userInputReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		sanitizedTextUserInput := strings.TrimSuffix(textUserInput, "\n")
		if strings.EqualFold("exit", sanitizedTextUserInput) {
			break
		}

		if len(strings.Fields(sanitizedTextUserInput)) > 10 {
			fmt.Println("Wrong input - check length")
			continue
		}

		sanitizedUserIntList := make([]int, 0)
		for _, v := range strings.Fields(sanitizedTextUserInput) {

			n, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				continue
			}

			sanitizedUserIntList = append(sanitizedUserIntList, n)
		}

		BubbleSort(sanitizedUserIntList)

		fmt.Println(sanitizedUserIntList)
	}
}
