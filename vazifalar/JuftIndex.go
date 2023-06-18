package vazifalar

import "fmt"

func JuftlarIndexVaSoni(massiv []int) (juftlar []int, juftlarSoni int) {
	for i := 0; i < len(massiv); i += 2 {
		fmt.Printf("%d ", massiv[i])
		juftlar = append(juftlar, i)
	}
	juftlarSoni = len(juftlar)
	return
}
