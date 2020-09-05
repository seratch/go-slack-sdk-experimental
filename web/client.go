package web

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type WebClient struct {
	Token *string
	BaseURL string
}

func NewWebClient(token *string) WebClient {
	client := WebClient{
		Token: token,
		BaseURL: "https://slack.com/api/",
	}
	return client
}

func (client *WebClient) CallApi(apiMethod string, data url.Values) (*http.Response, error) {
	apiEndpoint := fmt.Sprintf("https://slack.com/api/%s", apiMethod)
	req, _ := http.NewRequest(
		"POST",
		apiEndpoint,
		strings.NewReader(data.Encode()),
	)
	if client.Token != nil {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %q", &client.Token))
	}
	req.Header.Add("Content-Type", fmt.Sprintf("application/x-www-form-urlencoded"))
	httpClient := &http.Client{}
	return httpClient.Do(req)

}