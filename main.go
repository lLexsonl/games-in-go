package main

import (
	"fmt"

	"github.com/lLexsonl/games-in-go/games"
	"github.com/lLexsonl/games-in-go/utils"
)

var GAMES = []string{"ahorcado", "roshambo", "triqui", "sudoku", "salir"}

var LENGTH = len(GAMES)

func main() {

	exit := false
	for !exit {

		showMenu()

		text, err := scanGameOption()
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

func play(option string) bool {
	var exit = false
	game := games.NewGame()
	switch option {
	case "roshambo":
		game.SetGame(&games.Rock{})
	case "triqui":
		game.SetGame(&games.Triqui{})
	case "sudoku":
		game.SetGame(&games.SudokuLiang{})
	case "salir":
		exit = true
	default:
	}
	game.Play()
	return exit
}

func showMenu() {
	fmt.Println()
	for i, opt := range GAMES {
		fmt.Printf("%d.. Para %s\n", i+1, opt)
	}
	fmt.Println()
}

func scanGameOption() (string, error) {
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
