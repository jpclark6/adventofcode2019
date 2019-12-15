package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	cookbook := parseFile("./puzzledata/14day.txt")
	ore := findOre("FUEL", cookbook)
	fmt.Println("Total ore required:", ore)
}

type recipe struct {
	ingredients []ingredient
	output      ingredient
}

type ingredient struct {
	amount int
	name   string
}

func findOre(output string, cookbook []recipe) (totalOre int){
	required := make(map[string]int)
	product := getProduct(output, cookbook)
	required[product.name] = product.amount

	for iterations := 1; ; iterations++{
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
			break
		}
	}
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
