package spot

import (
	"context"

	"github.com/adshao/go-binance/v2"
)

func (c *Client) GetAccountInfo() (Account, error) {
	account, err := binance.NewClient(c.ApiKey, c.ApiSecret).NewGetAccountService().Do(context.Background())
	if err != nil {
		return Account{}, err
	}
	return Account{Account: *account}, nil
}
