package palindrome

import (
	"bytes"
	"unicode"
)

func reverseString(s string) (reverse string) {
	for _, ch := range s {
		reverse = string(ch) + reverse
	}
	return
}

func formatString(s string) string {
	var b bytes.Buffer
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			b.WriteRune(unicode.ToLower(ch))
		}
	}
	return b.String()
}

func IsPalindromeTwoPointers(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		l := rune(s[left])
		r := rune(s[right])
		if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			left += 1
		} else if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			right -= 1
		} else if unicode.ToLower(l) == unicode.ToLower(r) {
			left += 1
			right -= 1
		} else {
			return false
		}
	}
	return true
}

func IsPalindromeReverse(s string) bool {
	formattedS := formatString(s)
	reversedS := reverseString(formattedS)
	return formattedS == reversedS
}

func AlmostPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
		left += 1
		right -= 1
	}
	return true
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}
