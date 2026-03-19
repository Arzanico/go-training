package main

import "fmt"

func main() {
	s := "bbbbbbbbbbbbbbbbbbbbbbbbd"

	var validWindow string
	var left int

	control := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		for {
			if _, exist := control[s[i]]; !exist {
				break
			}

			delete(control, s[left])
			left++
		}

		control[s[i]] = struct{}{}

		if len(s[left:i+1]) > len(validWindow) {
			validWindow = s[left : i+1]
		}
	}

	fmt.Printf("Max chain found is %v with a length of %d\n", validWindow, len(validWindow))
}
