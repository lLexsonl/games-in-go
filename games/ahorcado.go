package games

import (
	"fmt"
	"github.com/lLexsonl/games-in-go/utils"
	"strings"
)

var IMAGES = [5]string{
	`
  +---+
  |   |
      |
      |
      |
      |
=========
`,
	`
  +---+
  |   |
  O   |
      |
      |
      |
=========
`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========
`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========
`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========
`,
}

const size_img int = len(IMAGES)

const (
	unknown_letter string = "_"
	letter_to_exit string = "q"
)

type Ahorcado struct{}

func (a *Ahorcado) Play() {
	ahorcado()
}

func ahorcado() {
	for {
		var word string
		var err error

		word, err = scanWord()

		if err != nil {
			fmt.Println(err)
			word = "q"
		}

		if word == "q" {
			break
		}

		word_size := len(word)
		fmt.Printf("La palabra tiene %d letras\n", word_size)

		playGame(word, word_size)

	}
}

func fillAllOcurrences(secret []string, indexes_found map[int]string) string {
	for key, value := range indexes_found {
		secret[key] = value
	}
	secret_word := strings.Join(secret, "")
	return secret_word
}

func findAllOcurrences(word string, serched string) map[int]string {
	var letters_found = make(map[int]string)
	for index, lett := range word {
		letter := string(lett)
		if letter == serched {
			letters_found[index] = letter
		}
	}
	return letters_found
}

func createSecret(size int, unknown_letter string) []string {
	secret := make([]string, size)
	for i := 0; i < size; i++ {
		secret[i] = unknown_letter
	}
	return secret
}

func scanWord() (string, error) {
	return utils.Scan("\nIngrese la palabra a adivinar o 'q' para salir: ")
}

func playGame(word string, word_size int) {
	secret := createSecret(word_size, unknown_letter)
	exit := false
	lifes := size_img

	var letter string
	var secret_word string
	var err error

	for !exit {

		fmt.Println(secret)

		letter, err = utils.Scan("Ingrese una letra: ")

		if err != nil {
			fmt.Println(err)
			exit = true
			continue
		}

		indexes_found := findAllOcurrences(word, letter)

		if len(indexes_found) == 0 {
			fmt.Print(IMAGES[size_img-lifes])
			lifes--
		}

		secret_word = fillAllOcurrences(secret, indexes_found)

		if win(word, secret_word) {
			exit = true
		}

		if lose(word, lifes) {
			exit = true
		}
	}
}

func win(word string, secret_word string) bool {
	if word == secret_word {
		fmt.Printf("\nLa palabra es: %s\n", word)
		fmt.Println("٩(^‿^)۶\nGanasteee!")
		return true
	}
	return false
}

func lose(word string, vidas int) bool {
	if vidas <= 0 {
		fmt.Printf("\nLa palabra era: %s\n\n(╥﹏╥)\nPerdisteee!\n", word)
		return true
	}
	return false
}
