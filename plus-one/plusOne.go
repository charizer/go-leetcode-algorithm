package plus_one

func PlusOne(digits []int) []int {
	for i:= len(digits)-1; i>=0; i++ {
		digits[i]++
		if digits[i] / 10 <= 0 {
			return digits
		}
		digits[i] %= 10
	}
	digits = append([]int{1}, digits...)
	return digits
}
