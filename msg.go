package discgo

type NewMessage struct {
	Content string `json:"content"`
	TTS     bool   `json:"tts"`
	//Embed   Embed  `json:"embed"`
}

type Embed struct {
}
