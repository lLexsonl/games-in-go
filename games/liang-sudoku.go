package games

import (
	"fmt"
)

type SudokuLiang struct{}

func (s *SudokuLiang) Play() {
	sudokuLiang()
}

var grid = [][]int{
	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 4, 5, 2, 8, 6, 1, 7, 9},
}

var _ = [][]int{
	{9, 6, 3, 1, 7, 4, 2, 5, 8},
	{1, 7, 8, 3, 2, 5, 6, 4, 9},
	{2, 5, 4, 6, 8, 9, 7, 3, 1},
	{8, 2, 1, 4, 3, 7, 5, 9, 6},
	{4, 9, 6, 8, 5, 2, 3, 1, 7},
	{7, 3, 5, 9, 6, 1, 8, 2, 4},
	{5, 8, 9, 7, 1, 3, 4, 6, 2},
	{3, 1, 7, 2, 4, 6, 9, 8, 5},
	{6, 4, 2, 5, 9, 8, 1, 7, 3},
}

func readASolution() {
	var inp int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Scanf("%d", &inp)
			grid[i][j] = inp
		}
	}
}

func isValid(grid [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] < 1 || grid[i][j] > 9 || !isValido(i, j, grid) {
				return false
			}
		}
	}
	return true
}

func isValido(i int, j int, grid [][]int) bool {
	for col := 0; col < 9; col++ {
		if col != j && grid[i][col] == grid[i][j] {
			return false
		}
	}

	for row := 0; row < 9; row++ {
		if row != i && grid[row][j] == grid[i][j] {
			return false
		}
	}

	for row := (i / 3) * 3; row < (i/3)*3+3; row++ {
		for col := (j / 3) * 3; col < (j/3)*3+3; col++ {
			if !(row == i && col == j) && grid[row][col] == grid[i][j] {
				return false
			}
		}
	}
	return true
}

func show_grid() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

func sudokuLiang() {
	fmt.Println("Sudoku")
	show_grid()
	fmt.Printf("Es solucion: %t\n", isValid(grid))
}
