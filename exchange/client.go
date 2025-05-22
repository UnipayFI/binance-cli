package exchange

type Client struct {
	ApiKey    string
	ApiSecret string
}

func NewClient(apiKey, apiSecret string) *Client {
	return &Client{ApiKey: apiKey, ApiSecret: apiSecret}
}
