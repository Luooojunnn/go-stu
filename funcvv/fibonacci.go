package funcv 

import (
	"fmt"
)

func Fibonacci1(){
	r := 0
	for i := 0; i < 10; i++ {
		r = fibonacci(i)
		fmt.Printf("fibonacci(%d) is : %d \n", i, r)
	}
}
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return 
}
