package dynamicProgramming

import (
	"fmt"
)

const (
	replace = 0
	remove  = 1
	insert  = 2
)

// MinEditDistance get min edit distance of two slice
func MinEditDistance(s1, s2 string) (int, []string) {
	len1 := len(s1) + 1
	len2 := len(s2) + 1

	// Create two-dimensional structure such that m[i][j] = 0 for i in 0 .. len1 and for j in 0 .. len2
	// dynamic programming memotizer
	m := make([][]int, len1)
	// index
	op := make([][]int, len1)

	for i := range m {
		// init with 0
		m[i] = make([]int, len1)
	}

	for i := range op {
		op[i] = make([]int, len1)
		for j := range op[i] {
			op[i][j] = -1
		}
	}

	//set up initial costs on horizontal and vertical
	for i := 1; i < len2; i++ {
		m[0][i] = i
	}

	for j := 1; j < len1; j++ {
		m[j][0] = j
	}

	// compute best
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			cost := 1
			if s1[i-1] == s2[j-1] {
				cost = 0
			}

			replaceCost := m[i-1][j-1] + cost
			removeCost := m[i-1][j] + 1
			insertCost := m[i][j-1] + 1

			costs := [...]int{replaceCost, removeCost, insertCost}
			m[i][j] = replaceCost
			for index, e := range costs {
				if e < m[i][j] {
					m[i][j] = e
					op[i][j] = index
				}
			}
		}
	}

	ops := make([]string, 0)
	x := len(s1)
	y := len(s2)

	for x > 0 || y > 0 {
		if op[x][y] == remove || y == 0 {
			ops = append(ops, fmt.Sprintf("remove %d-th char %q of %s", x, s1[x-1], s1))
			x = x - 1
		} else if op[x][y] == insert || x == 0 {
			ops = append(ops, fmt.Sprintf("insert %d-th char %q of %s", y, s2[y-1], s2))
			y = y - 1
		} else {
			if m[x-1][y-1] < m[x][y] {
				ops = append(ops, fmt.Sprintf("replace %d-th char %s of %q with %q", x, s1, s1[x-1], s2[y-1]))
			}
			x, y = x-1, y-1
		}
	}

	return m[len(s1)][len(s2)], ops
}
