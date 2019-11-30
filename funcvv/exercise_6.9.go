package funcv


func Febona() func(int) int {
	num1 := 1 // 0号位初始值
	num2 := 1 // 1号位初始值
	temp := 0 // 交换位
	return func(num int)(int) {
		if num <= 1 {
			return 1
		}
		for i := 1; i < num; i++ {
			temp = num2
			num2 = num1 + num2
			num1 = temp
		}
		return num2
	}
}