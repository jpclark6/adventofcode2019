package main

import (
	"fmt"
	// "strings"
)

func main() {
	x := 0
	for i := 0; i < 10; i++ {
		if i == 5 {
			x = i
		}
	}
	fmt.Println(x)
}