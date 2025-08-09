package portfolio

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) AutoCollection() (*portfolio.SuccessResponse, error) {
	return portfolio.NewClient(c.ApiKey, c.ApiSecret).NewFundAutoCollectionService().Do(context.Background())
}

func (c *Client) AssetCollection(asset string) (*portfolio.SuccessResponse, error) {
	return portfolio.NewClient(c.ApiKey, c.ApiSecret).NewFundCollectionByAssetService().Asset(asset).Do(context.Background())
}
