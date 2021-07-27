package utils

func IntInRange(number int, lower int, upper int) bool {

	if number >= lower && number < upper {
		return true
	}
	return false
}

func OptionInRange(option int, lower int, upper int) bool {
	return IntInRange(option, lower, upper)
}
