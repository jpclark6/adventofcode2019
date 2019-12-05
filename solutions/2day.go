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
	part1()
	midTime := time.Now()
	part2()
	endTime := time.Now()
	fmt.Println("Elapsed time for part 1:", midTime.Sub(startTime))
	fmt.Println("Elapsed time for part 2:", endTime.Sub(midTime))
	fmt.Println("Elapsed time for both parts:", endTime.Sub(startTime))
}

func getProgram() []int {
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
	return program
}

func fixProgram(noun, verb int) int {
	program := getProgram()

	program[1] = noun
	program[2] = verb

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
	return program[0]
}

func part1() {
	ans := fixProgram(12, 2)
	fmt.Println("Solution to part 1:", ans)
}

func part2() {
	expectedOutput := 19690720
Loop:
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			output := fixProgram(noun, verb)
			if output == expectedOutput {
				fmt.Println("Answer to part 2:", 100 * noun + verb)
				break Loop
			}
		}
	}
}