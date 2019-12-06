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
	fmt.Println("Part 1:")
	runDiagnostic(3)
	midTime := time.Now()
	fmt.Println("\nPart 2:")
	runDiagnostic(5)
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

func runDiagnostic(input int) {
	program := getProgram()
	i := 0
Loop:
	for {
		instruction := append([]string{"0", "0", "0"}, strings.Split(strconv.Itoa(program[i]), "")...)

		instructionOpcode := instruction[len(instruction)-1]
		var valueOne, valueTwo, loc int

		if strings.Contains("125678", instructionOpcode) {
			loc = program[i+3]
			modeOne := instruction[len(instruction)-3]
			modeTwo := instruction[len(instruction)-4]

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

		switch instructionOpcode {
		case "1":
			program[loc] = valueOne + valueTwo
			i += 4
		case "2":
			program[loc] = valueOne * valueTwo
			i += 4
		case "3":
			program[program[i+1]] = input
			i += 2
		case "4":
			mode := instruction[len(instruction)-3]

			if mode == "0" {
				valueOne = program[program[i+1]]
			} else {
				valueOne = program[i+1]
			}
			fmt.Println("Output:", valueOne)
			i += 2
		case "5":
			if valueOne != 0 {
				i = valueTwo
			} else {
				i += 3
			}
		case "6":
			if valueOne == 0 {
				i = valueTwo
			} else {
				i += 3
			}
		case "7":
			if valueOne < valueTwo {
				program[loc] = 1
			} else {
				program[loc] = 0
			}
			i += 4
		case "8":
			if valueOne == valueTwo {
				program[loc] = 1
			} else {
				program[loc] = 0
			}
			i += 4
		case "9":
			println("Diagnostic complete")
			break Loop
		default:
			println("Something went wrong... Found case:", instructionOpcode)
			break Loop
		}
	}
}
