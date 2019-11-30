package forFc

import (
	"fmt"
)

func Ran(){
	str := "Go is a beautiful language!"

	fmt.Printf("str 的len是 %d", len(str))

	for i, key := range str {
		fmt.Printf("Character on position %d is: %c \n", i, key)
	}
}