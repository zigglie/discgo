package main

import (
	"fmt"
	"os"
	"time"

	"github.com/zigglie/discgo"
	"github.com/zigglie/envite"
)

func main() {
	envite.Load()
	c := discgo.NewClient(os.Getenv("discgo"))
	startTime := time.Now()

	c.OnMessage(func(m *discgo.MsgCreate) {
		fmt.Println(m.Content)
		if m.Content == ".dank" {
			c.SendMessage(m.ChannelId, "That's pretty dank <:FeelsDankMan:426470052414488600>")
		}
		if m.Content == ";uptime" {
			t := time.Now()
			elapsed := t.Sub(startTime)
			c.SendMessage(m.ChannelId, fmt.Sprintf("Uptime is %v <:Awoo:430840716873433099>", elapsed))
		}
		if m.Content == ";uptime nano" {
			t := time.Now()
			elapsed := t.Sub(startTime)
			message := fmt.Sprintf("Uptime is: %vns <:Awoo:430840716873433099>", elapsed.Nanoseconds())
			c.SendMessage(m.ChannelId, message)
		}
	})

	c.Connect()
}
