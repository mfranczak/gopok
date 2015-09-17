package main

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

type Game struct {
	players []string
}

func (g *Game) AddPlayer(player string) {
	g.players = append(g.players, player)
}

func main() {
	var g = new(Game)

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		so.Join("game")
		so.BroadcastTo("game", "new_player", g.players)

		so.On("estimate", func(msg string) {
			log.Println(msg)
			so.BroadcastTo("game", "estimate", msg)
		})

		so.On("new_player", func(msg string) {
			g.AddPlayer(msg)
			so.Emit("new_player", g.players)
			so.BroadcastTo("game", "new_player", g.players)
		})

		so.On("disconnection", func() {
			log.Println("player left")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}
