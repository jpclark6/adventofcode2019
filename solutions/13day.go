package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	// "time"
)

func main() {
	// startTime := time.Now()
	program := getProgram("./puzzledata/13day.txt")
	input := []int{0}
	cursorLocation := 0
	rBase := 0
	input, program, _, cursorLocation, rBase = runIntCode(input, program, cursorLocation, rBase)
	tileCount := 0
	for i := 2; i < len(program); i += 3 {
		if input[i] == 2 {
			tileCount++
		}
	} 
	fmt.Println("Answer to part 1:", tileCount)
	
	// midTime := time.Now()

	// endTime := time.Now()
	// fmt.Println("Elapsed time for part 1:", midTime.Sub(startTime))
	// fmt.Println("Elapsed time for part 2:", endTime.Sub(midTime))
	// fmt.Println("Elapsed time for both parts:", endTime.Sub(startTime))
}

func getProgram(file string) map[int]int {
	content, err := ioutil.ReadFile(file)
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
	return hashProgram(programInt)
}

func hashProgram(program []int) map[int]int {
	mem := make(map[int]int)
	for i := 0; i < len(program); i++ {
		mem[i] = program[i]
	}
	return mem
}

func runIntCode(input []int, program map[int]int, cursorLocation int, rBase int) (output []int, newProgramState map[int]int, finished bool, newLocation int, newRBase int) {
	var saveLocation, nextInput int

	var threeParamGroup = map[int]bool{
		1: true,
		2: true,
		5: true,
		6: true,
		7: true,
		8: true,
	}

	var twoParamGroup = map[int]bool{
		3: true,
		4: true,
		9: true,
	}

Loop:
	for {
		instruction := program[cursorLocation]
		instructionOpcode := instruction % 100
		modeOne := (instruction / 100) % 10
		modeTwo := (instruction / 1000) % 10
		modeThree := (instruction / 10000) % 10

		var valueOne, valueTwo int
		if threeParamGroup[instructionOpcode] {
			if modeOne == 0 {
				valueOne = program[program[cursorLocation+1]]
			} else if modeOne == 1 {
				valueOne = program[cursorLocation+1]
			} else {
				valueOne = program[program[cursorLocation+1]+rBase]
			}

			if modeTwo == 0 {
				valueTwo = program[program[cursorLocation+2]]
			} else if modeTwo == 1 {
				valueTwo = program[cursorLocation+2]
			} else {
				valueTwo = program[program[cursorLocation+2]+rBase]
			}

			if modeThree == 0 {
				saveLocation = program[cursorLocation+3]
			} else {
				saveLocation = program[cursorLocation+3] + rBase
			}
		}

		if twoParamGroup[instructionOpcode] {
			if modeOne == 0 {
				valueOne = program[program[cursorLocation+1]]
				saveLocation = program[cursorLocation+1]
			} else if modeOne == 1 {
				valueOne = program[cursorLocation+1]
				saveLocation = cursorLocation + 1
			} else {
				valueOne = program[program[cursorLocation+1]+rBase]
				saveLocation = program[cursorLocation+1] + rBase
			}
		}

		if instructionOpcode == 3 {
			if len(input) > 1 {
				nextInput, input = input[0], input[1:]
			} else if len(input) == 1 {
				nextInput, input = input[0], nil
			} else {
				// fmt.Println("Pausing for additional input")
				return output, program, false, cursorLocation, rBase
			}
		}

		switch instructionOpcode {
		case 1:
			program[saveLocation] = valueOne + valueTwo
			cursorLocation += 4
		case 2:
			program[saveLocation] = valueOne * valueTwo
			cursorLocation += 4
		case 3:
			program[saveLocation] = nextInput
			cursorLocation += 2
		case 4:
			output = append(output, valueOne)
			// fmt.Println("Output:", output)
			cursorLocation += 2
		case 5:
			if valueOne != 0 {
				cursorLocation = valueTwo
			} else {
				cursorLocation += 3
			}
		case 6:
			if valueOne == 0 {
				cursorLocation = valueTwo
			} else {
				cursorLocation += 3
			}
		case 7:
			if valueOne < valueTwo {
				program[saveLocation] = 1
			} else {
				program[saveLocation] = 0
			}
			cursorLocation += 4
		case 8:
			if valueOne == valueTwo {
				program[saveLocation] = 1
			} else {
				program[saveLocation] = 0
			}
			cursorLocation += 4
		case 9:
			rBase += valueOne
			cursorLocation += 2
		case 99:
			return output, program, true, cursorLocation, rBase
		default:
			println("Something went wrong... Found case:", instructionOpcode, cursorLocation, program)
			break Loop
		}
	}
	return []int{0}, program, false, 0, 0
}
