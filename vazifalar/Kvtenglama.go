package vazifalar

import (
	"math"
)

func Kvtenglama(a, b, c float64) (x1, x2 float64) {
	D := b*b - 4*a*c
	if D < 0 {
		return math.NaN(), math.NaN() //
	} else if D == 0 {
		x1 = -b / (2 * a)
		x2 = x1
	} else {
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
	}

	return x1, x2
}
