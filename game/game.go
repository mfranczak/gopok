package game

type Game struct {
	players []string
}

func (g *Game) AddPlayer(player string) {
	g.players = append(g.players, player)
}

func (g *Game) GetPlayers() []string {
	return g.players
}
