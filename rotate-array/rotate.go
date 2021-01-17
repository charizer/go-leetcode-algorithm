package rotate_array

/*
原始数组                  : 1 2 3 4 5 6 7
反转所有数字后             : 7 6 5 4 3 2 1
反转前 k 个数字后          : 5 6 7 4 3 2 1
反转后 n-k 个数字后        : 5 6 7 1 2 3 4 --> 结果
*/

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func Rotate(nums []int, k int)  {
	k %= len(nums)
	reverse(nums, 0, len(nums) -1)
	reverse(nums, 0, k -1)
	reverse(nums, k, len(nums) - 1)
}


func Rotate2(nums []int, k int)  {
	k %= len(nums)
	start, count, n := 0, 0, len(nums)
	for ; count < len(nums); start++ {
		cur := start
		pre := nums[cur]
		for  {
			// 计算新位置
			next := (cur + k) % n
			// 先把要被占位的记录下来
			tmp := nums[next]
			// 占位
			nums[next] = pre
			// 被占位的继续处理
			pre = tmp
			cur = next
			count++
			// 回到起始，一轮换完
			if start == cur {
				break
			}
		}
	}
}

//  a b c d | e f  2
//  e f | a b c d
//  d c b a | f e
//  e f a b c d

// d c    b a

//  | b a d   c

//  c b a d

func reverse3(nums []int, start, end int){
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func Rotate3 (nums []int, k int) {
	k %= len(nums)
	reverse3(nums, 0, len(nums)-k-1)
	reverse3(nums, len(nums)-k, len(nums)-1)
	reverse3(nums, 0, len(nums)-1)
}

func Rotate4 (nums []int, k int) {
	n := len(nums)
	k = k % n
	start, count := 0, 0
	for ; count < len(nums); start++ {
		cur := start
		pre := nums[cur]
		for {
			// 当前需要移动到的位置
			next := (cur + k) % n
			// 把被占用的位置的数据先暂存
			tmp := nums[next]
			// 当前元素放到目标位置
			nums[next] = pre
			pre = tmp
			cur = next
			count++
			// 回到了起点，需要开始新一轮
			if cur == start {
				break
			}
		}
	}
}