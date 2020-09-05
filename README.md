# Experimental Slack SDK in Go

This is an experimental Slack API client library.

```go
package main

import (
    "fmt"
    "io/ioutil"
    "net/url"
    "os"
    "github.com/seratch/go-slack-sdk-experimental/webapi"
    "github.com/seratch/go-slack-sdk-experimental/webapi/chat_postMessage"
)

func main() {
    data := url.Values{}
    data.Set("text", "Hi!")
    data.Set("channel", "#random")
    token := os.Getenv("SLACK_BOT_TOKEN")
    client := webapi.NewClient(&token)
    resp, err := client.CallApi(chat_postMessage.ApiMethod, data)
    if err != nil {
        fmt.Println(fmt.Sprintf("HTTP error: %s", err))
        return
    }
    bytes, _ := ioutil.ReadAll(resp.Body)
    body, _ := chat_postMessage.UnmarshalChatPostMessage(bytes)
    if !*body.Ok {
        fmt.Println(fmt.Sprintf("Slack error: %s", *body.Error))
    }
}
```

# License

The MIT License