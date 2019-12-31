package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

var foundPart1 = false

func main() {
	s := time.Now()
	bothParts()
	e := time.Now()
	fmt.Println("Finished in", e.Sub(s))
}

func makeInitialStates([50]programState) [50]programState {
	file := "./puzzledata/23day.txt"
	var states [50]programState
	for i := 0; i < 50; i++ {
		input := []int{i}
		// fmt.Println("Initial input", input)
		states[i] = programState{
			input:          input,
			program:        getProgram(file),
			cursorLocation: 0,
			rBase:          0,
		}
		states[i] = runIntCode(states[i])
	}
	return states
}

func findInputs(states [50]programState) [51][]int {
	var nextInput [51][]int
	for i := 0; i < 50; i++ {
		state := states[i%50]
		previousOutput := state.output
		if len(previousOutput) > 1 {
			for k := 0; k < len(previousOutput)/3; k++ {
				computer := previousOutput[3*k]
				x := previousOutput[3*k+1]
				y := previousOutput[3*k+2]
				if computer == 255 {
					if !foundPart1 {
						fmt.Println("Found part 1:", y)
						foundPart1 = true
					}
					nextInput[50] = []int{x, y}
				} else {
					nextInput[computer] = append(nextInput[computer], []int{x, y}...)
				}
			}
		}
	}
	return nextInput
}

func bothParts() {
	var states [50]programState
	prevY := 0
	states = makeInitialStates(states)
	for j := 0; j < 500; j++ {
		nextInput := findInputs(states)
		currentlyIdle := 0
		for i := 0; i < 50; i++ {
			if len(nextInput[i]) == 0 {
				nextInput[i] = []int{-1}
				currentlyIdle++
			}
		}
		if currentlyIdle == 50 && j != 0 {
			nextInput[0] = nextInput[50]
			if nextInput[0][1] == prevY {
				fmt.Println("Found part 2:", prevY)
				return
			}
			prevY = nextInput[0][1]
		}
		for i := 0; i < 50; i++ {
			state := states[i%50]
			state.input = nextInput[i]
			state = runIntCode(state)
			states[i%50] = state
		}
	}
}

type programState struct {
	input          []int
	output         []int
	program        map[int]int
	cursorLocation int
	rBase          int
	finished       bool
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

func runIntCode(state programState) programState {
	input := state.input
	program := state.program
	cursorLocation := state.cursorLocation
	rBase := state.rBase
	output := []int{}

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
				return programState{output: output,
					program:        program,
					finished:       false,
					cursorLocation: cursorLocation,
					rBase:          rBase,
				}
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
			fmt.Println("Done")
			return programState{output: output,
				program:        program,
				finished:       true,
				cursorLocation: cursorLocation,
				rBase:          rBase,
			}
		default:
			println("Something went wrong... Found case:", instructionOpcode, cursorLocation, program)
			break Loop
		}
	}
	return programState{output: []int{0},
		program:        program,
		finished:       false,
		cursorLocation: 0,
		rBase:          0,
	}
}
