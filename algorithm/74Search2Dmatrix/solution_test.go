package search2dmatrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	testCases := []struct {
		input  [][]int
		target int
		want   bool
	}{
		{
			input: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			target: 3,
			want:   true,
		},
		{
			input: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			target: 13,
			want:   false,
		},
	}
	for _, tt := range testCases {
		if got := searchMatrix(tt.input, tt.target); got != tt.want {
			t.Errorf("searchMatrix(%v, %v) = %v", tt.input, tt.target, got)
		}
	}
}
