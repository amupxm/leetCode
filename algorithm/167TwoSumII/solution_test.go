package twosumii

import (
	"reflect"
	"testing"
)

func Test_TwoSum(t *testing.T) {

	var tests = []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}
	for _, test := range tests {
		if got := twoSum(test.nums, test.target); !reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum(%v, %v) = %v", test.nums, test.target, got)
		}
	}
}
