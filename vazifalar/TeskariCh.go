package vazifalar

import "fmt"

func ArrayTeskari(n []int) {
	for i := len(n) - 1; i >= 0; i-- {
		fmt.Println(n[i])
	}
}
