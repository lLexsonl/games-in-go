package main

import (
	"fmt"
	games "github.com/lLexsonl/games-in-go/games"
)

var games_opt = []string{"ahorcado", "piedra papel tijeras", "triqui", "salir"}
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
		default:
			continue
		}
	}
}

func menu() {
	for i, opt := range games_opt {
		fmt.Printf("%d.. Para %s\n", i, opt)
	}
}

func choose() int {
	for {
		fmt.Println("Elige una opcion: ")
		var option int
		fmt.Scanln(&option)
		if option > 0 && option <= g_size {
			return option - 1
		}
		fmt.Println("Opcion no valida")
	}
}
