package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "z"

	nst := append([]string{"0", "0", "0"}, strings.Split(x, "")...)
	fmt.Println("X:", nst[len(nst) - 1])

}