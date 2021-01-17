package intersection_of_two_arrays_ii

func Intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	// 记录元素出现的次数
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num]++
	}
	for _, num := range nums2 {
		if v, ok := m[num]; ok {
			// 在nums1中出现的元素在nums2也出现， 并将记录次数-1，防止重复
			result = append(result, num)
			m[num] = v - 1
			// 次数为0时删除
			if m[num] == 0 {
				delete(m, num)
			}
		}
	}
	return result
}
