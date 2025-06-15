package unified

import (
	"context"

	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/adshao/go-binance/v2/portfolio"
)

type Client struct {
	*exchange.Client
}

func (c *Client) GetBalances() (AccountBalanceList, error) {
	balances, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetBalanceService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	list := AccountBalanceList{}
	for _, asset := range balances {
		list = append(list, *asset)
	}
	return list, nil
}
