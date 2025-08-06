package futures

import (
	"context"

	"github.com/adshao/go-binance/v2/futures"
)

func (c *Client) GetMultiAssetsMode() (bool, error) {
	multiAssetMode, err := futures.NewClient(c.ApiKey, c.ApiSecret).NewGetMultiAssetModeService().Do(context.Background())
	if err != nil {
		return false, err
	}
	return multiAssetMode.MultiAssetsMargin, nil
}

func (c *Client) SetMultiAssetsMode(multiAssetsMode bool) error {
	return futures.NewClient(c.ApiKey, c.ApiSecret).NewChangeMultiAssetModeService().MultiAssetsMargin(multiAssetsMode).Do(context.Background())
}
