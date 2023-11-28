package main

import (
	_ "embed"
	"fmt"
	"maps"
	"math"
	"strings"
)

//go:embed input.txt
var input string

type Ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
}

var ingredientOrder []string

var maxSpoons = 100

func main() {
	ingredients := createIngredients(input)

	recipe := map[string]int{}
	for ingredientName := range ingredients {
		recipe[ingredientName] = 0
	}

	bestRecipe, bestScore := bestRecipe(recipe, ingredients)

	fmt.Printf("Best recipe is %v with a score of %d\n", bestRecipe, bestScore)
}

func bestRecipe(recipe map[string]int, ingredients map[string]Ingredient) (bestRecipe map[string]int, bestScore int) {
	bestRecipe = map[string]int{}
	bestScore = 0
	for i := 0; i < int(math.Pow(float64(maxSpoons), float64(len(recipe)))); i++ {
		potentialRecipe := nextRecipe(recipe)
		potentialScore := getScore(potentialRecipe, ingredients)
		if totalSpoons(potentialRecipe) == maxSpoons && getCalories(potentialRecipe, ingredients) == 500 && potentialScore > bestScore {
			bestRecipe = maps.Clone(potentialRecipe)
			bestScore = potentialScore
		}
	}
	return
}

func nextRecipe(recipe map[string]int) map[string]int {
	carryOver := 1
	for _, ingredient := range ingredientOrder {
		recipe[ingredient] += carryOver
		carryOver = 0
		if recipe[ingredient] > maxSpoons {
			carryOver = 1
			recipe[ingredient] = 0
			continue
		}
		break
	}
	return recipe
}

func totalSpoons(recipe map[string]int) (total int) {
	for _, amount := range recipe {
		total += amount
	}
	return
}

func createIngredients(input string) (ingredients map[string]Ingredient) {
	ingredients = map[string]Ingredient{}
	for _, line := range strings.Split(input, "\n") {
		ingredient := createIngredient(line)
		ingredients[ingredient.name] = ingredient
		ingredientOrder = append(ingredientOrder, ingredient.name)
	}
	return
}

func createIngredient(line string) (ingredient Ingredient) {
	_, err := fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
		&ingredient.name, &ingredient.capacity, &ingredient.durability, &ingredient.flavor, &ingredient.texture, &ingredient.calories)
	if err != nil {
		panic(err)
	}
	return
}

func getScore(recipe map[string]int, ingredients map[string]Ingredient) int {
	var capacity, durability, flavor, texture int
	for ingredientName, amount := range recipe {
		capacity += ingredients[ingredientName].capacity * amount
		durability += ingredients[ingredientName].durability * amount
		flavor += ingredients[ingredientName].flavor * amount
		texture += ingredients[ingredientName].texture * amount
	}

	if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 {
		return 0
	}

	return capacity * durability * flavor * texture
}

func getCalories(recipe map[string]int, ingredients map[string]Ingredient) int {
	var calories int
	for ingredientName, amount := range recipe {
		calories += ingredients[ingredientName].calories * amount
	}
	return calories
}
