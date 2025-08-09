package portfolio

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) GetCommissionRate(symbol string) (*portfolio.CommissionRate, error) {
	return portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetUMCommissionRateService().Symbol(symbol).Do(context.Background())
}
