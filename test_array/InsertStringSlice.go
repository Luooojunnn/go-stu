package arr

import "fmt"

// i插入的开始位置
// arr1 目标数组
// arr2 源数组

// arr1 输入 1，2，3
func InsertStringSlice(i int, arr1 []int, arr2 []int){
	resArr := make([]int, (len(arr1)+len(arr2)))
	copy(resArr, arr1[0:i])
	resArr = append(append(resArr[0:i], arr2...), arr1[i:]...)
	fmt.Println(resArr)
}
