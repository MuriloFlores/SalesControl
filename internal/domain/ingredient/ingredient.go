package ingredient

import (
	"github.com/google/uuid"
)

type ingredient struct {
	id          uuid.UUID
	name        string
	quantity    int
	value       float64
	unit        string
	minQuantity int
}

type InterfaceIngredient interface {
	GetID() uuid.UUID
	GetName() string
	GetQuantity() int
	GetValue() float64
	GetUnit() string
	GetMinQuantity() int

	SetName(name string)
	SetQuantity(quantity int)
	SetValue(value float64)
	SetUnit(unit string)
	SetMinQuantity(minQuantity int)

	VerifyMinQuantity() bool
}

func NewIngredient(name string, quantity int, value float64, unit string, minQuantity int) InterfaceIngredient {
	return &ingredient{
		id:          uuid.New(),
		name:        name,
		quantity:    quantity,
		value:       value,
		unit:        unit,
		minQuantity: minQuantity,
	}
}

func (i *ingredient) SetName(name string) {
	i.name = name
}

func (i *ingredient) SetQuantity(quantity int) {
	i.quantity = quantity
}

func (i *ingredient) SetValue(value float64) {
	i.value = value
}

func (i *ingredient) SetUnit(unit string) {
	i.unit = unit
}

func (i *ingredient) SetMinQuantity(minQuantity int) {
	i.minQuantity = minQuantity
}

func (i *ingredient) GetID() uuid.UUID {
	return i.id
}

func (i *ingredient) GetName() string {
	return i.name
}

func (i *ingredient) GetQuantity() int {
	return i.quantity
}

func (i *ingredient) GetValue() float64 {
	return i.value
}

func (i *ingredient) GetUnit() string {
	return i.unit
}

func (i *ingredient) GetMinQuantity() int {
	return i.minQuantity
}

func (i *ingredient) VerifyMinQuantity() bool {
	if i.quantity < i.minQuantity {
		return false
	}
	return true
}
