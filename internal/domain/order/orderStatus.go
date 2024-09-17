package order

type StatusOrder int

const (
	Canceled StatusOrder = iota
	Sent
	Preparing
	UnderAnalysis
)

func (s StatusOrder) String() string {
	return [...]string{"Canceled", "Sent", "Preparing", "UnderAnalysis"}[s]
}
