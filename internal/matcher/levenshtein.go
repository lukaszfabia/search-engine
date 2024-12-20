package matcher

// Calculates levenshtein distance between words
func Dist(lhs, rhs string) int {
	self := lhs
	other := rhs

	if self == other {
		return 0
	}

	m := len(self)
	n := len(other)

	if n == 0 {
		return m
	}

	if m == 0 {
		return n
	}

	m += 1
	n += 1

	distances := make([][]int, m)
	for i := range distances {
		distances[i] = make([]int, n)
	}

	// fill vertically
	for i := 0; i < m; i++ {
		distances[i][0] = i
	}

	// skip 0, 0 position cuz its reserved
	for j := 1; j < n; j++ {
		distances[0][j] = j
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			var cost = 0

			if self[i-1] != other[j-1] {
				cost = 1
			}

			distances[i][j] = min(
				distances[i-1][j]+1,
				distances[i][j-1]+1,
				distances[i-1][j-1]+cost,
			)
		}
	}

	return distances[m-1][n-1]
}
