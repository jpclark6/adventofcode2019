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
	// part2()
	endTime := time.Now()
	fmt.Println("Elapsed time for part 1:", midTime.Sub(startTime))
	fmt.Println("Elapsed time for part 2:", endTime.Sub(midTime))
	fmt.Println("Elapsed time for both parts:", endTime.Sub(startTime))
}

func getProgram() []int {
	content, err := ioutil.ReadFile("./puzzledata/5day.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := string(content)
	rawProgram := strings.Split(data, ",")
	programInt := make([]int, len(rawProgram))

	for i, s := range rawProgram {
		n, _ := strconv.Atoi(s)
		programInt[i] = n
	}
	return programInt
}

func runDiagnostic() {
	program := getProgram()
	i := 0
Loop:
	for {
		instruction := append([]string{"0", "0", "0"}, strings.Split(strconv.Itoa(program[i]), "")...)

		instructionOpcode := instruction[len(instruction) - 1]
		var valueOne, valueTwo, loc int

		if instructionOpcode == "1" || instructionOpcode == "2" || instructionOpcode == "4" {
			loc = program[i + 3]
			modeOne := instruction[len(instruction) - 3]
			modeTwo := instruction[len(instruction) - 4]
	
			if modeOne == "0" {
				valueOne = program[program[i+1]]
			} else {
				valueOne = program[i+1]
			}

			if modeTwo == "0" {
				valueTwo = program[program[i+2]]
			} else {
				valueTwo = program[i+2]
			}
		}

		fmt.Println("Instruction:", instruction, "Opcode:", instructionOpcode)

		switch instructionOpcode {
		case "1":
			fmt.Println("Op1. Current loc:", i, "Writing", valueOne, "+", valueTwo, "to location", loc)
			program[loc] = valueOne + valueTwo
			i += 4
		case "2":
			fmt.Println("Op1. Current loc:", i, "Writing", valueOne, "*", valueTwo, "to location", loc)
			program[loc] = valueOne * valueTwo
			i += 4
		case "3":
			fmt.Println("Op3. Writing 1 to location", program[i + 1])
			program[program[i + 1]] = 1
			i += 2
		case "4":
			fmt.Println("Opcode4. Loc:", i + 1, "Value:", valueOne)
			fmt.Println("Output:", valueOne)
			i += 2
		case "9":
			println("Diagnostic complete")
			break Loop
		default:
			println("Something went wrong... Found case:", instructionOpcode)
			break Loop
		}
		fmt.Println()
	}
}

func part1() {
	runDiagnostic()
}

// func part2() {
// 	expectedOutput := 19690720
// Loop:
// 	for noun := 0; noun < 100; noun++ {
// 		for verb := 0; verb < 100; verb++ {
// 			output := fixProgram(noun, verb)
// 			if output == expectedOutput {
// 				fmt.Println("Answer to part 2:", 100 * noun + verb)
// 				break Loop
// 			}
// 		}
// 	}
// }