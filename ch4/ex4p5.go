// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.

package ch4

func inplaceEliminateAdjacentDup(xs []string) []string {
	if len(xs) < 2 {
		return xs
	}

	// i stands for index to be compared and insert place - 1

	n := len(xs)

	i, j := 0, 0

	for j < n {

		if xs[i] != xs[j] {
			i++
			j++
			continue
		}

		// xs[i] == xs[j]
		for j < n && xs[i] == xs[j] {
			j++
		}

		if j != n {
			xs[i+1] = xs[j]
			i++
		}

		j++
	}

	return xs[:i+1]
}
