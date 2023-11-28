package containerwiththemostwater

import "testing"

func TestMaxArea(t *testing.T) {
	type test struct {
		input []int
		want  int
	}

	tests := []test{
		{input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{input: []int{1, 1}, want: 1},
	}

	for _, tc := range tests {
		got := maxArea(tc.input)
		if got != tc.want {
			t.Errorf("input: %v, want: %d, got: %d", tc.input, tc.want, got)
		}
	}
}
