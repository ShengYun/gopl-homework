// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.

package ch4

import "testing"

func compareStringSlice(lhs, rhs []string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i := 0; i < len(lhs); i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}

func Test_inplaceEliminateAdjacentDup(t *testing.T) {
	tests := []struct {
		name   string
		xs     []string
		expect []string
	}{
		{"test1", []string{}, []string{}},
		{"test1", []string{"1"}, []string{"1"}},
		{"test2", []string{"1", "2", "2", "2"}, []string{"1", "2"}},
		{"test3", []string{"2", "2", "2", "2"}, []string{"2"}},
		{"test4", []string{"2", "2", "4", "4"}, []string{"2", "4"}},
		{"test5", []string{"2", "2", "4"}, []string{"2", "4"}},
		{"test6", []string{"2", "2", "2", "2", "3", "3", "4", "4"}, []string{"2", "3", "4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inplaceEliminateAdjacentDup(tt.xs)
			if !compareStringSlice(got, tt.expect) {
				t.Errorf("inplaceEliminateAdjacentDup got = %v expect = %v", tt.xs, tt.expect)
			}
		})
	}
}
