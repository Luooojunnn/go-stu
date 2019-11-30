package arr 

import "fmt"


func SliceTest() {
	s := []byte{'n', 'i', 'h', 'a', 'o'}
	fmt.Println(s[:])
	fmt.Println(s[2:])
	fmt.Println(s[:4])
	fmt.Println(s[2:4])
}

// 输入一个num，返回长度为num的斐波那契切片
func CreateSlice(num int) (f []int) {
	f = make([]int, num)
	if num <= 0 {
		f = []int{}
		return 
	}

	if num == 1 {
		f = []int{1}
		return 
	}

	if num == 2 {
		f = []int{1,1}
		return 
	}

	f[0] = 1
	f[1] = 1
	n := 0

	for i := 2; i < num; i++ {
		n = f[i-1]
		f[i] = f[i-1] + f[i-2]
		f[i-1] = n
	}
	return 
}