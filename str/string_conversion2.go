package str

import (
	"fmt"
	"strconv"
)

func Main2() {
	var orig string = "123"
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, err := strconv.Atoi(orig)

	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return 
	}

	fmt.Printf("the interge is %d\n", an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("the new string is %s\n", newS)
}