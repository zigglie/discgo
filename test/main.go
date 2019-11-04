package main

import (
	"fmt"
	"os"

	discgo "github.com/zigglie/discgo"
	"github.com/zigglie/envite"
	socket "github.com/zigglie/websocks"
)

var _discord string = "wss://gateway.discord.gg/?v=6&encoding=json"

func goReadBytes(c chan []byte, s *socket.Socket) {
	for {
		tmp := make([]byte, 1024)
		n, err := s.GetConn().Read(tmp)

		if err != nil {
			panic(err)
		}

		if n != 0 {
			c <- tmp
		}
	}
}

func main() {
	envite.Load()
	c := discgo.NewClient(os.Getenv("discgo"))

	c.OnMessage(func(m *discgo.MsgCreate) {
		fmt.Println(m.Content)
		if m.Author.Id == "110927968133464064" && m.Content == "Hello?" {
			c.SendMessage(m.ChannelId, "Am-am.. I alive??? <:monkS:426470529709506578>")
		}
	})

	c.Connect()
}
