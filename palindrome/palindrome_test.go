package palindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct{ s, want string }{
		{"zzzcbabcyyy", "cbabc"},
		{"zzzcbaabcyyy", "cbaabc"},
		{"zcbacy", ""},
		{"a", ""},
		{"aaaaaa", "aaaaaa"},
	}
	for _, tt := range tests {
		if got := LongestPalindrome(tt.s); got != tt.want {
			t.Errorf("LongestPalindrome(%s) = %s, want %s", tt.s, got, tt.want)
		}
	}
}
