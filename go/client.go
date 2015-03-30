package main

import (
	"log"
	"net"

	"github.com/flyrpc/flyrpc"

	. "./msg"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		log.Fatal(err)
	}
	client := fly.NewClient(conn, fly.Protobuf)
	client.OnMessage(3, func(ctx *fly.Context, msg *Hello) {
		log.Println("client on message", msg)
	})
	reply := &Hello{}
	if err := client.Call(1, reply, &Hello{Id: 123}); err != nil {
		log.Fatal(err)
	}
	log.Println("reply 1", reply)
	client.SendMessage(2, &Hello{Id: 123})
}
