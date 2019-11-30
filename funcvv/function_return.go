package funcv

// import "fmt"

// 返回函数
func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

// 返回函数
func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}