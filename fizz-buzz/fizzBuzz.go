package fizz_buzz

import "strconv"

func FizzBuzz(n int) []string {
	ans := []string{}
	if n <= 0 {
		return ans
	}
	for i := 1 ; i <= n; i++ {
		if i % 15 == 0 {
			ans = append(ans, "FizzBuzz")
		}else if i % 3 == 0 {
			ans = append(ans, "Fizz")
		}else if i % 5 == 0 {
			ans = append(ans, "Buzz")
		}else{
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return ans
}
