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
	// start := time.Now()
	drawMaze()
	// end := time.Now()
	// fmt.Println("Total time:", end.Sub(start))
}

func drawMaze() {
	program := getProgram("./puzzledata/17day.txt")
	output, _, _, _, _ := runIntCode([]int{}, program, 0, 0)
	grid := make(map[string]int)
	fromTop := 0
	fromSide := 0
	foundLength := false
	length := 0
	for i := 0; i < len(output); i++ {
		switch output[i] {
		case 10:
			fromTop++
			if !foundLength {
				length = fromSide
				foundLength = true
			}
			fromSide = 0
			fmt.Println()
		default:
			xy := coordToKey(fromSide, fromTop)
			grid[xy] = output[i]
			r := rune(output[i])
			fmt.Printf("%c", r)
			fromSide++
		}
	}
	fromTop--
	totalAlignment := findAlignmentProduct(length, fromTop, grid)
	fmt.Println("Total for part 1:", totalAlignment, "\n ")
	path := findPath(grid, length, fromTop)
	intCodeInput := buildInput(path)
	program = getProgram("./puzzledata/17day.txt")
	program[0] = 2
	output, _, _, _, _ = runIntCode(intCodeInput, program, 0, 0)
	fmt.Println("Part 2 answer:", output[len(output) - 1])
}

func buildInput(path string) []int {
	remainingPath := path
	a, remainingPath := findSegment(remainingPath)
	b, remainingPath := findSegment(remainingPath)
	c, _ := findSegment(remainingPath)

	path = strings.Replace(path, a, "A", -1)
	path = strings.Replace(path, b, "B", -1)
	path = strings.Replace(path, c, "C", -1)

	path += "\n" + a + "\n" + b + "\n" + c + "\n" + "n" + "\n"

	fmt.Println("Path", path)

	rInput := []rune(path)
	inputIntCode := []int{}
	for i := 0; i < len(rInput); i++ {
		inputIntCode = append(inputIntCode, int(rInput[i]))
	}

	return inputIntCode
}

func findSegment(path string) (segment string, newPath string) {
	minLength := len(path)
	for i := 6; i <= 21; i++ {
		guess := string([]rune(path)[0:i])
		pathCopy := path
		pathCopy = strings.Replace(pathCopy, guess, "", -1)
		if len(pathCopy) < minLength {
			segment = guess
			minLength = len(pathCopy)
		} else {
			guess := string([]rune(path)[0:i - 1])
			path = strings.Replace(path, guess, "", -1)
			path = strings.Replace(path, ",,", ",", -1)
			path = strings.Trim(path, ",")
			guess = strings.Trim(guess, ",")
			return guess, strings.Replace(path, guess, "", -1)
		}
	}
	fmt.Println("Nope")
	return "", ""
}

func findPath(grid map[string]int, length int, height int) (path string) {
	robotLocation := []int{0, 0}
	for key, square := range(grid) {
		if square == 94 {
			x, y := keyToCoord(key)
			robotLocation[0] = x
			robotLocation[1] = y
			break
		}
	}
	var atDeadEnd bool
	atTurn := true
	direction := "n"
	turnDirection := ""
	forwardSpaces := 0
	pathInstructions := ""
	for {
		for atTurn != true {
			robotLocation, atTurn = goStraight(robotLocation, direction, grid, length, height)
			if !atTurn {
				forwardSpaces++
			}
		}
		pathInstructions += strconv.Itoa(forwardSpaces) + ","
		atTurn = false
		forwardSpaces = 0
		direction, turnDirection, atDeadEnd = goTurn(robotLocation, direction, grid, length, height)
		if atDeadEnd {
			pathInstructions = string([]rune(pathInstructions)[2:len(pathInstructions) - 1])
			return pathInstructions
		}
		pathInstructions += turnDirection + ","
	}
}

func goTurn(robotLocation []int, direction string, grid map[string]int, length int, height int) (string, string, bool) {
	sideCoords := findSideCoords(robotLocation, direction)
	if grid[coordToKey(sideCoords[0][0], sideCoords[0][1])] == 35 {
		if direction == "n" {
			return "w", "L", false
		} else if direction == "s" {
			return "w", "R", false
		} else if direction == "e" {
			return "n", "L", false
		}
		return "n", "R", false
	} else if grid[coordToKey(sideCoords[1][0], sideCoords[1][1])] == 35 {
		if direction == "n" {
			return "e", "R", false
		} else if direction == "s" {
			return "e", "L", false
		} else if direction == "e" {
			return "s", "R", false
		}
		return "s", "L", false
	}
	return direction, "\n", true
}

func findSideCoords(loc []int, direction string) [][]int {
	if direction == "n" || direction == "s" {
		d1 := []int{loc[0] - 1, loc[1]}
		d2 := []int{loc[0] + 1, loc[1]}
		return [][]int{d1, d2}
	}
	d1 := []int{loc[0], loc[1] - 1}
	d2 := []int{loc[0], loc[1] + 1}
	return [][]int{d1, d2}
}

func goStraight(loc []int, direction string, grid map[string]int, l int, h int) ([]int, bool) {
	// for {
		step := moveForwardOne(loc, direction)
		nextLoc := []int{loc[0] + step[0], loc[1] + step[1]}
		// if nextLoc[0] >= 0 && nextLoc[0] < l && nextLoc[1] >= 0 && nextLoc[1] < h {
		if grid[coordToKey(nextLoc[0], nextLoc[1])] == 35 {
			return nextLoc, false
		}
	// }
	return loc, true
}

func moveForwardOne(loc []int, direction string) []int {
	step := []int{0, 0}
	if direction == "n" {
		step[1]--
	} else if direction == "e" {
		step[0]++
	} else if direction == "s" {
		step[1]++
	} else if direction == "w" {
		step[0]--
	}
	return step
}

func findAlignmentProduct(length int, fromTop int, grid map[string]int) int {
	totalAlignment := 0
	for x := 1; x < length - 1; x++ {
		for y := 1; y < fromTop - 1; y++ {
			if atIntersection(x, y, grid) {
				totalAlignment += x * y
			}
		}
	}
	return totalAlignment
}

func atIntersection(x int, y int, grid map[string]int) bool {
	if grid[coordToKey(x, y)] == 35 && 
		grid[coordToKey(x + 1, y)] == 35 &&
		grid[coordToKey(x - 1, y)] == 35 &&
		grid[coordToKey(x, y + 1)] == 35 &&
		grid[coordToKey(x, y - 1)] == 35 {
			return true
		}
	return false
}

func coordToKey(fromSide int, fromTop int) string {
	x := strconv.Itoa(fromSide)
	y := strconv.Itoa(fromTop)
	return x + "," + y
}

func keyToCoord(key string) (x int, y int) {
	xy := strings.Split(key, ",")
	x, _ = strconv.Atoi(xy[0])
	y, _ = strconv.Atoi(xy[1])
	return x, y
}


// func drawMaze() {
// 	visited := map[string]bool{"0,0": true}
// 	queue := []string{"1", "2", "3", "4"}
// 	output := []int{0}
// 	input := []int{}
// 	minutes := 0

// 	for {
// 		minutes++
// 		currentQueue := queue
// 		for i := 0; i < len(currentQueue); i++ {
// 			currentLocation := parseLocationKey(currentQueue[i])
// 			visited[currentLocation] = true
// 			path := currentQueue[i]
// 			queue = queue[1:len(queue)]
// 			input = parsePathToInput(path)
// 			program := getProgram("./puzzledata/15day.txt")
// 			output, _, _, _, _ = runIntCode(input, program, 0, 0)
// 			latestMove := output[len(output)-1]
// 			if latestMove == 2 {
// 				fmt.Println("Answer to part 1:", minutes, "steps")
// 				minutes = -1
// 				visited = make(map[string]bool)
// 				queue = []string{path + "1", path + "2", path + "3", path + "4"}
// 				visited[parseLocationKey(path)] = true
// 				break
// 			} else if latestMove == 1 {
// 				queue = addNewMoves(queue, path, visited)
// 			}
// 		}
// 		if len(queue) == 0 {
// 			fmt.Println("Answer to part 2:", minutes, "minutes")
// 			return
// 		}
// 	}
// }

// func parseLocationKey(path string) string {
// 	p := strings.Split(path, "")
// 	x, y := 0, 0
// 	dir := map[string][]int{"1": []int{0, 1}, "2": []int{0, -1}, "3": []int{-1, 0}, "4": []int{1, 0}}
// 	for i := 0; i < len(p); i++ {
// 		x += dir[p[i]][0]
// 		y += dir[p[i]][1]
// 	}
// 	xS, yS := strconv.Itoa(x), strconv.Itoa(y)
// 	return xS + "," + yS
// }

// func addNewMoves(queue []string, path string, visited map[string]bool) []string {
// 	directions := []string{"1", "2", "3", "4"}
// 	queuePaths := map[string]bool{}
// 	for i := 0; i < len(queue); i++ {
// 		queuePaths[parseLocationKey(queue[i])] = true
// 	}
// 	for i := 0; i < len(directions); i++ {
// 		newPath := path + directions[i]
// 		loc := parseLocationKey(newPath)
// 		if !visited[loc] && !queuePaths[loc] {
// 			queue = append(queue, newPath)
// 			queuePaths[loc] = true
// 		}
// 	}
// 	return queue
// }

// func parsePathToInput(path string) []int {
// 	list := strings.Split(path, "")
// 	queue := []int{}
// 	for i := 0; i < len(list); i++ {
// 		direction, _ := strconv.Atoi(list[i])
// 		queue = append(queue, direction)
// 	}
// 	return queue
// }

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
