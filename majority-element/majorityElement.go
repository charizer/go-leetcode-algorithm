package majority_element

func MajorityElement(nums []int) int {
	max := 0
	ans := 0
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] > max {
			max = m[num]
			ans = num
		}
 	}
 	return ans
}


func MajorityElement2(nums []int) int {
	ans, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == ans {
			count++
		}else{
			count--
			if count == 0 {
				ans = nums[i]
				count = 1
			}
		}
	}
	return ans
}