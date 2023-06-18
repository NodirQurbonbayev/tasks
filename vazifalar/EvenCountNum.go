package vazifalar

func EvenCountNum(arr []int) (even []int, count int) {
	for i, v := range arr {
		if v%2 != 1 {
			even = append(even, i)
			count++
		}
	}
	return even, count
}
