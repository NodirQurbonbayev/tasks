package vazifalar

import "fmt"

func ToqlarIndexVaSoni(massiv []int) (toqlar []int, toqlarSoni int) {
	for i := 1; i < len(massiv); i += 2 {
		fmt.Printf("%d ", massiv[i])
		toqlar = append(toqlar, i)
	}
	toqlarSoni = len(toqlar)
	return
}
