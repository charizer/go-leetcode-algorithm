package check_if_it_is_a_straight_line

func CheckStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 3 {
		return true
	}
	// (x1 - x0) / (y1 - y0) = (x - x1) / (y - y1)
	// (x1 - x0) * (y - y1) = (y1 - y0) * (x - x1)
	// dx * (y - y1) = dy * (x - x1)
	x0, y0, x1, y1 := coordinates[0][0], coordinates[0][1], coordinates[1][0], coordinates[1][1]
	dx := x1 - x0
	dy := y1 - y0
	for i := 2; i < len(coordinates); i++ {
		if dx*(coordinates[i][1]-coordinates[i-1][1]) != dy*(coordinates[i][0]-coordinates[i-1][0]) {
			return false
		}
	}
	return true
}
