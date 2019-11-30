package arr

import "fmt"

func Filter(s []int) {
	for _, v := range s {
		fn(v)
	}
}

func fn(s int) {
	if s%2 == 0 {
		fmt.Printf("%d is 偶数\n", s)
	}
}
