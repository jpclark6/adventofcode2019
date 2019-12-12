package main

import (
	"fmt"
	// "strings"
)

func main() {
	var x = map[int]bool{
		1: true,
		2: true,
	}

	if x[1] {
		fmt.Println("Yes")
	}
	if x[5] {
		fmt.Println("No")
	}
}