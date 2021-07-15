package games

import (
	"fmt"
)

var wins_table = [][][]int{
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
	{{0, 0}, {1, 1}, {2, 2}},
	{{0, 2}, {1, 1}, {2, 0}},
}

const X string = "X"
const O string = "O"
const G string = "_"

var table = [][]string{
	{G, G, G},
	{G, G, G},
	{G, G, G},
}

var turno = O

func Triqui() {

	for {
		var x_coo, y_coo int
		for {
			fmt.Println("Ingrese posicion en x: ")
			fmt.Scanln(&x_coo)
			if x_coo >= 0 && x_coo < 3 {
				break
			}
		}
		for {
			fmt.Println("Ingrese posicion en y: ")
			fmt.Scanln(&y_coo)
			if y_coo >= 0 && y_coo < 3 {
				break
			}
		}

		if table[x_coo][y_coo] == G {
			if turno == X {
				table[x_coo][y_coo] = X
				turno = O
			} else {
				table[x_coo][y_coo] = O
				turno = X
			}
		} else {
			fmt.Println("Ya ha sido utilizada esa posicion")
			continue
		}

		is_full := table_p()

		win, winner := winner()

		if win {
			fmt.Printf("Triqui -> '%s' gana! <-", winner)
			break
		}

		if is_full {
			fmt.Println("Triqui -> Empate! <-", winner)
			break
		}

	}
}

func table_p() bool {
	counter := 0
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[0]); j++ {
			fmt.Printf(" %s ", table[i][j])
			if table[i][j] == G {
				counter++
			}
		}
		fmt.Println()
	}
	return counter == 0
}

func winner() (bool, string) {
	var win = false
	var winner string
	var one, two, three string
	for i := 0; i < len(wins_table) && !win; i++ {
		one = table[wins_table[i][0][0]][wins_table[i][0][1]]
		two = table[wins_table[i][1][0]][wins_table[i][1][1]]
		three = table[wins_table[i][2][0]][wins_table[i][2][1]]

		if (one != G && two != G && three != G) && (one == two && one == three) {
			win = true
		}
	}

	if win {
		if one == X {
			winner = X
		} else {
			winner = O
		}
	}
	return win, winner
}
