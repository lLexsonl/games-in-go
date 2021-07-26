package utils

import (
	"fmt"
	"strconv"
)

func ParseToInt(text string) (int, error) {
	option, err := strconv.ParseInt(text, 10, 64)

	if err != nil {
		err = fmt.Errorf("solo ingresar numeros")
	}
	return int(option), err
}
