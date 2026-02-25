package main

// IsPalindrome returns true if x is a palindrome number in base 10.
//
// Algorithm: reverse only the second half of the number.
// - Reject negatives.
// - Reject numbers ending with 0 (except 0 itself).
// - Build reversedHalf until it reaches/exceeds the remaining x.
// - For even digits: x == reversedHalf
// - For odd digits:  x == reversedHalf/10 (drop the middle digit)
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	return x == reversedHalf || x == reversedHalf/10
}

func main() {
	IsPalindrome(121)
}
