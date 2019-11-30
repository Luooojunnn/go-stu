package funcv

import (
	"fmt"
)

func Mut_recurs() {
	fmt.Printf("%d is 偶数:%t\n", 16, even(16))
	fmt.Printf("%d is 奇数:%t\n", 17, odd(17))
	fmt.Printf("%d is 奇数:%t\n", 18, odd(18))
}

// 奇数判断函数
func odd(num int) bool {
	if num == 0 {
		return false 
	}
	return even(revSign(num) - 1)
}
// 偶数判断函数
func even(num int) bool {
	if num == 0 {
		return true
	}
	return odd(revSign(num) - 1)
}
// 负数处理函数
func revSign(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}