package binarysearch

import "testing"

func TestSearch(t *testing.T) {

	var tests = []struct {
		input  []int
		target int
		want   int
	}{
		{input: []int{-1, 0, 3, 5, 9, 12}, target: 9, want: 4},
		{input: []int{-1, 0, 3, 5, 9, 12}, target: 2, want: -1},
		{input: []int{5}, target: 5, want: 0},
		{input: []int{5}, target: 6, want: -1},
		{input: []int{5, 6}, target: 5, want: 0},
		{input: []int{5, 6}, target: 6, want: 1},
		{input: []int{5, 6}, target: 1, want: -1},
	}

	for _, test := range tests {
		if got := Search(test.input, test.target); got != test.want {
			t.Errorf("search(%v, %v) = %v", test.input, test.target, got)
		}
	}
}
