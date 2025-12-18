package wordy

import (
	"errors"
	"strconv"
	"strings"
)

const PREFIX = "What is "
const SUFFIX = "?"

func Answer(question string) (int, bool) {
	str, err := removePrefixSuffix(question)
	if err != nil {
		return 0, false
	}

	if i, ok := checkNumberIteration1(str); ok {
		return i, true
	}

	operations, err := extractOperations(str)
	if err != nil {
		return 0, false
	}

	result, err := calculate(operations)
	if err != nil {
		return 0, false
	}

	return result, true
}

type Operation struct {
	name *string
	n1   *int
	n2   *int
}

func (op *Operation) isFullyInitiated() bool {
	return op.name != nil && op.n1 != nil && op.n2 != nil
}
func (op *Operation) n1MustBeFilled() bool {
	return op.n1 == nil && op.name == nil && op.n2 == nil
}
func (op *Operation) n2MustBeFilled() bool {
	return op.n1 != nil && op.name != nil && op.n2 == nil
}
func (op *Operation) nameMustBeFilled() bool {
	return op.n1 != nil && op.name == nil && op.n2 == nil
}
func (op *Operation) nameAcceptCompletion() bool {
	return op.n2MustBeFilled()
}
func (op *Operation) result() (int, error) {
	switch *op.name {
	case "plus":
		return *op.n1 + *op.n2, nil
	case "minus":
		return *op.n1 - *op.n2, nil
	case "multiplied by":
		return *op.n1 * *op.n2, nil
	case "divided by":
		return *op.n1 / *op.n2, nil
	default:
		return 0, errors.New("Syntax Error")
	}
}

var validOpName = map[string]bool{
	"plus":       true,
	"minus":      true,
	"multiplied": true,
	"divided":    true,
	"by":         true,
}

func extractOperations(str string) (list []Operation, err error) {
	split := strings.Split(str, " ")
	list = []Operation{{}}
	for _, item := range split {
		lastOp := &list[len(list)-1]
		if lastOp.isFullyInitiated() {
			list = append(list, Operation{n1: lastOp.n2})
			lastOp = &list[len(list)-1]
		}

		i, numErr := strconv.Atoi(item)
		isNum := numErr == nil
		switch {
		case validOpName[item]:
			switch {
			case lastOp.nameAcceptCompletion() && item == "by":
				*lastOp.name += " by"
			case lastOp.nameMustBeFilled():
				lastOp.name = &item
			default:
				return list, errors.New("Syntax Error")
			}
		case isNum:
			switch {
			case lastOp.n1MustBeFilled():
				lastOp.n1 = &i
			case lastOp.n2MustBeFilled():
				lastOp.n2 = &i
			default:
				return list, errors.New("Syntax Error")
			}
		default:
			return list, errors.New("Syntax Error")

		}
	}
	for _, op := range list {
		if !op.isFullyInitiated() {
			return list, errors.New("Syntax Error")
		}
	}
	return
}

func calculate(operations []Operation) (int, error) {
	currNumber := *operations[0].n1
	for _, op := range operations {
		*op.n1 = currNumber
		result, err := op.result()
		if err != nil {
			return 0, errors.New("Syntax Error")
		}
		currNumber = result
	}
	return currNumber, nil
}

func checkNumberIteration1(str string) (int, bool) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, false
	}
	return i, true
}

func removePrefixSuffix(question string) (str string, err error) {
	if len(question) <= len(PREFIX) {
		return str, errors.New("Syntax Error")
	}
	if question[:len(PREFIX)] == PREFIX && question[len(question)-1:] == SUFFIX {
		str = question[len(PREFIX) : len(question)-len(SUFFIX)]
		return str, err
	}
	return str, errors.New("Syntax Error")
}
