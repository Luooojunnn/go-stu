package arr

import "fmt"

const (
	WIDTH = 1920
	HEIGHT = 1080
)

type pixel int 
var screen [WIDTH][HEIGHT]pixel

// 二维数组
func MultidimArr() {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			screen[j][i] = 0
		}
	}
	fmt.Println(screen)
}