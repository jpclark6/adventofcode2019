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
	part1()
}

func part1() {
	program := getProgram("./puzzledata/21day.txt")
	input := parseInput()
	output, _, _, _, _ := runIntCode(input, program, 0, 0)
	// fmt.Println("SOLDUTIOND", output)
	parseOutput(output)
}

func parseOutput(output []int) {
	for i := 0; i < len(output); i++ {
		fmt.Printf(string(output[i]))
	}
}

func parseInput() []int {
// 	input := 
// `NOT C T
// AND D T
// NOT A J
// OR T J
// WALK
// `

	inputArr := []string{"NOT C T",
		"AND D T",
		"NOT E J",
		"AND T J",
		"NOT A T",
		"OR T J",
		"NOT E T",
		"AND A T",
		"AND B T",
		"AND F T",
		"OR T J",
		"AND D J",
		"NOT A T",
		"OR T J",
		"RUN",
		""}

/*
Jump if:
Next 3 spaces are holes
Out of 4 spaces, only 1 is land, jump to land there
Out of 4 spaces after 4 spaces, if only 1 is land jump there
*/
// `NOT C T
// AND D T
// AND G T
// NOT A J
// OR T J
// RUN
// `
	input := strings.Join(inputArr, "\n")
	intRunes := []int{}
	inputRunes := []rune(input)
	for i := 0; i < len(inputRunes); i++ {
		intRunes = append(intRunes, int(inputRunes[i]))
	}
	return intRunes
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
			// fmt.Println("Pausing for additional input")
			if len(input) > 1 {
				nextInput, input = input[0], input[1:]
			} else if len(input) == 1 {
				nextInput, input = input[0], nil
			} else {
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
			// fmt.Println(valueOne)
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
