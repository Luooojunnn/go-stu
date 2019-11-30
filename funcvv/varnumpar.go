package funcv

import (
	"fmt"
)

func Main() {
	min := minFc(2,3,4,12,1)
	fmt.Println(min)
	// sliceNums := []string{"0.82","8.21","88,27","0.22"}
	min = minFc([]int{2,3,1}...)
	fmt.Println(min)
}

func minFc(nums ...int) (int) {
	if len(nums) == 0 {
		return 0
	}
	minNum := nums[0]
	for _, v := range nums {
		if v < minNum {
			minNum = v
		}
	}
	return minNum
}