package main

import (
	"github.com/googollee/go-socket.io"
	"github.com/mfranczak/gopok/game"
	"log"
	"net/http"
)

func main() {
	var game = new(game.Game)

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		so.Join("game")

		so.On("estimate", func(msg string) {
			log.Println(msg)
			so.BroadcastTo("game", "estimate", msg)
		})

		so.On("new_player", func(msg string) {
			game.AddPlayer(msg)
			so.Emit("new_player", game.GetPlayerNames())
			so.BroadcastTo("game", "new_player", game.GetPlayerNames())
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
