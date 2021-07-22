package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	games "github.com/lLexsonl/games-in-go/games"
)

var games_opt = []string{"ahorcado", "piedra papel tijeras", "triqui", "sudoku", "salir"}
var g_size = len(games_opt)

func main() {

	salir := false
	for !salir {

		menu()
		option := choose()

		switch games_opt[option] {
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
			continue
		}
	}
}

func menu() {
	fmt.Println()
	for i, opt := range games_opt {
		fmt.Printf("%d.. Para %s\n", i+1, opt)
	}
	fmt.Println()
}

func scan() (int, error) {
	var err error
	var opt int
	var option int64

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		option, err = strconv.ParseInt(line, 10, 64)

		if err != nil {
			opt, err = 0, fmt.Errorf("solo ingresar numeros")
		} else {
			opt, err = int(option), nil
		}

	} else {
		opt, err = 0, fmt.Errorf("error inesperado")
	}
	return opt, err
}

func choose() int {
	for {

		option, err := scan()

		if err != nil {
			fmt.Println(err)
			menu()
			continue
		}

		if option > 0 && option <= g_size {
			return option - 1
		}
		fmt.Println("Opcion no valida")
		menu()
	}
}
