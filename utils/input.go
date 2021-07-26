package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func ToLetters(word string) []string {
	return strings.SplitAfter(word, "")
}

func FindLetterInSlice(letter string, slices []string) map[int]string {
	sl_size := len(slices)
	var indexes_found = make(map[int]string)
	for i := 0; i < sl_size; i++ {
		if slices[i] == letter {
			indexes_found[i] = letter
		}
	}
	return indexes_found
}
