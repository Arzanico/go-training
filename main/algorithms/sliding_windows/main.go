package main

import "fmt"

func main() {
	s := "bbbbbbbbbbbbbbbbbbbbbbbbd"
	v := findLongestSubStringWithNoDuplicated(s)
	fmt.Printf("Max unique chain found %s length %d\n", v, len(v))

	i := []int{1, 2, 6, 4, 3, 5, 4}
	target := 11
	vI := findFirstSubGroupGraterThanTarget(i, target)
	fmt.Printf("found %v for which sum is equal or grater than %d\n", vI, len(v))

}

func findLongestSubStringWithNoDuplicated(s string) string {
	runes := []rune(s)

	var validWindow []rune
	var left int

	control := make(map[rune]struct{})
	for i := 0; i < len(runes); i++ {
		for {
			if _, exist := control[runes[i]]; !exist {
				break
			}

			delete(control, runes[left])
			left++
		}

		control[runes[i]] = struct{}{}
		if len(runes[left:i+1]) > len(validWindow) {
			validWindow = runes[left : i+1]
		}
	}
	return string(validWindow)
}

func findFirstSubGroupGraterThanTarget(g []int, t int) []int {
	var sum int
	for i := 0; i < len(g); i++ {
		sum += g[i]
		if sum >= t {
			return g[0 : i+1]
		}

	}

	return nil
}
