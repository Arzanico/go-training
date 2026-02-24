package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please provide a list of integers separated by a space")
		fmt.Print(">")

		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(strings.TrimSpace(userInput)) == 0 {
			fmt.Println("invalid input")
			continue
		}

		userList := make([]int, 0)
		for _, v := range strings.Fields(strings.TrimSuffix(userInput, "\n")) {
			n, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				continue
			}
			userList = append(userList, n)
		}

		sortedLists := make([][]int, 0)
		if len(userList) < 4 {
			fmt.Println("invalid input")
			continue
		} else {
			splitList := splitIn4(userList)
			wg := &sync.WaitGroup{}

			wg.Add(4)
			go sortRoutine(splitList[0], wg)
			go sortRoutine(splitList[1], wg)
			go sortRoutine(splitList[2], wg)
			go sortRoutine(splitList[3], wg)
			wg.Wait()

			for _, v := range splitList {
				sortedLists = append(sortedLists, v)
			}
		}

		newSortedList := make([]int, 0)
		for _, list := range sortedLists {
			newSortedList = append(newSortedList, list...)
		}

		sortList(newSortedList)
		fmt.Println(newSortedList)
	}

}

func sortRoutine(l []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sorting:", l)
	sortList(l)
}

func splitIn4(l []int) [][]int {
	n := len(l)
	base := n / 4
	rem := n % 4

	result := make([][]int, 4)

	start := 0
	for i := 0; i < 4; i++ {
		size := base
		if i < rem {
			size++
		}
		end := start + size
		result[i] = l[start:end]
		start = end
	}

	return result
}

func sortList(l []int) {
	for pass := 0; pass < len(l)-1; pass++ {
		for i := 0; i < len(l)-1-pass; i++ {
			if l[i] > l[i+1] {
				swap(l, i)
			}
		}
	}
}

func swap(l []int, index int) {
	l[index], l[index+1] = l[index+1], l[index]
}
