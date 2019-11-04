package discgo

import (
	"encoding/json"
	"runtime"

	socket "github.com/zigglie/websocks"
)

// Packet to unmarshal to
// https://discordapp.com/developers/docs/topics/gateway#payloads
//
// FIELD	TYPE		DESCRIPTION								PRESENT
//	op		integer		opcode for the payload						Always
//	d		?mixed 		(any JSON value) event data					Always
//	s		integer		sequence number, used for resuming
//						sessions and heartbeats						Only for Opcode 0
//
//	t		string		the event name for this payload				Only for Opcode 0
//
//

type Payload struct {
	Opcode int         `json:"op"`
	D      interface{} `json:"d"`
	T      string      `json:"t"`
	S      int         `json:"s"`
}

type D struct {
	Type      string `json:"t"`
	Heartbeat int    `json:"heartbeat_interval"`
}

type ClientPayload struct {
	Opcode int         `json:"op"`
	D      interface{} `json:"d"`
	T      string      `json:"t"`
	S      int         `json:"s"`
}

type Identify struct {
	Token string     `json:"token"`
	Props Properties `json:"properties"`
}

//$os	string	your operating system
//$browser	string	your library name
//$device	string	your library name
//https://discordapp.com/developers/docs/topics/gateway#identify
type Properties struct {
	OS      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

func DecodeJSON(p socket.Packet) Payload {
	test := Payload{}
	test.D = &D{}
	json.Unmarshal([]byte(p.Msg), &test)

	if test.T == "MESSAGE_CREATE" {
		test.D = &MsgCreate{}
	}

	json.Unmarshal([]byte(p.Msg), &test)
	return test
}

func createProperties() Properties {
	return Properties{
		runtime.GOOS,
		"discgo",
		"discgo",
	}
}
