// Package word 에는 낱말 게임을 위한 유틸리티가 있다
package word

import "unicode"

// IsPalindrome 은 s가 회문인지 여부를 보고한다
// 대소문자와 비문자는 무시한다
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters) - 1 - i] {
			return false
		}
	}
	return true
}