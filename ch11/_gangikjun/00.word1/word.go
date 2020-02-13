// Package word 는 낱말 게임을 위한 유틸리티가 있다
package word

// IsPalindrome 은 s가 회문인지 여부를 보고한다
// (첫 번째 시도)
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s) - 1 -i] {
			return false
		}
	}
	return true
}