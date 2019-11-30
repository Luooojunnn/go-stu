package arr

import "fmt"

func ForFc() {
	sl_from := []int{1, 2, 3}
	sl_to := []int{10,20,0,0}

	n := copy(sl_to[2:], sl_from)
	fmt.Println(sl_to)
	fmt.Println(n)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, sl_to...)
	fmt.Println(sl3)
}