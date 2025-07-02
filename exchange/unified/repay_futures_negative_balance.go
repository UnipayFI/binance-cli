package unified

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) RepayFuturesNegativeBalance() (*portfolio.SuccessResponse, error) {
	return portfolio.NewClient(c.ApiKey, c.ApiSecret).NewRepayFuturesNegativeBalanceService().Do(context.Background())
}
