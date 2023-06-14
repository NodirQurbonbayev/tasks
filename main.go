package main

import (
	"fmt"
	"tasks/vazifalar"
)

func main() {
	fmt.Println(vazifalar.Kabisa(1200))
	fmt.Println(vazifalar.Swapper(4,6))
	fmt.Println(vazifalar.DayMonth(5, 6))
	fmt.Println(vazifalar.Isprime(14))
	a := 5
    b := 0
    fmt.Println(vazifalar.Ishora(a) + vazifalar.Ishora(b))
	fmt.Println(vazifalar.Kvtenglama(2,4,5))
}