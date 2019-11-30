package str

import (
	"fmt"
	"strconv"
	"../trans"
)

func Main() {
	var orig string = "666"
	var an int
	var newS string

	fmt.Println(strconv.IntSize)
	an, _ = strconv.Atoi(orig)
	fmt.Println(an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Println(newS)

	trans.Main()
}