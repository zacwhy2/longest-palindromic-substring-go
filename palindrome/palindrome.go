package palindrome

// LongestPalindrome returns the longest palindromic substring in string s
func LongestPalindrome(s string) string {
	var isEven bool
	longestIndex := -1
	longestLength := 0
	for i := range s {
		isOddPalindrome := true
		isEvenPalindrome := true
		for j := 0; ; j++ {
			if i-j < 0 || i+j >= len(s) {
				// index out of range
				break
			}
			if isOddPalindrome {
				isOddPalindrome = isPalindrome(s, i, j, false)
			}
			if isEvenPalindrome {
				isEvenPalindrome = isPalindrome(s, i, j, true)
			}
			if !isOddPalindrome && !isEvenPalindrome {
				break
			}
			if j > longestLength {
				longestIndex = i
				longestLength = j
				isEven = isEvenPalindrome
			}
		}
	}
	if longestLength == 0 {
		return ""
	}
	start := longestIndex - longestLength
	end := longestIndex + longestLength + 1
	if isEven {
		end++
	}
	return s[start:end]
}

func isPalindrome(s string, i, j int, isEven bool) bool {
	offset := 0
	if isEven {
		offset++
	}
	if i-j < 0 || i+j+offset >= len(s) {
		// index out of range
		return false
	}
	if s[i-j] != s[i+j+offset] {
		// not a palindrome anymore
		return false
	}
	return true
}
