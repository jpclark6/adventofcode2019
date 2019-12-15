package main

import (
	"fmt"
	"io/ioutil"
	// "math"
	"strconv"
	"strings"
	// "time"
)

func main() {
	cookbook := parseFile("./puzzledata/14day.txt")
	fmt.Println(cookbook)
}

type recipe struct {
	ingredients []ingredient
	output      ingredient
}

type ingredient struct {
	amount int
	name   string
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
