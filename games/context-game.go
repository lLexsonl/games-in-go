package games

type Contex struct {
	game Game
}

func NewGame() *Contex {
	return &Contex{
		game: &Ahorcado{},
	}
}

func (c *Contex) SetGame(game Game) {
	c.game = game
}

func (c *Contex) Play() {
	c.game.Play()
}
