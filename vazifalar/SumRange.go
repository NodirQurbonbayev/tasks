package vazifalar

func SumRange(x, y, z int) int {
	sum := 0
	if x > y {
		return 0
	}
	for i := x; i <= y; i++ {
		for j := y; j <= z; j++ {
			sum += j + i
		}
	}
	return sum
}
