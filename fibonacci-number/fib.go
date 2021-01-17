package fibonacci_number

import "fmt"

func Fib(n int) int {
	f := (1 - 5) / 2
	fmt.Printf("f:%d\n", f)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	f0,f1,f2 := 0, 1, 1
	for i:=3; i<=n; i++ {
		f0 = f1
		f1 = f2
		f2 = f0 + f1
	}
	return f2
}
