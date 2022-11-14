package longestsubstring

func LengthOfLongestSubstring(s string) int {
	longest := 0
	left := 0
	seen := make(map[rune]int)
	for right, ch := range s {
		if val, ok := seen[ch]; ok && val >= left {
			left = val + 1
		} else {
			windowSize := right - left + 1
			if windowSize > longest {
				longest = windowSize
			}
		}
		seen[ch] = right
	}
	return longest
}
