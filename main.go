package go_slack_sdk_experimental

import (
	"encoding/json"
	"github.com/seratch/go-slack-sdk-experimental/web"
	"io/ioutil"
	"os"
)
import (
	"fmt"
	"net/url"
)

func main() {
	data := url.Values{}
	data.Set("text", "Hi!")
	data.Set("channel", "#random")
	token := os.Getenv("SLACK_BOT_TOKEN")
	fmt.Println(token)
	client := web.NewWebClient(&token)
	resp, _ := client.CallApi("chat.postMessage", data)
	bytes, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(bytes))
	body, _ := chat_postMessage.UnmarshalChatPostMessage(bytes)
	if body.Message != nil {
		msg, _ := json.Marshal(*body.Message)
		fmt.Println(string(msg))
	}
}
