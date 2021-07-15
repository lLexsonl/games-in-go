package games

import (
	"fmt"
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

var size_img int = len(IMAGES)

func Ahorcado() {
	for {
		var palabra string
		fmt.Print("\nIngrese la palabra a adivinar o 'q' para salir: ")
		fmt.Scanln(&palabra)

		if palabra == "q" {
			break
		}

		slices := strings.SplitAfter(palabra, "")
		sl_size := len(slices)
		secret := make([]string, sl_size)
		fmt.Printf("La palabra tiene %d letras\n", sl_size)
		salir := false
		vidas := size

		for i := 0; i < sl_size; i++ {
			secret[i] = "_"
		}
		fmt.Println()

		for !salir {
			for i := 0; i < sl_size; i++ {
				fmt.Printf(" %s ", secret[i])
			}
			fmt.Println()

			var letra string
			fmt.Print("Ingrese una letra: ")
			fmt.Scanln(&letra)

			indexes_found := FindIndexes(slices, letra)
			if len(indexes_found) > 0 {
				for key, value := range indexes_found {
					secret[key] = value
				}
			} else {
				fmt.Print(IMAGES[size_img-vidas])
				vidas -= 1
			}

			secret_word := strings.Join(secret, "")

			if palabra == secret_word {
				fmt.Printf("\nLa palabra es: %s\n", palabra)
				fmt.Println("٩(^‿^)۶\nGanasteee!")
				salir = true
			}

			if vidas <= 0 {
				fmt.Printf("\nLa palabra era: %s\n\n(╥﹏╥)\nPerdisteee!\n", palabra)
				salir = true
			}
		}
	}
}

func FindIndexes(slices []string, letra string) map[int]string {
	sl_size := len(slices)
	var indexes_found = make(map[int]string)
	for i := 0; i < sl_size; i++ {
		if slices[i] == letra {
			indexes_found[i] = letra
		}
	}
	return indexes_found
}
