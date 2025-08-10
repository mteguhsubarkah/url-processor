package domain

import "errors"

type OperationType string

const (
    OperationRedirection OperationType = "redirection"
    OperationCanonical   OperationType = "canonical"
    OperationAll         OperationType = "all"
)

func (op OperationType) IsValid() bool {
    switch op {
    case OperationRedirection, OperationCanonical, OperationAll:
        return true
    default:
        return false
    }
}

func ParseOperation(op string) (OperationType, error) {
    operation := OperationType(op)
    if !operation.IsValid() {
        return "", errors.New("invalid operation type")
    }
    return operation, nil
}
