package vazifalar

import "fmt"

func EvenArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			fmt.Println("Juftlari= ", i)
		}
	}
	return arr
}
