package main

import (
	"fmt"
	"tasks/vazifalar"
)

func main() {
	//kabisa
	fmt.Println(vazifalar.Kabisa(1200))
	//Swapper
	fmt.Println(vazifalar.Swapper(4, 6))
	//DayMonth
	fmt.Println(vazifalar.DayMonth(5, 6))
	//ISprime
	fmt.Println(vazifalar.Isprime(14))
	//Ishora
	a := 5
	b := 0
	fmt.Println(vazifalar.Ishora(a) + vazifalar.Ishora(b))
	//Kvtenglama
	fmt.Println(vazifalar.Kvtenglama(2, 4, 5))
	//DoiraYuzi
	r1 := 23
	r2 := 22
	r3 := 21
	s1 := vazifalar.DoiraYuzi(float64(r1))
	s2 := vazifalar.DoiraYuzi(float64(r2))
	s3 := vazifalar.DoiraYuzi(float64(r3))
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	//AylanaYuzi
	fmt.Println(vazifalar.AylanaYuzi(4, 5))
	//Triangle
	fmt.Println(vazifalar.Triangle(3, 4))
	//SumRange
	fmt.Println(vazifalar.SumRange(5, 2, 8))
	//EvenOdd
	a = 5
	b = 7
	c := 10
	fmt.Println(vazifalar.EvenOdd(a))
	fmt.Println(vazifalar.EvenOdd(b))
	fmt.Println(vazifalar.EvenOdd(c))
	//SumArray
	fmt.Println(vazifalar.SumArray(2, 3, 4))
	//TeskariCh
	n := []int{1, 2, 3, 4, 5}
	vazifalar.ArrayTeskari(n)
	//ToqIndex
	arr := []int{4, 5, 6, 7, 69}
	odds, _ := vazifalar.OddCountNum(arr)
	for _, v := range odds {
		fmt.Println("Toqlari= ", arr[v])
	}
	_, count := vazifalar.OddCountNum(arr)
	fmt.Println("Toqlari soni= ", count)
	fmt.Println("/////////////////////////////////////////////////////////////////////////")
	//EvenIndex
	massiv := []int{2, 3, 4, 6, 8, 89}
	even, _ := vazifalar.EvenCountNum(massiv)
	for _, v := range even {
		fmt.Println("Juftlari= ", massiv[v])
	}
	_, EvenCount := vazifalar.EvenCountNum(massiv)
	fmt.Println("Juftlar soni= ", EvenCount)
	//EvenOddIndex
	sum := []int{4, 5, 7, 8, 69}
	fmt.Println(vazifalar.EvenCount(sum))
	fmt.Println(vazifalar.Oddcount(sum))
	//KarraliArray
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 5

	fmt.Print(vazifalar.KarraliArray(array, k))
	//EvenArray
	son := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(vazifalar.EvenArray(son))
	//ArraySum
	Easy := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	j := 2
	l := 7
	medium := vazifalar.ArraySum(Easy, j, l)
	fmt.Println(medium)

}
