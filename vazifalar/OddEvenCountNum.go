package vazifalar

import "fmt"

func EvenCount(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			fmt.Printf("Juftlari= %d\n",arr[i])
		}
	}
	return arr
}
func Oddcount(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			fmt.Printf("Toqalari= %d\n",arr[i])
		}
	}
	return arr
}
