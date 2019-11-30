package funcv

import (
	"fmt"
)

var a string = "123"

func T1() {
	// 0x115c228
	t(a)
	// a= 456
	// fmt.Println(&a)
}

func t(a string) {
	
	fmt.Println(&a)
}