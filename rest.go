package discgo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetChannelURL(channelId string) *url.URL {
	post := "https://discordapp.com/api/channels/" + channelId + "/messages"
	u, _ := url.Parse(post)
	return u
}

func SendMessage(token, channelId, msg string) {
	fmt.Println(msg)
	nb := ioutil.NopCloser(bytes.NewBufferString(msg))
	fmt.Println("\n\n\n\n\n", nb)

	h := http.Header{}
	h.Add("Authorization", "Bot "+token)
	h.Add("Content-Type", "application/json")

	url := GetChannelURL(channelId)
	req := http.Request{
		Method: "POST",
		Header: h,
		URL:    url,
		Body:   nb,
	}

	c := http.Client{}
	resp, err := c.Do(&req)

	if err != nil {
		panic(err)
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err2 != nil {
		panic(err2)
	}

	fmt.Println(string(body))
}
