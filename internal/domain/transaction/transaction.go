package transaction

import (
	"github.com/google/uuid"
	"time"
)

type transaction struct {
	id               uuid.UUID
	transactionType  transactionType
	transactionValue float64
	date             time.Time
	description      string
}

type InterfaceTransaction interface {
	GetId() uuid.UUID
	GetTransactionType() transactionType
	GetTransactionValue() float64
	GetDate() time.Time
	GetDescription() string

	SetTransactionType(transactionType transactionType)
	SetTransactionValue(value float64)
	SetDate(date time.Time)
	SetDescription(description string)
}

func NewTransaction(transactionType transactionType, transactionValue float64, date time.Time, description string) InterfaceTransaction {
	return &transaction{
		id:               uuid.New(),
		transactionType:  transactionType,
		transactionValue: transactionValue,
		date:             date,
		description:      description,
	}
}

func (t *transaction) GetId() uuid.UUID {
	return t.id
}

func (t *transaction) GetTransactionType() transactionType {
	return t.transactionType
}

func (t *transaction) GetTransactionValue() float64 {
	return t.transactionValue
}

func (t *transaction) GetDate() time.Time {
	return t.date
}

func (t *transaction) GetDescription() string {
	return t.description
}

func (t *transaction) SetTransactionType(transaction transactionType) {
	t.transactionType = transaction
}

func (t *transaction) SetTransactionValue(value float64) {
	t.transactionValue = value
}

func (t *transaction) SetDate(date time.Time) {
	t.date = date
}

func (t *transaction) SetDescription(description string) {
	t.description = description
}
