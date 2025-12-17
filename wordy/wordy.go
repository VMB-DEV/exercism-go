package wordy

import (
	"errors"
	"strconv"
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

	println(str)
	return 0, false
}

func checkNumberIteration1(str string) (int, bool) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return i, true
	}
	return 0, false
}

func removePrefixSuffix(question string) (str string, err error) {
	if len(question) <= len(PREFIX) {
		return str, errors.New("Syntax Error")
	}
	if question[:len(PREFIX)-1] == PREFIX || question[len(question)-1:] == SUFFIX {
		str = question[len(PREFIX) : len(question)-len(SUFFIX)]
		println(str)
		return str, err
	}
	return str, errors.New("Syntax Error")
}
