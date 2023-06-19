package vazifalar

func ArraySum(arr []int, a, b int) int {
	sum := 0
	for i := a - 1; i <= b; i++ {
		sum += arr[i]
	}
	return sum
}
