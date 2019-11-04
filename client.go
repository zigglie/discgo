package discgo

import (
	"encoding/json"
	"fmt"
	"time"

	socket "github.com/zigglie/websocks"
)

type Client struct {
	Token     string
	sock      *socket.Socket
	lastSeq   int
	heartbeat int
	events    EventHandler
}

type EventHandler struct {
	onMessage func(*MsgCreate)
}

var _discord string = "wss://gateway.discord.gg/?v=6&encoding=json"
var _echo string = "wss://echo.websocket.org"

func goReadBytes(c chan []byte, s *socket.Socket) {
	for {
		tmp := make([]byte, 1024)
		n, err := s.GetConn().Read(tmp)

		if err != nil {
			fmt.Println(tmp)
			panic(err)
		}

		if n != 0 {
			c <- tmp
		}
	}
}

func (c *Client) OnMessage(f func(*MsgCreate)) {
	c.events.onMessage = f
}

func NewClient(token string) *Client {
	c := Client{}
	c.events = EventHandler{}
	c.Token = token

	return &c
}

func (c *Client) Connect() {
	s := &socket.Socket{}

	err := s.Init(_discord)

	if err != nil {
		panic(err)
	}
	c.sock = s
	c.heartbeat = -1
	c.lastSeq = 0

	ch := make(chan []byte)
	go goReadBytes(ch, c.sock)

	beatingHearts := false

	c.sendIdentity()

	// go goSendBytes("Bytes", s)

	for {
		tmp := <-ch
		pck := socket.DecodePacket(&tmp)
		fmt.Println(pck.Msg)
		dj := DecodeJSON(pck)

		fmt.Printf("%+v %s\n", dj, dj.T)

		if dj.T == "MESSAGE_CREATE" {
			fmt.Printf("%+v\n", dj.D.(*MsgCreate))
			c.events.onMessage(dj.D.(*MsgCreate))
		}

		if dj.Opcode == 10 {
			c.heartbeat = dj.D.(*D).Heartbeat
			fmt.Println("Setting heartbeat")
			if !beatingHearts {
				go c.startHeartbeats()
				beatingHearts = true
			}
		}

		if dj.S != 0 {
			c.lastSeq = dj.S
		}
	}
}

func (c *Client) sendIdentity() {
	iden := Identify{c.Token, createProperties()}
	pay := ClientPayload{
		Opcode: 2,
		D:      iden,
	}
	p, err := json.Marshal(pay)

	if err != nil {
		panic(err)
	}

	fmt.Println("Debug json:", string(p))
	c.sock.SendMessage(string(p))
}

func (c *Client) sendHeartbeat() {
	cp := ClientPayload{}
	cp.Opcode = 1
	cp.D = c.lastSeq

	if c.lastSeq == 0 {
		cp.D = nil
	}

	p, err := json.Marshal(cp)

	if err != nil {
		panic(err)
	}

	fmt.Println("Sending heartbeat")
	fmt.Println(string(p))
	c.sock.SendMessage(string(p))
}

func (c *Client) startHeartbeats() {
	for {
		if c.heartbeat == -1 {

		} else {
			time.Sleep(time.Duration(c.heartbeat-3000) * time.Millisecond)
			c.sendHeartbeat()
		}
	}
}

func (c *Client) SendMessage(channelId, msg string) {
	payload := NewMessage{
		Content: msg,
		TTS:     false,
	}

	js, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	SendMessage(c.Token, channelId, string(js))
}
