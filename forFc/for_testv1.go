package forFc

import (
	"fmt"
)

func Testv1() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d", v)
		v = 5
	}
}

func Testv2() {
	for i := 0; ; i++ {
		fmt.Println("Value of i is now:", i)
	}
}

func Testv3() {
	for i := 0; i < 3; {
		fmt.Println("value is now:", i)
	}
}

func Testv4() {
	// s := ""
	// for s != "aaaa" {
	// 	fmt.Println(s)
	// 	s += "a"
	// }
	for i := 0; i<7 ; i++ {
		if i%2 == 0 { continue }
		fmt.Println("Odd:", i)
	}
}