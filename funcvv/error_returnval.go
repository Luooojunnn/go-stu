package funcv

import (
	"fmt"
	"math"	
)

// 非命名函数
func MySqrt(num float64) (float64, bool) {
	if num < 0 {
		panic("不能输入小于0的数值")
	}
	fmt.Println(math.Sqrt(num))
	return math.Sqrt(num), false
}

// 命名函数
func MySqrt2(num float64) (num2 float64, e bool) {
	if num < 0 {
		panic("不能输入小于0的数值")
	}
	num2 = math.Sqrt(num)
	e = true
	return 
}