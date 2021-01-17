package n_queens

import "testing"

func TestSolveNQueens(t *testing.T) {
	n := 4
	want := [][]string{
		{
			".Q..",
			"...Q",
			"Q...",
			"..Q.",
		},
		{
			"..Q.",
			"Q...",
			"...Q",
			".Q..",
		},
	}
	got := SolveNQueens(n)
	for idx, w := range want {
		for i := range w {
			if got[idx][i] != want[idx][i] {
				t.Errorf("want:%s got:%s", want[idx][i], got[idx][i])
			}
		}
	}
}
