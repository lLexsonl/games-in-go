package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/lLexsonl/games-in-go/games"
)

var GAMES = []string{"ahorcado", "piedra papel tijeras", "triqui", "sudoku", "salir"}

var LENGTH = len(GAMES)

func main() {

	exit := false
	for !exit {

		showMenu()
		game, err := chooseGame()

		if err != nil {
			fmt.Println(err)
			continue
		}
		exit = play(game)
	}
}

func play(index int) bool {
	var salir = false
	switch GAMES[index] {
	case "ahorcado":
		games.Ahorcado()
	case "piedra papel tijeras":
		games.Rock()
	case "triqui":
		games.Triqui()
	case "salir":
		salir = true
	case "sudoku":
		games.SudokuLiang()
	default:
	}
	return salir
}

func showMenu() {
	fmt.Println()
	for i, opt := range GAMES {
		fmt.Printf("%d.. Para %s\n", i+1, opt)
	}
	fmt.Println()
}

func parseInt(text string) (int, error) {
	option, err := strconv.ParseInt(text, 10, 64)

	if err != nil {
		err = fmt.Errorf("solo ingresar numeros")
	}
	return int(option), err
}

func scan() (string, error) {
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

func validateOption(option int) (int, error) {
	var err error

	if option > 0 && option <= LENGTH {
		option = option - 1
	} else {
		err = fmt.Errorf("opcion no valida")
	}
	return option, err
}

func chooseGame() (int, error) {
	var text string
	var err error
	var option int

	text, err = scan()
	if err == nil {
		option, err = parseInt(text)
		if err == nil {
			option, err = validateOption(option)
		}
	}
	return option, err
}
