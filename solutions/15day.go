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
	start := time.Now()
	breatheOxygen()
	end := time.Now()
	fmt.Println("Total time:", end.Sub(start))
}

func breatheOxygen() {
	visited := map[string]bool{"0,0": true}
	queue := []string{"1", "2", "3", "4"}
	output := []int{0}
	input := []int{}
	minutes := 0

	for {
		minutes++
		currentQueue := queue
		for i := 0; i < len(currentQueue); i++ {
			currentLocation := parseLocationKey(currentQueue[i])
			visited[currentLocation] = true
			path := currentQueue[i]
			queue = queue[1:len(queue)]
			input = parsePathToInput(path)
			program := getProgram("./puzzledata/15day.txt")
			output, _, _, _, _ = runIntCode(input, program, 0, 0)
			latestMove := output[len(output)-1]
			if latestMove == 2 {
				fmt.Println("Answer to part 1:", minutes, "steps")
				minutes = -1
				visited = make(map[string]bool)
				queue = []string{path + "1", path + "2", path + "3", path + "4"}
				visited[parseLocationKey(path)] = true
				break
			} else if latestMove == 1 {
				queue = addNewMoves(queue, path, visited)
			}
		}
		if len(queue) == 0 {
			fmt.Println("Answer to part 2:", minutes, "minutes")
			return
		}
	}
}

func parseLocationKey(path string) string {
	p := strings.Split(path, "")
	x, y := 0, 0
	dir := map[string][]int{"1": []int{0, 1}, "2": []int{0, -1}, "3": []int{-1, 0}, "4": []int{1, 0}}
	for i := 0; i < len(p); i++ {
		x += dir[p[i]][0]
		y += dir[p[i]][1]
	}
	xS, yS := strconv.Itoa(x), strconv.Itoa(y)
	return xS + "," + yS
}

func addNewMoves(queue []string, path string, visited map[string]bool) []string {
	directions := []string{"1", "2", "3", "4"}
	queuePaths := map[string]bool{}
	for i := 0; i < len(queue); i++ {
		queuePaths[parseLocationKey(queue[i])] = true
	}
	for i := 0; i < len(directions); i++ {
		newPath := path + directions[i]
		loc := parseLocationKey(newPath)
		if !visited[loc] && !queuePaths[loc] {
			queue = append(queue, newPath)
			queuePaths[loc] = true
		}
	}
	return queue
}

func parsePathToInput(path string) []int {
	list := strings.Split(path, "")
	queue := []int{}
	for i := 0; i < len(list); i++ {
		direction, _ := strconv.Atoi(list[i])
		queue = append(queue, direction)
	}
	return queue
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
