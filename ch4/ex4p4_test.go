package ch4

import (
	"testing"
)

var rotateTableTest = []struct {
	s      []int
	k      int
	expect []int
}{
	{[]int{}, 0, []int{}},
	{[]int{}, 2, []int{}},
	{[]int{1}, 0, []int{1}},
	{[]int{1}, 1, []int{1}},
	{[]int{1}, 2, []int{1}},
	{[]int{1, 2}, 1, []int{2, 1}},
	{[]int{1, 2}, 2, []int{1, 2}},
	{[]int{1, 2}, 3, []int{2, 1}},
	{[]int{1, 2, 3, 4, 5}, 3, []int{4, 5, 1, 2, 3}},
}

func TestRotate(t *testing.T) {
	for _, v := range rotateTableTest {
		orig := make([]int, len(v.s))
		copy(orig, v.s)

		actual := v.s
		rotate(actual, v.k)
		if !sliceValueEqual(actual, v.expect) {
			t.Errorf("rotate(%v, %v) expect %v actual %v", orig, v.k, v.expect, actual)
		}
	}

}
