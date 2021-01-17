package eraseOverlapIntervals

func EraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 0 {
		return 0
	}
	removes := 0
	for i:=0; i< len(intervals)-1;i++ {
		for j := i+1; j < len(intervals); j++ {
			if intervals[i][0] <= intervals[j][0] && intervals[i][1] >= intervals[j][1] {
				removes++
				break
			}else if intervals[i][0] <= intervals[j][0] && intervals[i][1] <= intervals[j][1] &&
				(intervals[i][1] - intervals[i][0] > intervals[j][1]- intervals[j][0]) {
				removes++
				break
			}else if intervals[i][0] >= intervals[j][0] && intervals[i][1] >= intervals[j][1] &&
				(intervals[i][1] - intervals[i][0] > intervals[j][1]- intervals[j][0]) {
				removes++
			}
		}
	}
	return removes
}
