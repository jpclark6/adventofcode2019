package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	content, err := ioutil.ReadFile("./puzzledata/2day.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := string(content)
	rawProgram := strings.Split(data, ",")
	program := make([]int, len(rawProgram))

	for i, s := range rawProgram {
		n, _ := strconv.Atoi(s)
		program[i] = n
	}

	program[1] = 12
	program[2] = 2

	i := 0

Loop:
	for {
		switch program[i] {
		case 1:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
		case 2:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
		case 99:
			break Loop
		default:
			println("Something went wrong...")
			break Loop
		}
		i += 4
	}
	fmt.Println("Solution to part 1:", program[0])

	end := time.Now()
	fmt.Println("Elapsed time for part 1:", end.Sub(startTime))
}
