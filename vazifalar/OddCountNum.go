package vazifalar

func OddCountNum(arr []int) (odds []int,count int) {
	for i, v := range arr {
		if v%2 != 0 {
			odds = append(odds, i)
			count++
		}
	}
	return odds,count
}
