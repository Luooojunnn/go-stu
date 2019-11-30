package forFc

import (
	"fmt"
)

func For2(){
	i := 5
	for i > 0 {
		i -= 1
		fmt.Println(i)
	}
}