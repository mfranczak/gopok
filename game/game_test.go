package game

import (
	"testing"
)

func TestAddPlayers(t *testing.T) {
	var g = new(Game)
	g.AddPlayer("player")

	var players = g.GetPlayers()

	if players[0] != "player" {
		t.Errorf("Player on 0 is wrong %q", players[0])
	}
}
