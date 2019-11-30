package funcv

import (
	"fmt"
)

func com(a, b int, reply *int) {
	// 反引用
	*reply  = a * b
}

func Multiply() {
	n := 0
	com(10, 5, &n)
	fmt.Println(n)
}