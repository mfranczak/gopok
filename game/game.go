package game

import (
	"log"
)

type Game struct {
	players []*Player
}

type Player struct {
	Name string
	Vote int
}

func (g *Game) AddPlayer(playerName string) {
	var player = new(Player)
	player.Name = playerName
	player.Vote = -1

	g.players = append(g.players, player)
}

// go:deprecated
func (g *Game) GetPlayerNames() []string {
	var playerNames []string
	for _, player := range g.players {
		playerNames = append(playerNames, player.Name)
	}
	return playerNames
}

func (g *Game) GetPlayers() []*Player {
	return g.players
}

func (g *Game) Vote(playerName string, vote int) {
	for _, player := range g.players {
		if player.Name == playerName {
			player.Vote = vote
			return
		}
	}

	log.Printf("[ERROR] problem with voting player=%s, vote=%d", playerName, vote)
}
