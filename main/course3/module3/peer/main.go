package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil {
		return
	}

	line = strings.TrimSpace(line)
	fields := strings.Fields(line)

	if len(fields) < 4 {
		fmt.Println("The number should at least 4!")
		return
	}

	numbers := make([]int, len(fields))
	for i, field := range fields {
		num, err := strconv.Atoi(field)

		if err != nil {
			return
		}

		numbers[i] = num
	}

	snumbers := partsSlice(numbers)
	c := make(chan []int)

	for id, val := range snumbers {
		go func(id int, val []int) {
			fmt.Printf("Sorting Goroutine %d with value: %v\n", id, val)
			sort.Ints(val)
			c <- val
		}(id, val)
	}

	var storage [][]int
	for i := 0; i < 4; i++ {
		storage = append(storage, <-c)
	}

	mergeOne := merge(storage[0], storage[1])
	mergeTwo := merge(storage[2], storage[3])
	finalMerge := merge(mergeOne, mergeTwo)

	fmt.Print(finalMerge)
}

func partsSlice(nums []int) [][]int {
	n := len(nums)
	base := n / 4
	remainder := n % 4

	var nestedSlice [][]int
	start := 0

	for i := 0; i < 4; i++ {
		size := base

		if remainder > 0 {
			size++
			remainder--
		}

		end := start + size

		if end > n {
			end = n
		}

		currentSlice := nums[start:end]
		nestedSlice = append(nestedSlice, currentSlice)

		start = end
	}

	return nestedSlice
}

func merge(left, right []int) []int {
	size := len(left) + len(right)
	result := make([]int, 0, size)

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	if i < len(left) {
		result = append(result, left[i:]...)
	}

	if j < len(right) {
		result = append(result, right[j:]...)
	}

	return result
}
