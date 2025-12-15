package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	amount, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	multiplier, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return multiplier * float64(amount) / float64(cows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) (err error) {
	if cows < 0 {
		cowError := &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
		return cowError
	} else if cows == 0 {
		cowError := &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
		return cowError
	}
	return
}
