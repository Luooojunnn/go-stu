package arr

import "fmt"

func ForArr() {
	// int类型数组，数组的长度是10
	var arr [10]int

	// 改变数组的初始值
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 3
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for k,v := range arr {
		fmt.Println(k,v)
	}
}