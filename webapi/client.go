package webapi

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	Token   *string
	BaseURL string
}

func NewClient(token *string) Client {
	client := Client{
		Token:   token,
		BaseURL: "https://slack.com/api/",
	}
	return client
}

func (client *Client) CallApi(apiMethod string, data url.Values) (*http.Response, error) {
	apiEndpoint := fmt.Sprintf("%s%s", client.BaseURL, apiMethod)
	req, _ := http.NewRequest(
		"POST",
		apiEndpoint,
		strings.NewReader(data.Encode()),
	)
	if client.Token != nil {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *client.Token))
	}
	req.Header.Add("Content-Type", fmt.Sprintf("application/x-www-form-urlencoded"))
	httpClient := &http.Client{}
	return httpClient.Do(req)

}
