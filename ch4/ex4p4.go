// Write a version of `rotate` that operates in a single pass
// O(n)

// s := []int{0, 1, 2, 3, 4, 5} => [2 3 4 5 0 1]
package ch4

func rotate(s []int, k int) {
	// k >= len : k = k % len
	// if k == 0 return as is

	// k < len
	// copy left k elements to somewhere
	// shift [k+1, len(s)-1] to [0, len(s)-k-1]
	// copy k elements from right to left until k is 0

	n := len(s)

	if n == 0 {
		return
	}

	if k >= n {
		k = k % n
	}

	if k == 0 {
		return
	}

	//           n
	// 0 1 2 3 4 5
	// 1 2 3 4 5
	//       k

	tmp := make([]int, k)
	copy(tmp, s[:k])
	copy(s[:n-k], s[k:n])
	copy(s[n-k:], tmp)
}
