package recipe

import (
	"SalesControl/internal/domain/ingredient"
	"fmt"
	"github.com/google/uuid"
)

type recipe struct {
	id             uuid.UUID
	name           string
	ingredients    map[ingredient.InterfaceIngredient]int
	operatingCosts float64
	costValue      float64
	saleValue      float64
	profit         float64
}

type InterfaceRecipe interface {
	GetID() uuid.UUID
	GetName() string
	GetIngredients() map[ingredient.InterfaceIngredient]int
	GetOperatingCosts() float64
	GetCostValue() float64
	GetSaleValue() float64
	GetProfit() float64

	SetName(name string)
	AddIngredient(...ingredient.InterfaceIngredient)
	setCostValue()
	setSaleValue()
	SetProfit(profit float64)

	AlterSalesValue(salesValue float64)
	AlterCostsValue(costValue float64)
	ValidateIngredientsQuantity() bool
}

func NewRecipe(name string, ingredients map[ingredient.InterfaceIngredient]int, operationCosts float64, profit float64) InterfaceRecipe {
	newRecipe := &recipe{
		id:             uuid.New(),
		name:           name,
		ingredients:    ingredients,
		operatingCosts: operationCosts,
		profit:         profit,
	}

	newRecipe.setCostValue()
	newRecipe.setSaleValue()

	return newRecipe
}

func (r *recipe) String() string {
	return fmt.Sprintf("Recipe ID: %s \n Name: %s \n Cost Value: %.2f \n Sale Value: %.2f \n Ingredients: %v \n",
		r.id.String(), r.name, r.costValue, r.saleValue, r.ingredients)
}

func (r *recipe) GetID() uuid.UUID {
	return r.id
}

func (r *recipe) GetName() string {
	return r.name
}

func (r *recipe) SetName(name string) {
	r.name = name
}

func (r *recipe) GetIngredients() map[ingredient.InterfaceIngredient]int {
	return r.ingredients
}

func (r *recipe) GetCostValue() float64 {
	return r.costValue
}

func (r *recipe) setCostValue() {
	var totalValue float64

	for key, quantity := range r.ingredients {
		totalValue += key.GetValue() * float64(quantity)
	}

	r.costValue = totalValue
}

func (r *recipe) setSaleValue() {
	totalCosts := r.costValue + r.operatingCosts

	if r.profit <= 0 {
		r.saleValue = totalCosts
	} else {
		profitMargin := (r.costValue + r.operatingCosts) * (r.profit / 100)

		r.saleValue = profitMargin + totalCosts
	}
}

func (r *recipe) GetSaleValue() float64 {
	return r.saleValue
}

func (r *recipe) SetProfit(profit float64) {
	r.profit = profit
}

func (r *recipe) GetProfit() float64 {
	return r.profit
}

func (r *recipe) GetOperatingCosts() float64 {
	return r.operatingCosts
}

func (r *recipe) SetOperatingCosts(operatingCosts float64) {
	r.operatingCosts = operatingCosts
}

func (r *recipe) ValidateIngredientsQuantity() bool {
	for recipeIngredient, MinQuantity := range r.ingredients {
		if recipeIngredient.GetQuantity() < MinQuantity {
			return false
		}
	}
	return true
}

func (r *recipe) AddIngredient(ingredients ...ingredient.InterfaceIngredient) {
	for _, ing := range ingredients {
		r.ingredients[ing] = ing.GetMinQuantity()
	}
}

func (r *recipe) AlterCostsValue(costs float64) {
	r.costValue = costs
}

func (r *recipe) AlterSalesValue(salesValue float64) {
	r.saleValue = salesValue
}
