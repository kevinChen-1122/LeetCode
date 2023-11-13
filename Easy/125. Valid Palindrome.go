package Easy

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	clearedStr := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)
	strArr := []rune(clearedStr)

	var leftPtr, rightPtr rune
	for i := 0; i < len(strArr); i++ {
		leftPtr = strArr[i]
		rightPtr = strArr[len(strArr)-1-i]

		if i < len(strArr)-1-i {
			if leftPtr != rightPtr {
				return false
			}
		}
	}
	return true
}
