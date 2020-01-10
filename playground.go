package main

import (
	"fmt"
	// "strconv"
	"strings"
)

func main() {
	x := "hello"
	fmt.Println(strings.Index(x, "j"))
}

func copy(li []int) []int {
	newlist := []int{}
	for i := 0; i < len(li); i++ {
		newlist = append(newlist, li[i])
	}
	return newlist
}