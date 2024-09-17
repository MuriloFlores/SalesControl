package order

import (
	"SalesControl/internal/domain/client"
	"SalesControl/internal/domain/recipe"
	"github.com/google/uuid"
	"time"
)

type order struct {
	id           uuid.UUID
	client       client.InterfaceClient
	recipes      []recipe.InterfaceRecipe
	deliveryDate time.Time
	totalValue   float64
	shippingCost float64
	status       StatusOrder
}

type InterfaceOrder interface {
	GetID() uuid.UUID
	GetClient() client.InterfaceClient
	GetRecipes() []recipe.InterfaceRecipe
	GetDeliveryDate() time.Time
	GetTotalValue() float64
	GetShippingCost() float64
	GetStatus() StatusOrder

	SetClient(client client.InterfaceClient)
	SetDeliveryDate(deliveryDate time.Time)
	SetTotalValue(totalValue float64)
	SetShippingCost(shippingCost float64)
	SetStatus(status StatusOrder)

	AddRecipe(recipe ...recipe.InterfaceRecipe)
	RemoveRecipe(shouldRemove func(recipe ...recipe.InterfaceRecipe) bool)
}

func NewOrder(client client.InterfaceClient, recipes []recipe.InterfaceRecipe, deliveryData time.Time, totalValue float64, shippingCost float64, status StatusOrder) InterfaceOrder {
	return &order{
		id:           uuid.New(),
		client:       client,
		recipes:      recipes,
		deliveryDate: deliveryData,
		totalValue:   totalValue,
		shippingCost: shippingCost,
		status:       status,
	}
}

func (o *order) GetID() uuid.UUID {
	return o.id
}

func (o *order) GetClient() client.InterfaceClient {
	return o.client
}

func (o *order) GetRecipes() []recipe.InterfaceRecipe {
	return o.recipes
}

func (o *order) GetDeliveryDate() time.Time {
	return o.deliveryDate
}

func (o *order) GetTotalValue() float64 {
	return o.totalValue
}

func (o *order) GetShippingCost() float64 {
	return o.shippingCost
}

func (o *order) GetStatus() StatusOrder {
	return o.status
}

func (o *order) SetClient(client client.InterfaceClient) {
	o.client = client
}

func (o *order) SetDeliveryDate(date time.Time) {
	o.deliveryDate = date
}

func (o *order) SetTotalValue(totalValue float64) {
	o.totalValue = totalValue
}

func (o *order) SetShippingCost(shippingCost float64) {
	o.shippingCost = shippingCost
}

func (o *order) SetStatus(status StatusOrder) {
	o.status = status
}

func (o *order) AddRecipe(recipe ...recipe.InterfaceRecipe) {
	o.recipes = append(o.recipes, recipe...)
}

func (o *order) RemoveRecipe(shouldRemove func(recipes ...recipe.InterfaceRecipe) bool) {
	newRecipes := o.recipes[:0]

	for _, rec := range o.recipes {
		if !shouldRemove(rec) {
			newRecipes = append(newRecipes, rec)
		}
	}

	o.recipes = newRecipes
}
