package webapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"testing"
	"webapi/chat_postMessage"
)

func TestWebClient_CallApi(t *testing.T) {
	data := url.Values{}
	data.Set("text", "Hi!")
	data.Set("channel", "#random")
	token := os.Getenv("SLACK_BOT_TOKEN")
	client := NewClient(&token)
	//client.BaseURL = "http://locahost:3000/"
	resp, err := client.CallApi(chat_postMessage.ApiMethod, data)
	if err != nil {
		fmt.Println(fmt.Sprintf("HTTP error: %s", err))
		return
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	body, _ := chat_postMessage.UnmarshalChatPostMessage(bytes)
	bs, _ := json.Marshal(body)
	fmt.Println(string(bs))
	if !*body.Ok {
		fmt.Println(fmt.Sprintf("Slack error: %s", *body.Error))
		return
	}
}