package palindromenumber

// https://leetcode.com/problems/palindrome-number/description/
func isPalindrome(x int) bool {
	created := 0
	original := x
	ok := true
	if (x%10 == 0 && x != 0) || x < 0 {
		return false
	}
	for ok {
		remainder := x % 10
		created = (created * 10) + remainder
		x = (x - remainder) / 10
		if x == 0 {
			ok = false
		}
	}
	return original == created
}

// write unit test
