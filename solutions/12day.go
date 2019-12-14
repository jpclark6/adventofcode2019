package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	part1And2()
}

func part1And2() {
	moons := parseFile("./puzzledata/12day.txt")
	part1Steps := 1000
	moons = stepTime(part1Steps, moons)
}

func findTotalEnergy(moons []moon) int {
	tE := 0
	for i := 0; i < len(moons); i++ {
		body := &moons[i]
		pE := findPotentialEnergy(*body)
		kE := findKineticEnergy(*body)
		body.tEnergy = pE * kE
		tE += pE * kE
	}
	return tE
}

func findPotentialEnergy(body moon) int {
	x := math.Abs(float64(body.x))
	y := math.Abs(float64(body.y))
	z := math.Abs(float64(body.z))
	pE := int(x + y + z)
	return pE
}

func findKineticEnergy(body moon) int {
	xVel := math.Abs(float64(body.xVel))
	yVel := math.Abs(float64(body.yVel))
	zVel := math.Abs(float64(body.zVel))
	kE := int(xVel + yVel + zVel)
	return kE
}

func stepTime(steps int, moons []moon) []moon {
	startTime := time.Now()
	midTime := time.Now()
	endTime := time.Now()

	xInfo := make(map[string]bool)
	yInfo := make(map[string]bool)
	zInfo := make(map[string]bool)

	xIterations := 0
	yIterations := 0
	zIterations := 0

	var moonsAtSteps []moon

	for i := 0; ; i++ {
		masterX := ""
		masterY := ""
		masterZ := ""

		for j := 0; j < len(moons); j++ {
			x := strconv.Itoa(moons[j].x)
			xVel := strconv.Itoa(moons[j].xVel)
			xKey := x + "," + xVel + ","

			y := strconv.Itoa(moons[j].y)
			yVel := strconv.Itoa(moons[j].yVel)
			yKey := y + "," + yVel + ","

			z := strconv.Itoa(moons[j].z)
			zVel := strconv.Itoa(moons[j].zVel)
			zKey := z + "," + zVel + ","

			masterX += xKey
			masterY += yKey
			masterZ += zKey
		}

		if xInfo[masterX] && xIterations == 0 {
			xIterations = i
		} else {
			xInfo[masterX] = true
		}

		if yInfo[masterY] && yIterations == 0 {
			yIterations = i
		} else {
			yInfo[masterY] = true
		}

		if zInfo[masterZ] && zIterations == 0 {
			zIterations = i
		} else {
			zInfo[masterZ] = true
		}

		if xIterations != 0 && yIterations != 0 && zIterations != 0 {
			break
		}
		moons = applyGravity(moons)
		moons = applyVelocity(moons)
		if i == steps-1 {
			tE := findTotalEnergy(moons)
			fmt.Println("Answer to part 1:", tE)
			midTime = time.Now()
		}
	}
	fmt.Println("Answer to part 2. LCD of:", xIterations, yIterations, zIterations)
	endTime = time.Now()
	fmt.Println("Part 1 took", midTime.Sub(startTime))
	fmt.Println("Part 2 took", endTime.Sub(midTime))
	fmt.Println("Total time:", endTime.Sub(startTime))
	return moonsAtSteps
}

func applyGravity(moons []moon) []moon {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3-i; j++ {
			moons[i], moons[3-j] = applyPairGravity(moons[i], moons[3-j])
		}
	}
	return moons
}

func applyPairGravity(body1 moon, body2 moon) (newBody1 moon, newBody2 moon) {
	if body1.x > body2.x {
		body1.xVel--
		body2.xVel++
	} else if body1.x < body2.x {
		body1.xVel++
		body2.xVel--
	}
	if body1.y > body2.y {
		body1.yVel--
		body2.yVel++
	} else if body1.y < body2.y {
		body1.yVel++
		body2.yVel--
	}
	if body1.z > body2.z {
		body1.zVel--
		body2.zVel++
	} else if body1.z < body2.z {
		body1.zVel++
		body2.zVel--
	}
	return body1, body2
}

func applyVelocity(moons []moon) []moon {
	for i := 0; i < len(moons); i++ {
		body := &moons[i]
		body.x += body.xVel
		body.y += body.yVel
		body.z += body.zVel
	}
	return moons
}

type moon struct {
	x       int
	y       int
	z       int
	xVel    int
	yVel    int
	zVel    int
	tEnergy int
}

func parseFile(file string) []moon {
	var moons []moon
	content, _ := ioutil.ReadFile(file)
	data := string(content)
	rawProgram := strings.Split(data, "\n")
	for i := 0; i < len(rawProgram); i++ {
		rawMoon := rawProgram[i]
		pre := strings.Split(rawMoon, "=")
		x, _ := strconv.Atoi(strings.Split(pre[1], ",")[0])
		y, _ := strconv.Atoi(strings.Split(pre[2], ",")[0])
		z, _ := strconv.Atoi(strings.Split(pre[3], ">")[0])
		body := moon{x: x, y: y, z: z}
		moons = append(moons, body)
	}
	return moons
}
