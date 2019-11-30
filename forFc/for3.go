// 无限for
package forFc

import (
	"fmt"
)

func ForInfi() {
	i := 1
	for {
		i+=1
		if i == 10 {
			break
		}
		fmt.Print(i)
	}
	fmt.Println("==")
}