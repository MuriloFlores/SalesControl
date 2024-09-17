package cash_register

import (
	"SalesControl/internal/domain/transaction"
	"github.com/google/uuid"
)

type cashRegister struct {
	id          uuid.UUID
	balance     float64
	transaction []transaction.InterfaceTransaction
}

type InterfaceCashRegister interface {
	GetId() uuid.UUID
	GetBalance() float64
	GetTransaction() []transaction.InterfaceTransaction
	AddTransaction(transaction ...transaction.InterfaceTransaction)
}

func NewCashRegister(balance float64, transaction []transaction.InterfaceTransaction) InterfaceCashRegister {
	register := &cashRegister{
		id:          uuid.New(),
		transaction: transaction,
	}

	if balance < 0 {
		register.balance = 0.00
	}

	register.balance = balance

	return register
}

func (c *cashRegister) GetId() uuid.UUID {
	return c.id
}

func (c *cashRegister) GetBalance() float64 {
	return c.balance
}

func (c *cashRegister) GetTransaction() []transaction.InterfaceTransaction {
	return c.transaction
}

func (c *cashRegister) AddTransaction(transaction ...transaction.InterfaceTransaction) {
	c.transaction = append(c.transaction, transaction...)
}
