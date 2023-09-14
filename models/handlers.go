package models

type Handlers struct {
	Operations []*Operation
}
type Operation struct {
	OutputFileName string
	OperaID        string
	Param          string
}

func NewHandler() *Handlers {
	return &Handlers{}
}

func NewOperation(output, operaID string) *Operation {
	return &Operation{
		OutputFileName: output,
		OperaID:        operaID,
	}
}
