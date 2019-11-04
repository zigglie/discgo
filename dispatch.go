// Structs for Opcode 0 AKA dispatch events
// These are pretty much the meat and potatoes of discord events

package discgo

type MsgCreate struct {
	Type      int    `json:"type"` // Not sure what this is tbh
	TTS       bool   `json:"tts"`
	Timestamp string `json:"timestamp"`
	Pinned    bool   `json:"pinned"`
	//Sender
	Author    Author `json:"author"`
	Id        string `json:"id"`
	GuildId   string `json:"guild_id"`
	ChannelId string `json:"channel_id"`
	Content   string `json:"content"`
}

type Author struct {
	Id string `json:"id"`
}
