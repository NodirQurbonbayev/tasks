package vazifalar

import "math"

func Isprime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))); i++  {
		if n%i==0 {
			return false
		}
	}
	if n<=1 {
		return false
	}else {
		return true
	}
}