package game

type Game struct {
	Players []string
}

func (g *Game) AddPlayer(player string) {
	g.Players = append(g.Players, player)
}
