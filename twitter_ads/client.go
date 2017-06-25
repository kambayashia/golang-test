package twitter_ads

import (
	"github.com/dghubble/sling"
	"net/http"
	"fmt"
)

const baseUrl = "https://ads-api-sandbox.twitter.com/"

// Client is a Twitter client for making Twitter API requests.
type Client struct {
	sling *sling.Sling
  AccountService *AccountService
}

type ApiError struct {
	Errors []ErrorEntity `json:"errors"`
}

type ErrorEntity struct {
	Message string `json:"message"`
	Code int `json:"code"`
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(baseUrl)

	return &Client{
		sling: base,
		AccountService: NewAccountService(base),
	}
}

func ApiVersion(sling *sling.Sling, version int) *sling.Sling {
	path := fmt.Sprintf("%d/", 1)
	return sling.New().Path(path)
}
