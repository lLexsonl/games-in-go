package main

import (
	"fmt"
	"github.com/llexsonl/ahorcado"
)

var games = []string{"ahorcado", "piedra papel tijeras", "triqui", "salir"}
var g_size = len(games)

func main() {

	salir := false
	for !salir {

		menu()
		option := choose()

		switch games[option] {
		case "ahorcado":
			Ahorcado()
		case "piedra papel tijeras":
			Rock()
		case "triqui":
			Triqui()
		case "salir":
			salir = true
		default:
			continue
		}
	}
}

func menu() {
	for i, opt := range games {
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
