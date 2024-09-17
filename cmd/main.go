package main

import (
	"SalesControl/internal/domain/ingredient"
	"SalesControl/internal/domain/recipe"
	"fmt"
)

func main() {
	product := ingredient.NewIngredient("produto1", 3, 1, "unidade", 1)
	product2 := ingredient.NewIngredient("produto2", 3, 1, "unidade", 1)
	product3 := ingredient.NewIngredient("produto3", 3, 1, "unidade", 1)

	ingredients := make(map[ingredient.InterfaceIngredient]int)
	ingredients[product] = 2
	ingredients[product2] = 3
	ingredients[product3] = 1

	recipe1 := recipe.NewRecipe("receita1", ingredients, 5, 10)

	fmt.Println(recipe1)
	fmt.Println(recipe1.GetSaleValue())
}
