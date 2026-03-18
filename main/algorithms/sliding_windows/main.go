package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abbcccdbdrfgtrhiplfhjgrialcovnrrnlp"
	l := make([]string, 0)
	for i := range s {
		l = append(l, string(s[i]))
	}

	chains := make(map[string]int)
	window := make([]string, 0)
	var left int

	for i := 0; i < len(l); i++ {
		window = l[left:i]

		control := make(map[string]struct{})
		for j := range window {
			if _, exists := control[window[j]]; !exists {
				control[window[j]] = struct{}{}
			} else {
				chain := strings.Join(window, "")
				chains[chain] = len(chain)
				//fmt.Println("New chain created " + chain)
				left++
				continue
			}
		}

	}
	var maxChain string
	var maxLenChain int
	for k, v := range chains {
		if v > maxLenChain {
			maxLenChain = v
			maxChain = k
		}
	}

	fmt.Printf("Max chain found is %s with a length of %d\n", maxChain, maxLenChain)
}
