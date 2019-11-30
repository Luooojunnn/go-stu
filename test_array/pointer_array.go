package arr

import "fmt"

var a1 = []int{1,2,3,4,5}
var a2 = [...]int{10,20,30,40,50,60}

func PointerArr() {
	p1(a1)
	p2(&a2)
	fmt.Println(a2)
}

func p1(a []int) {
	fmt.Println(a)
}
func p2(a *[6]int) {
	a[0] = 100
	fmt.Println(a)
}