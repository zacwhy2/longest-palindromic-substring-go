package palindrome

// LongestPalindrome returns the longest palindromic substring in string s
func LongestPalindrome(s string) string {
	var start, end int
	for i := range s {
		var isOddPalindrome, isEvenPalindrome = true, true
		for j := 0; i-j >= 0 && i+j < len(s); j++ {
			if isOddPalindrome {
				isOddPalindrome = isPalindromic(s, i, j, false)
			}
			if isEvenPalindrome {
				isEvenPalindrome = isPalindromic(s, i, j, true)
			}
			if !isOddPalindrome && !isEvenPalindrome {
				break
			}
			if isOddPalindrome && (j*2+1) > (end-start+1) {
				start = i - j
				end = i + j
			}
			if isEvenPalindrome && (j*2+2) > (end-start+1) {
				start = i - j
				end = i + 1 + j
			}
		}
	}
	if start == end {
		return ""
	}
	return s[start : end+1]
}

func isPalindromic(s string, i, j int, isEven bool) bool {
	char1Index := i - j
	char2Index := i + j
	if isEven {
		char2Index++
	}
	return char1Index >= 0 &&
		char2Index < len(s) &&
		s[char1Index] == s[char2Index]
}
