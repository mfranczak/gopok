package main

import (
	"github.com/googollee/go-socket.io"
	"github.com/mfranczak/gopok/game"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func pick(index int, args []string) string {
	return args[index]
}

func main() {
	var game = new(game.Game)

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		so.Join("game")

		so.On("estimate", func(msg string) {
			var data = strings.Split(msg, ":")
			var vote, err = strconv.Atoi(pick(1, data))

			if err == nil {
				log.Printf("New vote %s %d", data[0], vote)
				game.Vote(data[0], vote)
				so.BroadcastTo("game", "estimate", msg)
			}
		})

		so.On("new_player", func(msg string) {
			game.AddPlayer(msg)
			so.Emit("new_player", game.GetPlayers())
			so.BroadcastTo("game", "new_player", game.GetPlayers())
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
