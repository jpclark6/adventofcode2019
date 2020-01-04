package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

func main() {
	x := []int{1, 2, 3, 4}
	y := copy(x)
	y[1] = 444
	fmt.Println(x, y)
}

func copy(li []int) []int {
	newlist := []int{}
	for i := 0; i < len(li); i++ {
		newlist = append(newlist, li[i])
	}
	return newlist
}