package game

import (
	"testing"
)

func TestAddPlayers(t *testing.T) {
	var g = new(Game)
	g.AddPlayer("player")

	var players = g.GetPlayerNames()

	if players[0] != "player" {
		t.Errorf("Player on 0 is wrong %q", players[0])
	}
}

func TestVoting(t *testing.T) {
	var game = new(Game)
	game.AddPlayer("player")
	game.Vote("player", 10)

	var players = game.GetPlayers()
	if players[0].Name != "player" {
		t.Error("Wrong player name")
	}

	if players[0].Vote != 10 {
		t.Error("Player vote was not stored")
	}

}
