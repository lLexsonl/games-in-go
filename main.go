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

func chooseGame() (int, error) {
	var text string
	var err error
	var option int

	text, err = utils.Scan("Elige un juego: ")
	if err == nil {
		option, err = utils.ParseToInt(text)
		if err == nil {
			option, err = utils.RangeToOption(option, 0, LENGTH)
		}
	}
	return option, err
}
