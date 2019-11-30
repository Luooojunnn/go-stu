/** 使用defer检测代码的执行 */

package funcv

import (
	"fmt"
)

func tr(s string) string {
	fmt.Println("进入", s)
	return s
}
func un(s string) {
	fmt.Println("离开", s)
}

func A(){
	defer un(tr("A"))
	fmt.Println("A的一些操作")
}
func B(){
	defer un(tr("B"))
	fmt.Println("B的一些操作")
	A()
}