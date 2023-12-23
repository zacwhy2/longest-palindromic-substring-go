package palindrome

func LongestPalindrome(s string) string {
	longestIndex := -1
	longestLength := 0
	for i := range s {
		for j := 0; ; j++ {
			if i-j < 0 || i+j >= len(s) {
				// index out of range
				break
			}
			if s[i-j] != s[i+j] {
				// not a palindrome anymore
				break
			}
			if j > longestLength {
				longestLength = j
				longestIndex = i
			}
		}
	}
	start := longestIndex - longestLength
	end := longestIndex + longestLength + 1
	return s[start:end]
}
