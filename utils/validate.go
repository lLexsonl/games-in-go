package utils

import "fmt"

func IntInRange(number int, lower int, upper int) bool {

	if number > lower && number <= upper {
		return true
	}
	return false
}

func RangeToOption(option int, lower int, upper int) (int, error) {
	var err error
	var opt int
	if IntInRange(option, lower, upper) {
		opt = option - 1
	} else {
		err = fmt.Errorf("opcion no valida")
	}
	return opt, err
}
