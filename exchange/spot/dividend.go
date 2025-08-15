package spot

import (
	"context"

	"github.com/adshao/go-binance/v2"
)

func (c *Client) GetDividendHistory(asset string, startTime int64, endTime int64, limit int) (DividendList, error) {
	service := binance.NewClient(c.ApiKey, c.ApiSecret).NewAssetDividendService()
	if asset != "" {
		service.Asset(asset)
	}
	if startTime != 0 {
		service.StartTime(startTime)
	}
	if endTime != 0 {
		service.EndTime(endTime)
	}
	if limit != 0 {
		service.Limit(limit)
	}
	dividends, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return DividendList(*dividends.Rows), nil
}
