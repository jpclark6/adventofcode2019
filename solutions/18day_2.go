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
    allKeys = append(allKeys, "@")
	keyPairs := makeKeyPairs(allKeys)
	keyDistances := map[string]int{}
	keyRequiredDoors := map[string][]string{}
	keyFreeKeys := map[string][]string{}
	for i := 0; i < len(keyPairs); i++ {
		keys := keyPairs[i]
        distance, doors, freeKeys := findDistance(puzzleMap, keys)
		keyDistances[keys] = distance
        keyRequiredDoors[keys] = doors
        keyFreeKeys[keys] = freeKeys
		// fmt.Println(keys, distance, doors)
    }
    allKeys = allKeys[0:len(allKeys)-1]
    // fmt.Println("keyDistances", keyDistances)
    // fmt.Println("keyRequiredDoors", keyRequiredDoors)
    fmt.Println("keyFreeKeys", keyFreeKeys)
    // fmt.Println("keyPairs", keyPairs)

    // solveKeyGraph(allKeys, keyPairs, keyDistances, keyRequiredDoors, keyFreeKeys)
}

func copyKeys(keys []string) []string {
    remainingKeys := []string{}
    for j := 0; j < len(keys); j++ {
        remainingKeys = append(remainingKeys, keys[j])
    }
    return remainingKeys
}

func solveKeyGraph(keys []string, keyPairs []string, distances map[string]int, doors map[string][]string, freeKeys map[string][]string) {
    initialQ := []string{}
    for i := 0; i < len(keyPairs); i++ {
        key := keyPairs[i]
        if strings.Index(key, "@") >= 0 && len(doors[key]) == 0 {
            initialQ = append(initialQ, key)
            // fmt.Println("Key", key, "Doors", doors[key])
        }
    }
    for i := 0; i < len(initialQ); i++ {
        remainingKeys := copyKeys(keys)

        fmt.Println(remainingKeys)
    }
}

func findDistance(puzzleMap puzzle, keys string) (distance int, doors []string, freeKeys []string) {
	startCoords := findKey(puzzleMap, []rune(keys)[0])
	endCoords := findKey(puzzleMap, []rune(keys)[1])
	loc := startCoords
    puzzleMap.visited = map[string]string{startCoords: startCoords}
    puzzleMap.hereFrom = map[string]string{}
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
                    if isKey(space) && !(currentLoc == startCoords || currentLoc == endCoords) {
                        freeKeys = append(freeKeys, space)
                    }
					if currentLoc == startCoords {
						return distance, doors, freeKeys
					}
				}
			}
			puzzleMap = checkAdjacent(puzzleMap, puzzleMap.queue[j])
			puzzleMap.visited[puzzleMap.queue[j]] = "visited"
		}

		puzzleMap.queue = copyQueue(puzzleMap.newQueue)
	}

	return distance, doors, freeKeys
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
                kP := []rune(keys[i] + keys[j])
                if kP[0] > kP[1] {
                    kP[0], kP[1] = kP[1], kP[0]
                }
                keyPair := string(kP)
				keyPairs = append(keyPairs, keyPair)
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
