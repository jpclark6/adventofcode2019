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
	part1Amps := []int{0, 1, 2, 3, 4}
	feedback := false
	ans := findMaxSignal(part1Amps, feedback)
	fmt.Println("Answer to part 1:", ans)

	midTime := time.Now()

	part2Amps := []int{5, 6, 7, 8, 9}
	feedback = true
	ans = findMaxSignal(part2Amps, feedback)
	fmt.Println("Answer to part 2:", ans)

	endTime := time.Now()
	fmt.Println("Elapsed time for part 1:", midTime.Sub(startTime))
	fmt.Println("Elapsed time for part 2:", endTime.Sub(midTime))
	fmt.Println("Elapsed time for both parts:", endTime.Sub(startTime))
}

func getProgram() []int {
	content, err := ioutil.ReadFile("./puzzledata/7day.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := string(content)
	rawProgram := strings.Split(data, ",")
	var programInt []int

	for _, s := range rawProgram {
		n, _ := strconv.Atoi(s)
		programInt = append(programInt, n)
	}
	return programInt[0:len(programInt) - 1]
}

// func runIntCode(input []int, program []int) (int, []int) {
// 	var output int
// 	i := 0
// Loop:
// 	for {
// 		instruction := append([]string{"0", "0", "0"}, strings.Split(strconv.Itoa(program[i]), "")...)
// 		instructionOpcode := instruction[len(instruction)-1]
// 		var valueOne, valueTwo, loc int

// 		if strings.Contains("125678", instructionOpcode) {
// 			loc = program[i+3]
// 			modeOne := instruction[len(instruction)-3]
// 			modeTwo := instruction[len(instruction)-4]

// 			if modeOne == "0" {
// 				valueOne = program[program[i+1]]
// 			} else {
// 				valueOne = program[i+1]
// 			}

// 			if modeTwo == "0" {
// 				valueTwo = program[program[i+2]]
// 			} else {
// 				valueTwo = program[i+2]
// 			}
// 		}

// 		switch instructionOpcode {
// 		case "1":
// 			program[loc] = valueOne + valueTwo
// 			i += 4
// 		case "2":
// 			program[loc] = valueOne * valueTwo
// 			i += 4
// 		case "3":
// 			var nextInput int
// 			if len(input) > 1 {
// 				nextInput, input = input[0], input[1:]
// 			} else {
// 				nextInput = input[0]
// 			}
// 			program[program[i+1]] = nextInput
// 			i += 2
// 		case "4":
// 			mode := instruction[len(instruction)-3]

// 			if mode == "0" {
// 				valueOne = program[program[i+1]]
// 			} else {
// 				valueOne = program[i+1]
// 			}
// 			// fmt.Println("Output:", valueOne)
// 			output = valueOne
// 			i += 2
// 		case "5":
// 			if valueOne != 0 {
// 				i = valueTwo
// 			} else {
// 				i += 3
// 			}
// 		case "6":
// 			if valueOne == 0 {
// 				i = valueTwo
// 			} else {
// 				i += 3
// 			}
// 		case "7":
// 			if valueOne < valueTwo {
// 				program[loc] = 1
// 			} else {
// 				program[loc] = 0
// 			}
// 			i += 4
// 		case "8":
// 			if valueOne == valueTwo {
// 				program[loc] = 1
// 			} else {
// 				program[loc] = 0
// 			}
// 			i += 4
// 		case "9":
// 			// fmt.Println("Diagnostic complete:", output)
// 			return output, program
// 		default:
// 			println("Something went wrong... Found case:", instructionOpcode, "loc", loc)
// 			break Loop
// 		}
// 	}
// 	return 0, program
// }

func runIntCodeRepeat(input []int, program []int, loc int) (int, []int, bool, int) {
	var output int
	i := loc
Loop:
	for {
		instruction := append([]string{"0", "0", "0"}, strings.Split(strconv.Itoa(program[i]), "")...)
		instructionOpcode := instruction[len(instruction)-1]
		var valueOne, valueTwo int

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
			var nextInput int
			if len(input) > 1 {
				nextInput, input = input[0], input[1:]
			} else if len(input) == 1{
				nextInput, input = input[0], nil
			} else {
				// fmt.Println("No more inputs currently")
				return output, program, false, i
			}
			program[program[i+1]] = nextInput
			i += 2
		case "4":
			mode := instruction[len(instruction)-3]

			if mode == "0" {
				valueOne = program[program[i+1]]
			} else {
				valueOne = program[i+1]
			}
			// fmt.Println("Output:", valueOne)
			output = valueOne
			i += 2
			// return output, program, false, i
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
			// println("Diagnostic complete:", output)
			return output, program, true, i
		default:
			println("Something went wrong... Found case:", instructionOpcode, i, program[i-1])
			break Loop
		}
	}
	return 0, program, false, 0
}

func amplifySignal(ampValues []int, feedback bool) int {
	currentProgram := getProgram()
	qtyAmplifiers := len(ampValues)
	var programs [][]int
	var programLocations []int

	for i := 0; i < qtyAmplifiers; i++ {
		programs = append(programs, getProgram())
		programLocations = append(programLocations, 0)
	}
	// fmt.Println("Program", getProgram())

	currentOutput := 0
	finished := false

	for i := 0; ; i++{
		// fmt.Println("i:", i, currentOutput)
		nextAmpValue := ampValues[i % qtyAmplifiers]
		var programInputs []int
		if i < qtyAmplifiers {
			programInputs = []int{nextAmpValue, currentOutput}
		} else {
			programInputs = []int{currentOutput}
		}

		if !feedback {
			currentOutput, currentProgram, finished, programLocations[0] = runIntCodeRepeat(programInputs, currentProgram, 0)
		} else {
			// fmt.Println("Before. programInputs:", programInputs, "program", programs[i % qtyAmplifiers], "location", programLocations[i % qtyAmplifiers])
			currentOutput, programs[i % qtyAmplifiers], finished, programLocations[i % qtyAmplifiers] = runIntCodeRepeat(programInputs, programs[i % qtyAmplifiers], programLocations[i % qtyAmplifiers])
			// fmt.Println("After. programInputs:", programInputs, "program", programs[i % qtyAmplifiers], "location", programLocations[i % qtyAmplifiers])
			// fmt.Println(programs, programLocations)
			// fmt.Println()
		}
		if !feedback && i == qtyAmplifiers - 1 {
			return currentOutput
		}
		if finished && i % qtyAmplifiers == qtyAmplifiers - 1 {
			return currentOutput
		}
		// if i > 3000 {
		// 	return 666
		// }
	}
}

func findMaxSignal(amplifiers []int, feedback bool) int {
	max := 0
	allCombos := permutations(amplifiers)
	for i:= 0; i < len(allCombos); i++ {
		// fmt.Println(allCombos[i])
		output := amplifySignal(allCombos[i], feedback)
		if output > max {
			// fmt.Println(max, allCombos[i])
			max = output
		}
		// break
	}
	return max
}

func permutations(arr []int) [][]int {
    var helper func([]int, int)
    res := [][]int{}

    helper = func(arr []int, n int) {
        if n == 1{
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++ {
                helper(arr, n - 1)
                if n % 2 == 1 {
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
}