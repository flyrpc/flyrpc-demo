package main

import (
	"log"
	"net"

	"github.com/flyrpc/flyrpc"
	. "github.com/flyrpc/flyrpc-demo/go/msg"
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
	done := make(chan int, 1)
	client.OnMessage(1, func(ctx *fly.Context, m *Hello) {
		reply := &Hello{}
		if err := client.Call(1, reply, &Hello{Id: 123}); err != nil {
			log.Fatal(err)
		}
		log.Println("reply 1", reply)
		client.SendMessage(2, &Hello{Id: 123})
		log.Println("done")
		done <- 1
	})
	<-done
}
