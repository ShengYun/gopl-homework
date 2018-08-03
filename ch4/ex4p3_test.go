package ch4

import (
	"testing"
)

var tableTest = []struct {
	input  []int
	expect []int
}{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{1, 2}, []int{2, 1}},
	{[]int{1, 2, 3}, []int{3, 2, 1}},
}

var tableTest2 = []struct {
	input  [3]int
	expect [3]int
}{
	{[...]int{1, 2, 3}, [...]int{3, 2, 1}},
}

func sliceValueEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestReverseOrig(t *testing.T) {

	for _, v := range tableTest {
		actual := v.input
		reverseOrig(actual)
		t.Logf("%v %v", actual, v.input)
		if !sliceValueEqual(actual, v.expect) {
			t.Errorf("reverseOrig(%v), expected %v, actual %v", v.input, v.expect, actual)
		}
	}
}

func TestReverse(t *testing.T) {
	for _, v := range tableTest2 {
		actual := v.input
		reverse(&actual)
		if !sliceValueEqual(actual[:], v.expect[:]) {
			t.Errorf("reverseOrig(%v), expected %v, actual %v", v.input, v.expect, actual)
		}
	}

}
