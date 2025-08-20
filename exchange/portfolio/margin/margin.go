package margin

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) GetMarginInterestHistory(asset string, startTime int64, endTime int64, current int64, size int64, archived bool) (MarginInterestHistoryList, error) {
	service := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetMarginInterestHistoryService()
	if asset != "" {
		service = service.Asset(asset)
	}
	if startTime != 0 {
		service = service.StartTime(startTime)
	}
	if endTime != 0 {
		service = service.EndTime(endTime)
	}
	if current != 0 {
		service = service.Current(current)
	}
	if size != 0 {
		service = service.Size(size)
	}
	if archived {
		service = service.Archived(archived)
	}
	interestHistory, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return MarginInterestHistoryList(interestHistory.Rows), nil
}
