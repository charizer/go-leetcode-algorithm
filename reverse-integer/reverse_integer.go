package reverse_integer

import "math"

//指定为int超过32位时不会溢出，按照leetcode上的要求就会出错
func Reverse(x int) int {
	result := 0
	for x != 0 {
		tail := x % 10
		newResult := result * 10 + tail
		if (newResult - tail) / 10 != result {
			return 0
		}
		result = newResult
		x = x / 10
	}
	return result
}


// int is a signed integer type that is at least 32 bits in size. It is a
// distinct type, however, and not an alias for, say, int32.
// type int int
// int不是int32， int的范围比int32要大
func Reverse3(x int32) int32 {
	var result int32 =  0
	var newResult int32
	for x != 0 {
		tail := x % 10
		newResult = result * 10 + tail
		if (newResult - tail) / 10 != result {
			return 0
		}
		result = newResult
		x = x / 10
	}
	return result
}

func Reverse2(x int) int {
	result := 0
	for x != 0 {
		if result > math.MaxInt32 / 10 || result == math.MaxInt32 / 10 && x % 10 > math.MaxInt32 % 10 ||
			result < math.MinInt32 / 10 || result == math.MinInt32 && x % 10 < math.MaxInt32 % 10 {
			return 0
		}
		result = result * 10 + x % 10
		x = x / 10
	}
	return result
}
