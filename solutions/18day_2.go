package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
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
	newQueue  []string
	hereFrom  map[string]string
}

func part1(filename string) {
	file, _ := ioutil.ReadFile(filename)
	data := string(file)
	lines := strings.Split(data, "\n")
	puzzleMap := puzzle{grid: map[string]string{},
		keys:     map[string]bool{},
		hereFrom: map[string]string{},
		visited:  map[string]string{},
	}
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
	keyRequiredDoors := map[string][]string{}
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
}

func findDistance(puzzleMap puzzle, keys string) (distance int, doors []string) {
	startCoords := findKey(puzzleMap, []rune(keys)[0])
	endCoords := findKey(puzzleMap, []rune(keys)[1])
	loc := startCoords
	puzzleMap.visited[startCoords] = startCoords
	puzzleMap = checkAdjacent(puzzleMap, loc)
	puzzleMap.queue = copyQueue(puzzleMap.newQueue)

	for len(puzzleMap.queue) > 0 {
		puzzleMap.iteration++
		puzzleMap.newQueue = []string{}

		for j := 0; j < len(puzzleMap.queue); j++ {
			if puzzleMap.queue[j] == endCoords {
				distance = puzzleMap.iteration
				takenPath := []string{endCoords}
				currentLoc := endCoords
				for {
					currentLoc = puzzleMap.hereFrom[currentLoc]
					takenPath = append(takenPath, currentLoc)
					space := puzzleMap.grid[currentLoc]
					if isWall(space) {
						doors = append(doors, space)
					}
					if currentLoc == startCoords {
						return distance, doors
					}
				}
			}
			puzzleMap = checkAdjacent(puzzleMap, puzzleMap.queue[j])
			puzzleMap.visited[puzzleMap.queue[j]] = "visited"
		}

		puzzleMap.queue = copyQueue(puzzleMap.newQueue)
	}

	return distance, doors
}

func checkAdjacent(puzzleMap puzzle, sLoc string) puzzle {
	loc := parseKey(sLoc)
	dir := directions()

	for i := 0; i < len(dir); i++ {
		nextKey := makeKey(loc[0]+dir[i][0], loc[1]+dir[i][1])
		nextLoc := puzzleMap.grid[nextKey]
		if nextLoc != "" && nextLoc != "#" {
			if _, ok := puzzleMap.visited[nextKey]; !ok {
				puzzleMap.newQueue = append(puzzleMap.newQueue, nextKey)
				puzzleMap.hereFrom[nextKey] = sLoc
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
	keyPairs := []string{}
	for i := 0; i < len(keys)-1; i++ {
		for j := 0; j < len(keys); j++ {
			if j > i {
				keyPairs = append(keyPairs, keys[i]+keys[j])
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

func directions() [][]int {
	dir := [][]int{}
	dir = append(dir, []int{0, 1})
	dir = append(dir, []int{0, -1})
	dir = append(dir, []int{1, 0})
	dir = append(dir, []int{-1, 0})
	return dir
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
