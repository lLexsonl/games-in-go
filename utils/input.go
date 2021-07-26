package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Scan() (string, error) {
	var err error
	var text string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	} else {
		err = fmt.Errorf("error inesperado")
	}
	return text, err
}
