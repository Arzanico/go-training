package main

import (
	"bufio"
	"fmt"
	"os"
)

func BubbleSort(a []int) {

}

func Swap(a []int, i int) {
	a[i], a[i+1] = a[i+1], a[i]
}

func main() {
	fmt.Print("\n> ")
	rawUserInput := bufio.NewScanner(os.Stdin)
	rawUserInput.Scan()
	textUserInput := rawUserInput.Text()
	fmt.Println(textUserInput)
}
