package transaction

type transactionType int

const (
	Inflow = iota
	Outflow
)

func (t transactionType) String() string {
	return [...]string{"Inflow", "Outflow"}[t]
}
