package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

func main() {
	path := []int{1, 2, 3, 4}
	x := path[:len(path)-1]
	fmt.Println(x)
}