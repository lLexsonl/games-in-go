package games

import (
	"fmt"
	"math/rand"
)

type Rock struct {
}

func (r *Rock) Play() {
	rock()
}

// global
var options = [3]string{"rock", "paper", "scissors"}
var wins, losses, ties int
var size = len(options)
var option_s = map[string]string{
	"r": "rock",
	"p": "paper",
	"s": "scissors",
}

func rock() {
	var option string
	for {
		fmt.Println("Rock, paper, scissors")
		fmt.Println("Enter (r)ock, (p)aper, (s)cissors or (q)uit")

		fmt.Scanln(&option)

		if option == "q" {
			break
		}

		option_comple := option_s[option]

		if !in_options(option_comple) {
			fmt.Println("Bad option!")
			continue
		}

		option_rand := options[rand.Intn(size)]

		fmt.Printf("%s vs %s\n", option_comple, option_rand)

		who_win(option_comple, option_rand)
	}
}

func who_win(you string, me string) {
	you_win := false

	if you == me {
		fmt.Println("Empateeee!")
		ties++
		fmt.Printf("Wins: %d Losses: %d Ties: %d\n", wins, losses, ties)
		return
	} else if you == options[0] {
		if me == options[2] {
			you_win = true
		}
	} else if you == options[1] {
		if me == options[0] {
			you_win = true
		}
	} else {
		if me == options[1] {
			you_win = true
		}
	}

	if you_win {
		fmt.Println("Ganasteee!")
		wins++
	} else {
		fmt.Println("Perdisteee!")
		losses++
	}

	fmt.Printf("Wins: %d Losses: %d Ties: %d\n", wins, losses, ties)
}

func in_options(option string) bool {

	for i := 0; i < size; i++ {
		if option == options[i] {
			return true
		}
	}
	return false
}
