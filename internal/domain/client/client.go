package client

import (
	"SalesControl/internal/domain/order"
	"github.com/google/uuid"
)

type client struct {
	id           uuid.UUID
	name         string
	address      string
	phoneNumber  string
	orderHistory []order.InterfaceOrder
}

type InterfaceClient interface {
	GetId() uuid.UUID
	GetName() string
	GetAddress() string
	GetPhoneNumber() string
	GetOrderHistory() []order.InterfaceOrder

	SetName(name string)
	SetAddress(address string)
	SetPhoneNumber(phoneNumber string)

	AddOrder(order order.InterfaceOrder)
	RemoveOrder(shouldRemove func(order order.InterfaceOrder) bool)
}

func NewOrder(name string, phoneNumber string, address string) InterfaceClient {
	return &client{
		id:          uuid.New(),
		name:        name,
		address:     address,
		phoneNumber: phoneNumber,
	}
}

func (c *client) GetId() uuid.UUID {
	return c.id
}

func (c *client) GetName() string {
	return c.name
}

func (c *client) GetAddress() string {
	return c.address
}

func (c *client) GetPhoneNumber() string {
	return c.phoneNumber
}

func (c *client) GetOrderHistory() []order.InterfaceOrder {
	return c.orderHistory
}

func (c *client) AddOrder(order order.InterfaceOrder) {
	c.orderHistory = append(c.orderHistory, order)
}

func (c *client) RemoveOrder(shouldRemove func(order order.InterfaceOrder) bool) {
	resultSlice := c.orderHistory[:0]

	for _, value := range c.orderHistory {
		if !shouldRemove(value) {
			resultSlice = append(resultSlice, value)
		}
	}

	c.orderHistory = resultSlice
}

func (c *client) SetName(name string) {
	c.name = name
}

func (c *client) SetAddress(address string) {
	c.address = address
}

func (c *client) SetPhoneNumber(phoneNumber string) {
	c.phoneNumber = phoneNumber
}
