package swi

import (
	"fmt"
)
// 需要导出的函数必须用大写开头
func SwFc(num int) {
	switch num {
		case 1: fallthrough
		case 2: fmt.Println("输入了2")
	case 3:
		fmt.Println("输入了3")
	default:
		a()
	}
}

func a(){
	panic(1)
}