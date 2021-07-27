package main

import (
	"fmt"

	"github.com/lLexsonl/games-in-go/games"
	"github.com/lLexsonl/games-in-go/utils"
)

var GAMES = []string{"ahorcado", "piedra papel tijeras", "triqui", "sudoku", "salir"}

var LENGTH = len(GAMES)

func main() {

	exit := false
	for !exit {

		showMenu()

		text, err := scanOptionGame()
		if err != nil {
			fmt.Println(err)
		}
		option, err := validateGameOption(text)
		if err != nil {
			fmt.Println(err)
		}

		exit = play(GAMES[option])
	}
}

func play(game string) bool {
	var exit = false
	switch game {
	case "ahorcado":
		games.Ahorcado()
	case "piedra papel tijeras":
		games.Rock()
	case "triqui":
		games.Triqui()
	case "salir":
		exit = true
	case "sudoku":
		games.SudokuLiang()
	default:
	}
	return exit
}

func showMenu() {
	fmt.Println()
	for i, opt := range GAMES {
		fmt.Printf("%d.. Para %s\n", i+1, opt)
	}
	fmt.Println()
}

func scanOptionGame() (string, error) {
	return utils.Scan("Elige un juego: ")
}

func validateGameOption(text string) (int, error) {
	var err error
	var option int

	option, err = utils.ParseToInt(text)
	if err == nil {
		option--
		if !utils.OptionInRange(option, 0, LENGTH) {
			err = fmt.Errorf("opcion no valida")
		}
	}
	return option, err
}
