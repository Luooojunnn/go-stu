package funcv

import (
	"fmt"
)

// 非命名函数
func ReturnVal(num1 int, num2 int)(int, int, int) {
	fmt.Println(num1+num2, num1*num2, num1-num2)
	return num1+num2, num1*num2, num1-num2
}

func ReturnVal2(num1 int, num2 int) (he int, ji int, cha int) {
	fmt.Println(num1+num2, num1*num2, num1-num2)
	he = num1+num2
	ji = num1*num2
	cha = num1-num2
	return
}