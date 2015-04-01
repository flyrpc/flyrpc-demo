package main

import (
	"log"

	"github.com/flyrpc/flyrpc"
	. "github.com/flyrpc/flyrpc-demo/go/msg"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	server := fly.NewServer(&fly.ServerOpts{
		Serializer: fly.Protobuf,
	})

	server.OnConnect(func(ctx *fly.Context) {
		ctx.SendMessage(1, &Hello{Id: 567, Name: "abc"})
	})

	server.OnMessage(1, func(ctx *fly.Context, u *Hello) *Hello {
		u.Id += 1
		return u
	})

	server.OnMessage(2, func(ctx *fly.Context, u *Hello) {
		u.Id += 300
		ctx.SendMessage(3, u)
	})

	server.OnMessage(5, func(ctx *fly.Context, u *Hello) {
		log.Println("server on message 5", u)
		u.Id += 500
		ctx.SendMessage(6, u)
	})

	err := server.Listen("0.0.0.0:5555")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server listen at 5555")
	server.HandleConnections()
}
