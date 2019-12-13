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
	solution("part 1")
	midTime := time.Now()
	solution("part 2")
	endTime := time.Now()
	fmt.Println("Elapsed time for part 1:", midTime.Sub(startTime))
	fmt.Println("Elapsed time for part 2:", endTime.Sub(midTime))
	fmt.Println("Elapsed time for both parts:", endTime.Sub(startTime))
}

func solution(part string) {
	file := "./puzzledata/11day.txt"
	y, x := 180, 120
	grid := makeGrid(y, x)
	program := getProgram(file)
	cursorLocation, rBase := 0, 0
	currentX, currentY := x / 2, y / 2
	finished := false
	var output []int
	direction := "N"
	paintList := make(map[string]bool)

	if part == "part 2" {
		grid[currentY][currentX] = 1
	}

	for ; !finished;  {
		currentColor := grid[currentY][currentX]
		input := []int{currentColor}
		output, program, finished, cursorLocation, rBase = 
			runIntCode(input, program, cursorLocation, rBase)

		paintList = addToPaintlist(paintList, currentX, currentY)

		color := output[0]
		instruction := output[1]

		grid = paintGrid(grid, currentX, currentY, color)
		grid[currentY][currentX] = output[0]
		currentX, currentY, direction = turnAndMoveForward(currentX, currentY, direction, instruction)
	}

	if part == "part 1" {
		fmt.Println("Part 1 answer:", len(paintList))
	} else {
		paintGrid := makePaintGrid(y, x)
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == 0 {
					paintGrid[i][x - 1 - j] = " "
				} else {
					paintGrid[i][x - 1 - j] = "O"
				}
			}
		}
		var picture []string
		for i := 0; i < 10; i++ {
			picture = append(picture, strings.Join(paintGrid[i + 83][18:60], ""))
		}
		fmt.Println("View in mirror to see answer to part 2")
		for i := 0; i < len(picture); i++ {
			fmt.Println(picture[len(picture) - 1 - i])
		}
	}
}

func addToPaintlist(paintList map[string]bool, x int, y int) map[string]bool {
	location := fmt.Sprintf("%v,%v", x, y)
	paintList[location] = true
	return paintList
}

func turnAndMoveForward(x int, y int, direction string, instruction int) (newX int, newY int, newDirection string) {
	switch direction {
	case "N":
		if instruction == 0 {
			direction = "W"
			x--
		} else {
			direction = "E"
			x++
		}
	case "E":
		if instruction == 0 {
			direction = "N"
			y++
		} else {
			direction = "S"
			y--
		}
	case "S":
		if instruction == 0 {
			direction = "E"
			x++
		} else {
			direction = "W"
			x--
		}
	case "W":
		if instruction == 0 {
			direction = "S"
			y--
		} else {
			direction = "N"
			y++
		}
	}
	return x, y, direction
}

func paintGrid(grid [][]int, x int, y int, color int) [][]int {
	grid[y][x] = color
	return grid
}

func makeGrid(y int, x int) [][]int {
	height := y
	width := x
	var grid [][]int
	for i := 0; i < height; i++ {
		var tempArray []int
		for j := 0; j < width; j++ {
			tempArray = append(tempArray, 0)
		}
		grid = append(grid, tempArray)
	}
	return grid
}

func makePaintGrid(y int, x int) [][]string {
	height := y
	width := x
	var grid [][]string
	for i := 0; i < height; i++ {
		var tempArray []string
		for j := 0; j < width; j++ {
			tempArray = append(tempArray, " ")
		}
		grid = append(grid, tempArray)
	}
	return grid
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
