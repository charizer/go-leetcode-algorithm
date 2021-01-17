package canPlaceFlowers

func CanPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	for i := 0; i < len(flowerbed); i+=2 {
		if flowerbed[i] == 0 {
			if i == len(flowerbed) - 1 || flowerbed[i + 1] == 0 {
				n--
			} else {
				i++
			}
		}
	}
	return n <= 0
}

func CanPlaceFlowers2(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	count := 0
	for i := 0; i < len(flowerbed) && count < n ; i++ {
		if flowerbed[i] == 0 {
			pre, next := 0, 0
			if i == 0 {
				pre = 0
			}else {
				pre = flowerbed[i-1]
			}
			if i == len(flowerbed) - 1 {
				next = 0
			}else{
				next = flowerbed[i+1]
			}
			if pre == 0 && next == 0 {
				count++
				flowerbed[i] = 1
 			}
		}
	}
	return count >= n
}
