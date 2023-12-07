package threesum

import "testing"

func TestThreeSum(t *testing.T) {
	type test struct {
		input []int
		want  [][]int
	}

	tests := []test{
		{input: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{input: []int{}, want: [][]int{}},
		{input: []int{0}, want: [][]int{}},
		{input: []int{0, 0, 0}, want: [][]int{{0, 0, 0}}},
	}

	for _, tc := range tests {
		got := threeSum(tc.input)
		if len(got) != len(tc.want) {
			t.Errorf("input: %v, want: %v, got: %v", tc.input, tc.want, got)
		}
	}
}
