package main

import (
	"fmt"

	"github.com/lLexsonl/games-in-go/games"
	"github.com/lLexsonl/games-in-go/utils"
)

var GAMES = []string{"ahorcado", "roshambo", "triqui", "sudoku"}

var LENGTH = len(GAMES)
var EXIT = "q"

func main() {

	exit := false
	for !exit {

		showMenu()

		text, err := scanGameOption()
		if err != nil {
			fmt.Println(err)
		}

		option, err := validateGameOption(text)

		if text == EXIT || option == LENGTH {
			exit = true
			continue
		}

		if err != nil {
			fmt.Println(err)
		}

		play(GAMES[option])
	}
}

func play(option string) {
	game := games.NewGame()
	switch option {
	case "roshambo":
		game.SetGame(&games.Rock{})
	case "triqui":
		game.SetGame(&games.Triqui{})
	case "sudoku":
		game.SetGame(&games.SudokuLiang{})
	default:
	}
	game.Play()
}

func showMenu() {
	fmt.Println()
	for i, opt := range GAMES {
		fmt.Printf("%d.. Para %s\n", i+1, opt)
		if i == LENGTH-1 {
			fmt.Printf("%d.. Para salir o 'q'\n", i+2)
		}
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
