package unified

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) GetAccount() (AccountInfo, error) {
	account, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetAccountService().Do(context.Background())
	if err != nil {
		return AccountInfo{}, err
	}
	return AccountInfo(*account), nil
}
