package vazifalar

import "math"

func Triangle(a, b float64) float64 {
	c := math.Sqrt(a*a+b*b)
	perimetr:=a+b+c
	return perimetr
}