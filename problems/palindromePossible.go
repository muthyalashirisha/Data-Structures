package problems

import "fmt"

// Function to check if a string is a palindrome
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// Function to check if concatenation of two strings is a palindrome
func isConcatPalindrome(s1, s2 string) bool {
	fmt.Println(s1, s2)
	return isPalindrome(s1 + s2)
}

// Function to check if it's possible to cut both strings at a common point to form a palindrome
func IsPalindromePossible(A, B string) bool {
	for i := 0; i < len(A); i++ {
		if isConcatPalindrome(A[:i], B[i:]) {
			return true
		}
	}
	return false
}

// Function to find the length of the longest palindrome that can be formed
func LongestPalindromeLength(A, B string) int {
	maxLen := 0
	for i := 0; i <= len(A); i++ {
		for j := 0; j <= len(B); j++ {
			// Concatenate substrings
			concat := A[:i] + B[j:]
			// Check if the concatenation forms a palindrome and update maxLen if it does
			if isPalindrome(concat) && len(concat) > maxLen {
				fmt.Println(concat)
				maxLen = len(concat)
			}
		}
	}
	return maxLen
}

// Function to compute the failure function for KMP algorithm
func computePrefix(s string) []int {
	n := len(s)
	prefix := make([]int, n)
	prefix[0] = 0
	k := 0

	for i := 1; i < n; i++ {
		for k > 0 && s[k] != s[i] {
			k = prefix[k-1]
		}
		if s[k] == s[i] {
			k++
		}
		prefix[i] = k
	}
	return prefix
}

// Helper function to reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Function to find the length of the longest palindrome that can be formed
func LongestPalindromeLengthKMP(A, B string) int {
	rev := reverseString(B)
	combined := A + "#" + rev
	n := len(combined)

	// Compute the failure function for the combined string
	prefix := computePrefix(combined)

	maxLen1 := 0
	for i := 0; i < n; i++ {
		// The length of the longest palindrome centered at position i
		palindromeLen := prefix[i]
		// If the palindrome spans both strings, update maxLen
		if palindromeLen > 0 && i-palindromeLen == len(A) {
			maxLen1 = palindromeLen * 2
		}
	}

	reva := reverseString(A)
	combined1 := reva + "#" + B
	n1 := len(combined1)

	// Compute the failure function for the combined string
	prefix1 := computePrefix(combined1)

	maxLen2 := 0
	for i := 0; i < n1; i++ {
		// The length of the longest palindrome centered at position i
		palindromeLen := prefix1[i]
		// If the palindrome spans both strings, update maxLen
		if palindromeLen > 0 && i-palindromeLen == len(A) {
			maxLen2 = palindromeLen * 2
		}
	}

	return max(maxLen1, maxLen2)
}
