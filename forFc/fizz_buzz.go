package forFc

import (
	"fmt"
)

func FbFc(len int) {
	for i := 1; i <= len; i++ {
		switch true {
		case i%3==0&&i%5==0: fmt.Println("FizzBuzz")
		case i%3==0: fmt.Println("Fizz")
		case i%5==0: fmt.Println("Buzz")
		default: fmt.Println(i)
		}
	}
}