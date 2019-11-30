package funcv

import "fmt"

// 练习 6.4


// 生成斐波那契数列的程序,数列中的位置和对应的值
func Febo(num int) {
	for i := 0; i < num; i++ {
		fmt.Printf("斐波那契额数列%d的位置是%d\n", febo(i), i )
	}
}
func febo(num int) (int) {
	if num <= 1 {
		return 1
	}
	return febo(num - 1) + febo(num - 2)
}

// 阶乘函数
func Jiec(num []int) {
	for _,v := range num {
		fmt.Printf("%d的阶乘的值是：%d\n", v, jiec(v))
	}
}
func jiec(num int) int {
	if num <= 1 {
		return 1
	}

	return jiec(num - 1) * num
}

