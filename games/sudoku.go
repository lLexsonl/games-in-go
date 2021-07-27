package games

import (
	"fmt"
	"math/rand"
)

type Sudoku struct{}

func (s *Sudoku) Play() {
	sudoku()
}

var null = 0
var table_sudo = [][]int{
	{null, null, null},
	{null, null, null},
	{null, null, null},
}
var sudo_size = len(table_sudo)

func sudoku() {

	rand_initial_table()
	show_table()

}

func show_table() {
	for _, row := range table_sudo {
		fmt.Println(row)
	}
}

func rand_initial_table() {

	for i := 0; i < sudo_size; i++ {
		one := rand.Intn(sudo_size)
		two := rand.Intn(sudo_size)
		three := rand.Intn(sudo_size) + 1
		table_sudo[one][two] = three
	}
}
