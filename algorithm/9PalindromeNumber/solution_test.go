package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{121, true},
		{1234321, true},
		{-121, false},
		{10, false},
		{0, true},
		{1000021, false},
	}
	for _, test := range tests {
		if got := isPalindrome(test.input); got != test.want {
			t.Errorf("isPalindrome(%v) = %v", test.input, got)
		}
	}
}
