package forFc

import (
	"fmt"
)

func ForString(nums int) {
	if nums <= 0 {
		println("不能0长度")
		return 
	}
	s := "Go is a beautiful language!"
	for i := 0; i < nums; i++ {
		fmt.Printf("Character on position %d is: %c \n", i, s[i])
	}
}