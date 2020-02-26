package word

import (
	"testing"
	"math/rand"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// randomPalindrome 은 길이와 내용이
// 의사 난수 생성기 rng에서 만들어진 회문을 반환한다
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // 24까지의 임의 길이
	runes := make([]rune, n)
	for i := 0; i < (n+1) / 2 ; i++ {
		r := rune(rng.Intn(0x1000)) // '\u0999' 까지의 임의의룬
		runes[i] = r
		runes[n - 1 - i] = r
	}
	return string(runes)
}

func TestRndomPalindrome(t *testing.T) {
	// 의사 난수 생성기 초기화
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}