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

func runIntCodeRepeat(input []int, program []int, cursorLocation int) (output int, newProgramState []int, finished bool, newLocation int) {
	var nextCursorLocation int
Loop:
	for {
		instruction := append([]string{"0", "0", "0"}, strings.Split(strconv.Itoa(program[cursorLocation]), "")...)
		instructionOpcode := instruction[len(instruction)-1]
		var valueOne, valueTwo int

		if strings.Contains("125678", instructionOpcode) {
			nextCursorLocation = program[cursorLocation+3]
			modeOne := instruction[len(instruction)-3]
			modeTwo := instruction[len(instruction)-4]

			if modeOne == "0" {
				valueOne = program[program[cursorLocation+1]]
			} else {
				valueOne = program[cursorLocation+1]
			}

			if modeTwo == "0" {
				valueTwo = program[program[cursorLocation+2]]
			} else {
				valueTwo = program[cursorLocation+2]
			}
		}

		switch instructionOpcode {
		case "1":
			program[nextCursorLocation] = valueOne + valueTwo
			cursorLocation += 4
		case "2":
			program[nextCursorLocation] = valueOne * valueTwo
			cursorLocation += 4
		case "3":
			var nextInput int
			if len(input) > 1 {
				nextInput, input = input[0], input[1:]
			} else if len(input) == 1{
				nextInput, input = input[0], nil
			} else {
				return output, program, false, cursorLocation
			}
			program[program[cursorLocation+1]] = nextInput
			cursorLocation += 2
		case "4":
			mode := instruction[len(instruction)-3]
			if mode == "0" {
				valueOne = program[program[cursorLocation+1]]
			} else {
				valueOne = program[cursorLocation+1]
			}
			output = valueOne
			cursorLocation += 2
		case "5":
			if valueOne != 0 {
				cursorLocation = valueTwo
			} else {
				cursorLocation += 3
			}
		case "6":
			if valueOne == 0 {
				cursorLocation = valueTwo
			} else {
				cursorLocation += 3
			}
		case "7":
			if valueOne < valueTwo {
				program[nextCursorLocation] = 1
			} else {
				program[nextCursorLocation] = 0
			}
			cursorLocation += 4
		case "8":
			if valueOne == valueTwo {
				program[nextCursorLocation] = 1
			} else {
				program[nextCursorLocation] = 0
			}
			cursorLocation += 4
		case "9":
			return output, program, true, cursorLocation
		default:
			println("Something went wrong... Found case:", instructionOpcode, cursorLocation, program[cursorLocation-1])
			break Loop
		}
	}
	return 0, program, false, 0
}

func amplifySignal(ampValues []int, feedback bool) int {
	qtyAmplifiers := len(ampValues)
	var programs [][]int
	var programLocations []int

	for i := 0; i < qtyAmplifiers; i++ {
		programs = append(programs, getProgram())
		programLocations = append(programLocations, 0)
	}

	currentOutput := 0
	finished := false

	for i := 0; ; i++{
		nextAmpValue := ampValues[i % qtyAmplifiers]
		var programInputs []int
		if i < qtyAmplifiers {
			programInputs = []int{nextAmpValue, currentOutput}
		} else {
			programInputs = []int{currentOutput}
		}

		if !feedback {
			currentOutput, programs[0], finished, programLocations[0] = 
				runIntCodeRepeat(programInputs, programs[0], 0)
		} else {
			currentOutput, programs[i % qtyAmplifiers], finished, programLocations[i % qtyAmplifiers] = 
				runIntCodeRepeat(programInputs, programs[i % qtyAmplifiers], programLocations[i % qtyAmplifiers])
		}
		if !feedback && i == qtyAmplifiers - 1 {
			return currentOutput
		}
		if finished && i % qtyAmplifiers == qtyAmplifiers - 1 {
			return currentOutput
		}
	}
}

func findMaxSignal(amplifiers []int, feedback bool) int {
	max := 0
	allCombos := permutations(amplifiers)
	for i:= 0; i < len(allCombos); i++ {
		output := amplifySignal(allCombos[i], feedback)
		if output > max {
			max = output
		}
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