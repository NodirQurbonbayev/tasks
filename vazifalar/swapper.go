package vazifalar

func Swapper(x, y int) (int, int) {
	if x > y {
		x,y =2*x*y, (x+y)/2
	}else if x<y{
		y,x=2*x*y,(x+y)/2
	}
	return x,y
}