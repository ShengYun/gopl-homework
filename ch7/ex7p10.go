// The sort.Interface type can be adapted to other uses. Write a function
// IsPalindrome(s sort.Interface) bool that reports whether the sequence s is a
// palindrome, in other words, reversing the sequence would not change it.
// Assume that the elements at indices i and j are equal if !s.Less(i, j) &&
// !s.Less(j, i)

package ch7

import (
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	// Finally, Go has no comma operator and ++ and -- are statements not
	// expressions. Thus if you want to run multiple variables in a for you
	// should use parallel assignment (although that precludes ++ and --).
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		}
		return false
	}
	return true
}
