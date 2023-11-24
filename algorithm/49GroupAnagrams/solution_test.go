package groupanagrams

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {

	var tests = []struct {
		input []string
		want  [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
	}
	for _, test := range tests {
		// use without isEqual
		if got := groupAnagrams(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("groupAnagrams(%v) = %v", test.input, got)
		}
	}
}
