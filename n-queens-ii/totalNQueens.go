package n_queens_ii

func TotalNQueens(n int) int {
	cols := make([]bool, n)
	d45 := make([]bool, 2 * n)
	d135 := make([]bool, 2 * n)
	count := 0
	backtrack(0, cols, d45, d135, n, &count)
	return count
}

func backtrack(row int, cols, d45, d135 []bool, n int, count *int) {
	if row == n {
		*count++
		return
	}
	for col := 0; col < n; col++ {
		id1 := row - col + n
		id2 := row + col
		if cols[col] || d45[id1] || d135[id2] {
			continue
		}
		cols[col], d45[id1], d135[id2] = true, true, true
		backtrack(row + 1, cols, d45, d135, n, count)
		cols[col], d45[id1], d135[id2] = false, false, false
	}
}
