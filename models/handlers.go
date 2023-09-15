package models

type Handlers struct {
	Operations []*Operation
}

type Operation struct {
	OperationsPath string
	OperationID    string
	OutputFileName string
	Param          string
}

func NewHandler() *Handlers {
	return &Handlers{}
}

func NewOperation(output, operationID, operationsPath string) *Operation {
	return &Operation{
		OutputFileName: output,
		OperationID:    operationID,
		OperationsPath: operationsPath,
	}
}
