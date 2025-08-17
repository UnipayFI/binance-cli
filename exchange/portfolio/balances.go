package portfolio

import (
	"context"

	"github.com/UnipayFI/binance-cli/common"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/adshao/go-binance/v2/portfolio"
)

type Client struct {
	*exchange.Client
}

func NewClient(apiKey, apiSecret string) *Client {
	return &Client{Client: exchange.NewClient(apiKey, apiSecret)}
}

func (c *Client) GetBalances() (AccountBalanceList, error) {
	balances, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetBalanceService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	list := AccountBalanceList{}
	for _, asset := range balances {
		if !common.IsZero(asset.TotalWalletBalance) {
			list = append(list, *asset)
		}
	}
	return list, nil
}
