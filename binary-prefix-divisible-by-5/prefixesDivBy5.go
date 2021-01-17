package binary_prefix_divisible_by_5

func PrefixesDivBy5(A []int) []bool {
	num := 0
	ans := []bool{}
	for _, a := range A {
		num = (num * 2 + a) % 5
		if num == 0 {
			ans = append(ans, true)
		}else{
			ans = append(ans, false)
		}
	}
	return ans
}
