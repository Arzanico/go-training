package main

import "testing"

func TestIsPalindrome_TableDriven(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   int
		want bool
	}{
		// --- invalid / quick rejects ---
		{"negative", -1, false},
		{"negative_large", -121, false},
		{"ends_with_zero_not_zero", 10, false},
		{"ends_with_zero_100", 100, false},
		{"ends_with_zero_1010", 1010, false},

		// --- zero & single digit ---
		{"zero", 0, true},
		{"single_digit_5", 5, true},
		{"single_digit_9", 9, true},

		// --- two digits ---
		{"two_digits_pal_11", 11, true},
		{"two_digits_not_pal_12", 12, false},
		{"two_digits_pal_22", 22, true},
		{"two_digits_not_pal_98", 98, false},

		// --- even length ---
		{"even_pal_1221", 1221, true},
		{"even_pal_4884", 4884, true},
		{"even_not_pal_1231", 1231, false},
		{"even_not_pal_1001_is_pal", 1001, true}, // important: has zeros inside

		// --- odd length (middle digit ignored) ---
		{"odd_pal_121", 121, true},
		{"odd_pal_12321", 12321, true},
		{"odd_pal_10501", 10501, true}, // zeros inside, still palindrome
		{"odd_not_pal_12341", 12341, false},
		{"odd_not_pal_12021_false", 12021, true}, // actually palindrome, good sanity

		// --- tricky / near palindromes ---
		{"near_pal_1002", 1002, false},
		{"near_pal_1223", 1223, false},
		{"pal_with_many_zeros_inside_1000001", 1000001, true},

		// --- larger numbers ---
		{"large_pal_123454321", 123454321, true},
		{"large_not_pal_123456789", 123456789, false},
		{"large_pal_2000000002", 2000000002, true},

		// --- int32 boundary-ish (still within int on 64-bit) ---
		{"max_int32_not_pal", 2147483647, false},
		{"max_int32_like_pal_2147447412", 2147447412, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := IsPalindrome(tt.in)
			if got != tt.want {
				t.Fatalf("IsPalindrome(%d) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome_DoesNotMutateInputVariable(t *testing.T) {
	t.Parallel()

	// Nota: En Go, los ints se pasan por valor, así que esto siempre “no muta” el caller.
	// Igual sirve para dejar explícito el contrato esperado.
	x := 12321
	_ = IsPalindrome(x)
	if x != 12321 {
		t.Fatalf("input was mutated: got %d, want %d", x, 12321)
	}
}

// Optional: property-style sanity check without external libs.
// For palindromes constructed by mirroring, IsPalindrome should return true.
func TestIsPalindrome_ConstructedPalindromes(t *testing.T) {
	t.Parallel()

	makeEvenPalindrome := func(prefix int) int {
		// Example: prefix=123 => 123321
		x := prefix
		y := prefix
		for y > 0 {
			x = x*10 + (y % 10)
			y /= 10
		}
		return x
	}

	makeOddPalindrome := func(prefix int) int {
		// Example: prefix=123 => 12321  (mirror excluding last digit)
		x := prefix
		y := prefix / 10
		for y > 0 {
			x = x*10 + (y % 10)
			y /= 10
		}
		return x
	}

	// Keep prefixes small to avoid overflow on 32-bit environments.
	for p := 1; p <= 5000; p += 137 {
		even := makeEvenPalindrome(p)
		if !IsPalindrome(even) {
			t.Fatalf("expected even constructed palindrome to be true, p=%d, n=%d", p, even)
		}

		odd := makeOddPalindrome(p)
		if !IsPalindrome(odd) {
			t.Fatalf("expected odd constructed palindrome to be true, p=%d, n=%d", p, odd)
		}
	}
}
