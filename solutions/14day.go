package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"math"
	"time"
)

func main() {
	start := time.Now()
	cookbook := parseFile("./puzzledata/14day.txt")
	ore := findOre("FUEL", cookbook, 1)
	fmt.Println("Part 1. Total ore required:", ore)

	mid := time.Now()

	totalFuel := findMaxFuel(1000000000000, cookbook)
	fmt.Println("Part 2. Max Fuel:", totalFuel)

	end := time.Now()
	fmt.Println("Part 1 time:", mid.Sub(start))
	fmt.Println("Part 2 time:", end.Sub(mid))
	fmt.Println("Total time:", end.Sub(start))
}

type recipe struct {
	ingredients []ingredient
	output      ingredient
}

type ingredient struct {
	amount int
	name   string
}

func findMaxFuel(ore int, cookbook []recipe) int {
	currentGuess := ore + 1
	multiplier := float64(ore) 
	currentlyPositive := false
	previouslyPositive := false
	for  {
		currentOre := findOre("FUEL", cookbook, currentGuess)
		currentOreMoreFuel := findOre("FUEL", cookbook, currentGuess + 1)
		leftoverOre := ore - currentOre
		leftoverOreMoreFuel := ore - currentOreMoreFuel

		if leftoverOre >= 0 {
			previouslyPositive = true
		} else {
			previouslyPositive = false
		}

		if leftoverOre >= 0 && leftoverOreMoreFuel < 0 {
			return currentGuess
		}

		if leftoverOre >= 0 {
			currentGuess = int(float64(currentGuess) + multiplier)
		} else {
			currentGuess = int(float64(currentGuess) - multiplier)
		}

		currentOre = findOre("FUEL", cookbook, currentGuess)
		leftoverOre = ore - currentOre

		if leftoverOre >= 0 {
			currentlyPositive = true
		} else {
			currentlyPositive = false
		}

		if currentlyPositive != previouslyPositive {
			multiplier /= 2
		} else {
			multiplier *= 1.5
		}
		
		previouslyPositive = currentlyPositive
	}
}

func findOre(output string, cookbook []recipe, outputAmount int) (totalOre int){
	required := make(map[string]int)
	product := getProduct(output, cookbook)
	product.amount = outputAmount
	required[product.name] = product.amount

	for iterations := 1; ; iterations++{
		for requiredItem, requiredAmount := range(required) {
			if requiredAmount > 0 && requiredItem != "ORE" {
				product = getProduct(requiredItem, cookbook)
				multiplier := int(math.Ceil(float64(requiredAmount) / float64(product.amount)))
				newRequired := getRecipe(product, cookbook)
				for i := 0; i < len(newRequired); i++ {
					nR := newRequired[i]
					nRAmount := required[nR.name]
					nRAmount += nR.amount * multiplier
					required[nR.name] = nRAmount
				}
				required[product.name] -= product.amount * multiplier
			}
		}
		finished := true
		for k, v := range(required) {
			if k != "ORE" && v > 0 {
				finished = false
			}
		}
		if finished {
			break
		}
	}
	return required["ORE"]
}

func findOre1Trillion(output string, cookbook []recipe) (totalFuel int){
	leftoverOre := 1000000000000
	required := make(map[string]int)
	product := getProduct(output, cookbook)
	required[product.name] = product.amount

	for ; leftoverOre > 0; {
		for requiredItem, requiredAmount := range(required) {
			if requiredAmount > 0 && requiredItem != "ORE" {
				product = getProduct(requiredItem, cookbook)
				newRequired := getRecipe(product, cookbook)
				for i := 0; i < len(newRequired); i++ {
					nR := newRequired[i]
					nRAmount := required[nR.name]
					nRAmount += nR.amount
					required[nR.name] = nRAmount
				}
				required[product.name] -= product.amount
			}
		}
		finished := true
		for k, v := range(required) {
			if k != "ORE" && v > 0 {
				finished = false
			}
		}
		if finished {
			leftoverOre -= required["ORE"]
			fuelMade := 1000000
			required["ORE"] = 0
			required["FUEL"] = fuelMade
			totalFuel += fuelMade
			fmt.Println(totalFuel, required)
		}
	}
	fmt.Println(totalFuel)
	return required["ORE"]
}

func getExtra(output ingredient, extras []ingredient) (ingredient, bool) {
	for i := 0; i < len(extras); i++ {
		if extras[i].name == output.name {
			return extras[i], true
		}
	}
	return ingredient{}, false
}

func getProduct(output string, cookbook []recipe) ingredient {
	for i := 0; i < len(cookbook); i++ {
		if output == cookbook[i].output.name {
			return cookbook[i].output
		}
	}
	return ingredient{}
}

func getRecipe(output ingredient, cookbook []recipe) []ingredient {
	for i := 0; i < len(cookbook); i++ {
		if output.name == cookbook[i].output.name {
			return cookbook[i].ingredients
		}
	}
	return []ingredient{}
} 

func parseFile(file string) (cookbook []recipe) {
	content, _ := ioutil.ReadFile(file)
	data := string(content)
	rawProgram := strings.Split(data, "\n")
	for i := 0; i < len(rawProgram); i++ {
		rawRecipe := rawProgram[i]
		ingredients := strings.Split(rawRecipe, " => ")
		ingredientList := strings.Split(ingredients[0], ",")
		r := recipe{}
		for j := 0; j < len(ingredientList); j++ {
			rawIngredient := strings.TrimSpace(ingredientList[j])
			splitIngredients := strings.Split(rawIngredient, " ")
			amount, _ := strconv.Atoi(splitIngredients[0])
			name := splitIngredients[1]
			ing := ingredient{amount: amount, name: name}
			r.ingredients = append(r.ingredients, ing)
		}
		ingredientOutput := strings.Split(ingredients[1], " ")
		amount, _ := strconv.Atoi(ingredientOutput[0])
		name := ingredientOutput[1]
		output := ingredient{amount: amount, name: name}
		r.output = output
		cookbook = append(cookbook, r)
	}
	return cookbook
}
