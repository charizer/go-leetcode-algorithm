package add_digits

func AddDigits(num int) int {
	if num < 10 {
		return num
	}
	return AddDigits(num%10 + num/10)
}


