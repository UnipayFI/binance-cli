package unified

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) AutoCollection() (*portfolio.SuccessResponse, error) {
	return portfolio.NewClient(c.ApiKey, c.ApiSecret).NewFundAutoCollectionService().Do(context.Background())
}
