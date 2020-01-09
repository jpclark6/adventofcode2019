package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
	"sort"
	"math/rand"
	"time"
)

func main() {
	part1("./puzzledata/18day_example.txt")
}

type puzzle struct {
	grid      map[string]string
	keys      map[string]bool
	visited   map[string]string
	queue     []string
	iteration int
	newQueue	[]string
}

func copyPuzzle(puzzleMap puzzle) puzzle {
	grid := puzzleMap.grid
	keys := puzzleMap.keys
	visited := puzzleMap.visited
	queue := puzzleMap.queue
	newQueue := puzzleMap.newQueue

	newGrid := map[string]string{}
	newKeys := map[string]bool{}
	newVisited := map[string]string{}
	nQueue := []string{}
	newNewQueue := []string{}

	for k, v := range grid {
		newGrid[k] = v
	}
	for k, v := range keys {
		newKeys[k] = v
	}
	for k, v := range visited {
		newVisited[k] = v
	}
	for i := 0; i < len(queue); i++ {
		nQueue = append(nQueue, queue[i])
	}
	for i := 0; i < len(newQueue); i++ {
		newNewQueue = append(newNewQueue, newQueue[i])
	}

	newPuzzle := puzzle{grid: newGrid, keys: newKeys, visited: newVisited, queue: nQueue, iteration: puzzleMap.iteration, newQueue: newNewQueue}

	return newPuzzle
}

func part1(filename string) {
	file, _ := ioutil.ReadFile(filename)
	data := string(file)
	lines := strings.Split(data, "\n")
	puzzleMap := puzzle{grid: map[string]string{}, keys: map[string]bool{}}
	allKeys := []string{}

	for row := 0; row < len(lines); row++ {
		line := strings.Split(lines[row], "")
		for column := 0; column < len(line); column++ {
			key := makeKey(column, row)
			char := line[column]
			puzzleMap.grid[key] = char
			rchar := []rune(char)[0]
			if unicode.IsLower(rchar) {
				puzzleMap.keys[char] = false
				allKeys = append(allKeys, string(rchar))
			}
		}
	}
	keyPairs := makeKeyPairs(allKeys)
	keyDistances := map[string]int{}
	keyRequiredDoors := map[string]string{}
	for i := 0; i < len(keyPairs); i++ {
		kP := []rune(keyPairs[i])
		if kP[0] > kP[1] {
			kP[0], kP[1] = kP[1], kP[0]
		}
		keys := string(kP)
		distance, doors := findDistance(puzzleMap, keys)
		keyDistances[keys] = distance
		keyRequiredDoors[keys] = doors
		fmt.Println(keys, distance, doors)
	}
	// solve(puzzleMap)
}

func findDistance(puzzleMap puzzle, keys string) (distance int, doors string) {
	startCoords := findKey(puzzleMap, []rune(keys)[0])
	endCoords := findKey(puzzleMap, []rune(keys)[1])
	loc := startCoords 	
	puzzleMap.visited = map[string]string{startCoords: startCoords}
	puzzleMap = checkAdjacent(puzzleMap, loc)
	puzzleMap.queue = copyQueue(puzzleMap.newQueue)

	for len(puzzleMap.queue) > 0 {
		puzzleMap.iteration++
		puzzleMap.newQueue = []string{}
		
		for j := 0; j < len(puzzleMap.queue); j++ {
			if puzzleMap.queue[j] == endCoords {
				distance = puzzleMap.iteration
				doors = "tbd"
				return distance, doors
			}
			puzzleMap = checkAdjacent(puzzleMap, puzzleMap.queue[j])
			puzzleMap.visited[puzzleMap.queue[j]] = puzzleMap.iteration
		}

		puzzleMap.queue = copyQueue(puzzleMap.newQueue)
	}
	
	return distance, doors
}

func checkAdjacent(puzzleMap puzzle, sLoc string) puzzle {
	loc := parseKey(sLoc)
	dir := directions()

	for i := 0; i < len(dir); i++ {
		nextKey := makeKey(loc[0] + dir[i][0], loc[1] + dir[i][1])
		nextLoc := puzzleMap.grid[nextKey]
		if nextLoc != "" && nextLoc != "#" {
			if _, ok := puzzleMap.visited[nextKey]; !ok {
				puzzleMap.newQueue = append(puzzleMap.newQueue, nextKey)
			}
		}
	}
	return puzzleMap
}


func findKey(puzzleMap puzzle, rKey rune) string {
	key := string(rKey)
	for loc, char := range puzzleMap.grid {
		if char == key {
			return loc
		}
	}
	return "not found"
}

func makeKeyPairs(keys []string) []string {
	fmt.Println(keys)
	keyPairs := []string{}
	for i := 0; i < len(keys) - 1; i++ {
		for j := 0; j < len(keys); j++ {
			if j > i {
				keyPairs = append(keyPairs, keys[i] + keys[j])
			}
		}
	}

	return keyPairs
}

func isKey(loc string) bool {
	rchar := []rune(loc)[0]
	if unicode.IsLower(rchar) {
		return true
	}
	return false
}

func isWall(loc string) bool {
	rchar := []rune(loc)[0]
	if unicode.IsUpper(rchar) {
		return true
	}
	return false
}

func findStart(puzzleMap puzzle) string {
	for loc, char := range puzzleMap.grid {
		if char == "@" {
			return loc
		}
	}
	return "not found"
}

func copyQueue(queue []string) []string {
	newQueue := []string{}
	for i := 0; i < len(queue); i++ {
		newQueue = append(newQueue, queue[i])
	}
	return newQueue
}

func solve(puzzleMap puzzle) {
	startCoords := findStart(puzzleMap)
	loc := startCoords
	puzzleMap.visited = map[string]int{startCoords: 0}
	puzzleMap, _ = checkSurroundingSpaces(puzzleMap, loc)
	puzzleMap.queue = copyQueue(puzzleMap.newQueue)
	puzzleMap.grid[loc] = "."
	puzzleMaps := []puzzle{puzzleMap}

	trimLimit := 5000
	trimIteration := 3

	for i := 0; i < 6000; i++ {
		extraMaps := []puzzle{}
		for m := 0; m < len(puzzleMaps); m++ {
			// fmt.Printf("Starting map: %+v\n", puzzleMaps[m])
			if len(puzzleMaps[m].queue) > 0 {
				puzzleMap = puzzleMaps[m]
				puzzleMap.iteration++
				puzzleMap.newQueue = []string{}
				
				for j := 0; j < len(puzzleMap.queue); j++ {
					extraMap := []puzzle{}
					puzzleMap, extraMap = checkSurroundingSpaces(puzzleMap, puzzleMap.queue[j])
					extraMaps = append(extraMaps, extraMap...)
					puzzleMap.visited[puzzleMap.queue[j]] = puzzleMap.iteration
				}

				puzzleMap.queue = copyQueue(puzzleMap.newQueue)
				puzzleMaps[m] = puzzleMap
			}
		}
		puzzleMaps = append(puzzleMaps, extraMaps...)
		nextMaps := []puzzle{}
		for m := 0; m < len(puzzleMaps); m++ {
			if len(puzzleMaps[m].queue) != 0 {
				nextMaps = append(nextMaps, puzzleMaps[m])
				// fmt.Printf("Remaining map: %+v\n", puzzleMaps[m])
			}
		}
		puzzleMaps = nextMaps
		// masterKeyList := map[string]int{}
		if i % trimIteration == 0 {
			// sort.Slice(puzzleMaps, func(i, j int) bool {
			// 	a := 0
			// 	b := 0
			// 	for _, v := range puzzleMaps[i].keys {
			// 		if v == true {
			// 			a++
			// 		}
			// 	}
			// 	for _, v := range puzzleMaps[j].keys {
			// 		if v == true {
			// 			b++
			// 		}
			// 	}
			// 	return a > b
			// })
			nextPuzzles := []puzzle{}
			rand.Seed(time.Now().UnixNano())
			if len(puzzleMaps) > trimLimit {
				for j := 0; j < trimLimit; j++ {
					nextPuzzles = append(nextPuzzles, puzzleMaps[rand.Intn(len(puzzleMaps))])
				}
				// puzzleMaps = puzzleMaps[0:trimLimit]
				puzzleMaps = nextPuzzles
			}
			
			
			// for p := 0; p < len(puzzleMaps); p++ {
			// 	if _, ok := masterKeyList[keyList(puzzleMaps[p])]; !ok {
			// 		masterKeyList[keyList(puzzleMaps[p])] = 1
			// 		nextPuzzles = append(nextPuzzles, puzzleMaps[p])
			// 	} else {
			// 		if masterKeyList[keyList(puzzleMaps[p])] < 2 {
			// 			nextPuzzles = append(nextPuzzles, puzzleMaps[p])
			// 		}
			// 	}
			// }
			// puzzleMaps = nextPuzzles
		}
		if i % 50 == 0 {
			fmt.Println("Iteration", i, "Open maps:", len(puzzleMaps))
		}
	}
}

func keyList(puzzleMap puzzle) string {
	allKeys := []string{}
	for k, v := range puzzleMap.keys {
		if v == true {
			allKeys = append(allKeys, k)
		}
	}
	sort.Strings(allKeys)
	keyList := strings.Join(allKeys, "")
	return keyList
}

func directions() [][]int {
	dir := [][]int{}
	dir = append(dir, []int{0, 1})
	dir = append(dir, []int{0, -1})
	dir = append(dir, []int{1, 0})
	dir = append(dir, []int{-1, 0})
	return dir
}

func checkSurroundingSpaces(puzzleMap puzzle, sLoc string) (puzzle, []puzzle) {
	loc := parseKey(sLoc)
	dir := directions()
	additionalPuzzles := []puzzle{}

	for i := 0; i < len(dir); i++ {
		nextKey := makeKey(loc[0] + dir[i][0], loc[1] + dir[i][1])
		nextLoc := puzzleMap.grid[nextKey]
		if nextLoc != "" {
			if _, ok := puzzleMap.visited[nextKey]; !ok {
				if nextLoc == "." {
					puzzleMap.newQueue = append(puzzleMap.newQueue, nextKey)
				} else if isKey(nextLoc) {
					nextPuzzle := copyPuzzle(puzzleMap)
					nextPuzzle.keys[nextLoc] = true
					nextPuzzle.grid[nextKey] = "."
					nextPuzzle.queue = []string{nextKey}
					nextPuzzle.newQueue = []string{}
					nextPuzzle.visited = map[string]int{}

					finished := true
					for _, v := range nextPuzzle.keys {
						if v == false {
							finished = false
						}
					}
					if finished == true {
						fmt.Println("Solution found:", nextPuzzle.iteration + 1)
						panic("Found solution")
					}
					additionalPuzzles = append(additionalPuzzles, nextPuzzle)
				} else if isWall(nextLoc) {
					if puzzleMap.keys[strings.ToLower(nextLoc)] == true {
						puzzleMap.grid[nextLoc] = "."
						puzzleMap.newQueue = append(puzzleMap.newQueue, nextKey)
					}
				}
			}
		}
	}
	return puzzleMap, additionalPuzzles
}

func parseKey(xy string) []int {
	xS := strings.Split(xy, ",")[0]
	x, _ := strconv.Atoi(xS)
	yS := strings.Split(xy, ",")[1]
	y, _ := strconv.Atoi(yS)
	return []int{x, y}
}

func makeKey(x int, y int) (key string) {
	key = strconv.Itoa(x) + "," + strconv.Itoa(y)
	return key
}
