package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(sli []int, idx int) {
	temp := sli[idx]
	sli[idx] = sli[idx+1]
	sli[idx+1] = temp
}

func BubbleSort(sli []int) {
	for j := 1; j < len(sli); j++ {
		for i := 0; i < len(sli)-1; i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
			}
		}
	}
}

func main() {
	fmt.Print("Enter multiple integers(seperated by space): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fields := strings.Fields(input)

	sli := []int{}
	for i := 0; i < len(fields); i++ {
		val, e := strconv.Atoi(fields[i])
		if e == nil {
			sli = append(sli, val)
		}
	}

	fmt.Printf("Input: %v ", sli)
	BubbleSort(sli)
	fmt.Printf("-> Sorted: %v\n", sli)

}
