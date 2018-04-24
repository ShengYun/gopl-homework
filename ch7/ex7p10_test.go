package ch7

import (
	"sort"
	"testing"
	// "string"
)

var tableTest = []struct {
	input  []string
	expect bool
}{
	{[]string{""}, true},
	{[]string{"a"}, true},
	{[]string{"a", "b", "a"}, true},
	{[]string{"a", "b", "b", "a"}, true},
	{[]string{"a", "a", "a", "a", "a"}, true},
	{[]string{"a", "a", "a", "a", "a", "a"}, true},

	{[]string{"a", "b"}, false},
	{[]string{"a", "b", "c"}, false},
	{[]string{"a", "b", "a", "a"}, false},
	{[]string{"a", "b", "a", "a", "a"}, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, v := range tableTest {
		actual := IsPalindrome(sort.StringSlice(v.input))
		if actual != v.expect {
			t.Errorf("IsPalindrome(%s), expected %v, actual %v", v.input, v.expect, actual)
		}
	}
}
