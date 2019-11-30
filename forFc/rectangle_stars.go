package forFc

import (
	"fmt"
	"strings"
)

func Rectangle_stars() {
	s := "*"
	for i := 0; i < 10; i++ {
		fmt.Println(strings.Repeat(s, 20))
	}
}