package main

import (
	"log"

	"github.com/flyrpc/flyrpc"
	. "github.com/flyrpc/flyrpc-demo/go/msg"
)

func main() {
	server := fly.NewServer(&fly.ServerOpts{
		Serializer: fly.Protobuf,
	})

	server.OnMessage(1, func(ctx *fly.Context, u *Hello) *Hello {
		u.Id += 1
		return u
	})

	server.OnMessage(2, func(ctx *fly.Context, u *Hello) {
		u.Id += 300
		ctx.SendMessage(3, u)
	})

	err := server.Listen("0.0.0.0:5555")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server listen at 5555")
	server.HandleConnections()
}
