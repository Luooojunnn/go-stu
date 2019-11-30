package pointere

import (
	"fmt"
)

func Po() {
	var p  int = 1
	// *p = 1
	fmt.Println(*&p)
	// var i1 = 132
	// fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	// var i1p *int
	// i1p = &i1
	// fmt.Printf("The value at memory location %p is %d\n", i1p, *i1p)
}