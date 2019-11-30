package forFc

import (
	"fmt"
	"strings"
)

func  For_character(){
	for i := 0; i < 25; i++ {
		for ii := 0; ii < i+1; ii++ {
			fmt.Print("c")
		}
		fmt.Println()
 	}
}

func For_character2() {
	s := ""
	for i := 0; i < 25; i++ {
		s += strings.Repeat("=", i)
		s += "\r\n"
	}
	fmt.Println(s)
}