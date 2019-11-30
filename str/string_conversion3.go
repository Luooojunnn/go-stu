package str

import (
	"fmt"
)

func Main3() {
	v, e := a(3)
	if (e) {
		fmt.Print("你好")
	} else {
		fmt.Print(v)
	}
	// fmt.Print(a(33))
}

func a(num float32)(n float32, b bool) {
	return num/2, true
}