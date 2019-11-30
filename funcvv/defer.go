/**
* defer 给我的感觉就是异步执行了
*/
package funcv

import (
	"fmt"
)

func Fun1() {
	i := 0
	fmt.Println("Fun1开始执行")
	defer fmt.Println(i)
	i = 10
	defer fun2() // 这里就异步了
	fmt.Println("Fun1结束执行")
	return 
}

func fun2() {
	fmt.Println("Fun2开始执行")
}

