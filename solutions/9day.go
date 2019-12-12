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
	fmt.Println("Answer to part 1:")
	runIntCode([]int{1}, getProgram(), 0)

	midTime := time.Now()

	fmt.Println("\nAnswer to part 2:")
	runIntCode([]int{2}, getProgram(), 0)

	endTime := time.Now()
	fmt.Println("Elapsed time for part 1:", midTime.Sub(startTime))
	fmt.Println("Elapsed time for part 2:", endTime.Sub(midTime))
	fmt.Println("Elapsed time for both parts:", endTime.Sub(startTime))
}

func getProgram() []int {
	content, err := ioutil.ReadFile("./puzzledata/9day.txt")
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
	return programInt
}

func hashProgram(program []int) map[int]int {
	mem := make(map[int]int)
	for i := 0; i < len(program); i++ {
		mem[i] = program[i]
	}
	return mem
}

func runIntCode(input []int, program []int, cursorLocation int) (output int, newProgramState map[int]int, finished bool, newLocation int) {
	var saveLocation int
	rBase := 0
	mem := hashProgram(program)

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
		instruction := mem[cursorLocation]
		instructionOpcode := instruction % 100
		modeOne := (instruction / 100) % 10
		modeTwo := (instruction / 1000) % 10
		modeThree := (instruction / 10000) % 10

		var valueOne, valueTwo int
		if threeParamGroup[instructionOpcode] {
			if modeOne == 0 {
				valueOne = mem[mem[cursorLocation+1]]
			} else if modeOne == 1 {
				valueOne = mem[cursorLocation+1]
			} else {
				valueOne = mem[mem[cursorLocation+1]+rBase]
			}

			if modeTwo == 0 {
				valueTwo = mem[mem[cursorLocation+2]]
			} else if modeTwo == 1 {
				valueTwo = mem[cursorLocation+2]
			} else {
				valueTwo = mem[mem[cursorLocation+2]+rBase]
			}

			if modeThree == 0 {
				saveLocation = mem[cursorLocation+3]
			} else {
				saveLocation = mem[cursorLocation+3] + rBase
			}
		}

		if twoParamGroup[instructionOpcode] {
			if modeOne == 0 {
				valueOne = mem[mem[cursorLocation+1]]
				saveLocation = mem[cursorLocation+1]
			} else if modeOne == 1 {
				valueOne = mem[cursorLocation+1]
				saveLocation = cursorLocation + 1
			} else {
				valueOne = mem[mem[cursorLocation+1]+rBase]
				saveLocation = mem[cursorLocation+1] + rBase
			}
		}

		switch instructionOpcode {
		case 1:
			mem[saveLocation] = valueOne + valueTwo
			cursorLocation += 4
		case 2:
			mem[saveLocation] = valueOne * valueTwo
			cursorLocation += 4
		case 3:
			var nextInput int
			if len(input) > 1 {
				nextInput, input = input[0], input[1:]
			} else if len(input) == 1 {
				nextInput, input = input[0], nil
			} else {
				return output, mem, false, cursorLocation
			}
			mem[saveLocation] = nextInput
			cursorLocation += 2
		case 4:
			output = valueOne
			fmt.Println("Output:", output)
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
				mem[saveLocation] = 1
			} else {
				mem[saveLocation] = 0
			}
			cursorLocation += 4
		case 8:
			if valueOne == valueTwo {
				mem[saveLocation] = 1
			} else {
				mem[saveLocation] = 0
			}
			cursorLocation += 4
		case 9:
			rBase += valueOne
			cursorLocation += 2
		case 99:
			return output, mem, true, cursorLocation
		default:
			println("Something went wrong... Found case:", instructionOpcode, cursorLocation, mem)
			break Loop
		}
	}
	return 0, mem, false, 0
}
