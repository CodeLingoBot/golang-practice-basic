package gotest

func IsPalindrome(text string) bool {
	for i := range text {
		if text[i] != text[len(text)-1-i] {
			return false
		}
	}
	return true
}
