package funcv

import (
	"log"
	"io"
)

func Logvalues(s string) (n int, e error) {
	defer func(){
		log.Printf("Logvalues(%q) = %d, %v", s, n, e)
	}()
	return 7, io.EOF
}