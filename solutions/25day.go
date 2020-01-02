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
	// s := time.Now()
	part1()
	// e := time.Now()
	// fmt.Println("Finished in", e.Sub(s))
}

func part1() {
	text := "north\neast\ntake coin\neast\nnorth\nsouth\nwest\nwest\nsouth\neast\ntake cake\nsouth\ntake weather machine\nwest\ntake ornament\nwest\ntake jam\neast\neast\nnorth\nwest\nsouth\ntake food ration\nwest\ntake sand\nnorth\nnorth\neast\ntake astrolabe\nwest\nsouth\nsouth\neast\nnorth\neast\neast\neast\neast\ninv\ndrop sand\ndrop coin\ndrop jam\ndrop cake\nsouth\ninv\n"

	input := textToIntcode(text)
	fmt.Println(input)
	state := programState{
		input: input,
		program: getProgram("./puzzledata/25day.txt"),
		cursorLocation: 0,
		rBase: 0,
	}
	state = runIntCode(state)
}

func textToIntcode(text string) []int {
	rText := []rune(text)
	iText := []int{}
	for i := 0; i < len(rText); i++ {
		// fmt.Println(rText[i], int(rText[i]))
		iText = append(iText, int(rText[i]))
	}
	return iText
}

func printOutput(state programState) {
	output := state.output
	for i := 0; i < len(output); i++ {
		fmt.Printf("%c", output[i])
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
			fmt.Printf("%c", valueOne)
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
